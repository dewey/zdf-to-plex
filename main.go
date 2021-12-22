package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/peterbourgon/diskv/v3"
	"github.com/peterbourgon/ff/v3"
)

var (
	reSeasonNumber = regexp.MustCompile(`Staffel\s(\d+)`)
)

func main() {
	fs := flag.NewFlagSet("zdf-to-plex", flag.ExitOnError)
	var (
		cachePath        = fs.String("cache-path", "cache", "provide where we successful download urls are cached.")
		youtubedlBinPath = fs.String("ytdl-bin", "/opt/homebrew/bin/yt-dlp", "path to youtube-dl / yt-dlp binary.")
		httpProxyAddress = fs.String("http-proxy-address", "", "http proxy string to be used for downloading. example: username:password@10.0.0.1:1234")
		entrypoints      = fs.String("entrypoints", "https://www.zdf.de/serien/inspector-barnaby", "comma separated list of zdf shows to download.")
		outputBasePath   = fs.String("output-base-path", "", "output path where downloaded files are stored.")
		debug            = fs.Bool("debug", false, "debug information and showing yt-dlp stdout.")
	)
	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarNoPrefix(),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)

	w := log.NewSyncWriter(os.Stderr)
	l := log.NewLogfmtLogger(w)
	level.Info(l).Log("msg", "starting zdf-to-plex")

	d := diskv.New(diskv.Options{
		BasePath: *cachePath,
	})

	resp, err := http.Get("https://zdf-cdn.live.cellular.de/static/export/live/google/catalog.json")
	if err != nil {
		level.Error(l).Log("err", err)
		return
	}
	var c CatalogData
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		level.Error(l).Log("err", err)
		return
	}
	var tasks []task
	for _, s := range c.DataFeedElement {
		// Only whitelisted entrypoints are being scraped
		for _, entrypoint := range strings.Split(*entrypoints, ",") {
			if s.PartOfSeries.ID == entrypoint && s.Type == "TVEpisode" {
				seasonNumber, err := extractSeasonNumber(s.URL)
				if err != nil {
					level.Error(l).Log("err", err)
					return
				}
				clearEpisodeName := fmt.Sprintf("S%dE%d - %s", seasonNumber, s.EpisodeNumber, s.Name)
				tasks = append(tasks, task{
					OutputFile: filepath.Join(*outputBasePath, fmt.Sprintf("%s/%s/%s.%%(ext)s", s.PartOfSeries.Name, fmt.Sprintf("Season %d", seasonNumber), clearEpisodeName)),
					ClearName:  clearEpisodeName,
					URL:        s.URL,
				})
			}
		}
	}

	level.Info(l).Log("msg", "collected tasks", "count", len(tasks))
	for _, t := range tasks {
		// Skip if we already downloaded the file previously
		if exists := d.Has(t.URL); exists {
			continue
		}
		level.Info(l).Log("msg", "working on task", "url", t.URL, "clear_name", t.ClearName, "output_path", t.OutputFile)

		// Building yt-dlp command to execute
		var args []string
		if *httpProxyAddress != "" {
			level.Info(l).Log("msg", "found proxy, using that to fetch file")
			args = append(args, "--proxy")
			args = append(args, *httpProxyAddress)
		}
		args = append(args, "--concurrent-fragments")
		args = append(args, "10")
		args = append(args, "-o")
		args = append(args, t.OutputFile)
		args = append(args, t.URL)
		cmd := exec.Command(*youtubedlBinPath, args...)

		// Show output in Stdout
		if *debug {
			cmd.Stdout = os.Stdout
		} else {
			fmt.Println("not")
		}

		if err := cmd.Run(); err != nil {
			level.Error(l).Log("err", err)
			return
		}

		// Cache is successful
		if err := d.Write(t.URL, []byte(t.ClearName)); err != nil {
			level.Error(l).Log("err", err)
			return
		}
	}
}

// extractSeasonNumber extracts the season number from ZDF website as the feed doesn't have the correct one and always uses season 0.
func extractSeasonNumber(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return 0, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return 0, err
	}

	var seasonNumber int
	doc.Find("div.details > span.teaser-cat").Each(func(i int, s *goquery.Selection) {
		matches := reSeasonNumber.FindStringSubmatch(s.Text())
		if len(matches) == 2 {
			sn, err := strconv.Atoi(matches[1])
			if err != nil {
				return
			}
			seasonNumber = sn
		}
	})
	if seasonNumber != 0 {
		return seasonNumber, nil
	} else {
		return 0, errors.New("couldn't extract season number")
	}
}

// task is an internal download task
type task struct {
	OutputFile string
	ClearName  string
	URL        string
}

// CatalogData as defined by ZDF feed
type CatalogData struct {
	Type            string    `json:"@type"`
	DateModified    time.Time `json:"dateModified"`
	DataFeedElement []struct {
		Type          string `json:"@type"`
		ID            string `json:"@id"`
		URL           string `json:"url"`
		SameAs        string `json:"sameAs,omitempty"`
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		EpisodeNumber int    `json:"episodeNumber,omitempty"`
		PartOfSeason  struct {
			Type         string `json:"@type"`
			SeasonNumber int    `json:"seasonNumber"`
		} `json:"partOfSeason,omitempty"`
		PartOfSeries struct {
			Type string `json:"@type"`
			ID   string `json:"@id"`
			Name string `json:"name"`
		} `json:"partOfSeries,omitempty"`
	} `json:"dataFeedElement"`
}
