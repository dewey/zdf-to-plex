package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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

	"github.com/twilio/twilio-go"

	"github.com/pkg/errors"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/peterbourgon/diskv/v3"
	"github.com/peterbourgon/ff/v3"
)

var (
	reSeasonNumber = regexp.MustCompile(`Staffel\s(\d+).+Folge\s(\d+)`)
)

func main() {
	fs := flag.NewFlagSet("zdf-to-plex", flag.ExitOnError)
	var (
		cachePath             = fs.String("cache-path", "cache", "provide where we successful download urls are cached.")
		youtubedlBinPath      = fs.String("ytdl-bin", "/opt/homebrew/bin/yt-dlp", "path to youtube-dl / yt-dlp binary.")
		httpProxyAddress      = fs.String("http-proxy-address", "", "http proxy string to be used for downloading. example: username:password@10.0.0.1:1234")
		entrypoints           = fs.String("entrypoints", "https://www.zdf.de/serien/inspector-barnaby", "comma separated list of zdf shows to download.")
		outputBasePath        = fs.String("output-base-path", "", "output path where downloaded files are stored.")
		twilioSubaccountID    = fs.String("twilio-subaccount-id", "", "account sid of twilio account used for sending.")
		twilioPhoneFromNumber = fs.String("twilio-phone-from-number", "", "number notifications are sent from.")
		twilioPhoneToNumber   = fs.String("twilio-phone-to-number", "", "number notifications are sent to.")
		twilioAuthToken       = fs.String("twilio-auth-token", "", "auth token of twilio account used for sending.")
		debug                 = fs.Bool("debug", false, "debug information and showing yt-dlp stdout.")
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

	var twilioClient *twilio.RestClient
	if *twilioPhoneFromNumber != "" && *twilioAuthToken != "" && *twilioSubaccountID != "" && *twilioPhoneToNumber != "" {
		client := twilio.NewRestClientWithParams(twilio.RestClientParams{
			Username: *twilioSubaccountID,
			Password: *twilioAuthToken,
		})
		twilioClient = client
		level.Info(l).Log("msg", "found twilio credentials, initializing sms notifications")
	}

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
				seasonNumber, episodeNumber, err := extractSeasonEpisodeNumber(s.URL)
				if err != nil {
					level.Error(l).Log("err", err)
					return
				}
				clearEpisodeName := fmt.Sprintf("S%dE%d - %s", seasonNumber, episodeNumber, s.Name)
				tasks = append(tasks, task{
					OutputFile:    filepath.Join(*outputBasePath, fmt.Sprintf("%s/%s/%s.%%(ext)s", s.PartOfSeries.Name, fmt.Sprintf("Season %d", seasonNumber), clearEpisodeName)),
					ClearName:     clearEpisodeName,
					ClearNameShow: s.PartOfSeries.Name,
					EpisodeName:   s.Name,
					URL:           s.URL,
				})
			}
		}
	}

	level.Info(l).Log("msg", "collected tasks", "count", len(tasks))
	for _, t := range tasks {
		// Skip if we already downloaded the file previously
		hashedURL := hashURL(t.URL)
		if exists := d.Has(hashedURL); exists {
			level.Info(l).Log("msg", "url already in cache, skipping", "url", t.URL, "md5", hashedURL)
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
		// Progress bar as new lines, so that it looks more readable in log files
		args = append(args, "--newline")
		args = append(args, t.URL)
		cmd := exec.Command(*youtubedlBinPath, args...)

		// Show output in Stdout
		var combinedOutput []byte
		if *debug {
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				level.Error(l).Log("err", errors.Wrap(err, "running exec command"))
				return
			}
		} else {
			out, err := cmd.CombinedOutput()
			if err != nil {
				level.Error(l).Log("err", errors.Wrap(err, "running exec command collecting combined output"))
				return
			}
			combinedOutput = out
		}

		// Cache is successful
		if err := d.Write(hashURL(t.URL), []byte(t.ClearName)); err != nil {
			level.Error(l).Log("err", errors.Wrap(err, "writing cache key"))
			return
		}

		// Send notification of new episode
		if twilioClient != nil {
			// Don't send a notification if we already downloaded the file
			if strings.Contains(string(combinedOutput), "already been downloaded") {
				level.Info(l).Log("msg", "already been downloaded, skip sending a notification")
				continue
			}
			message := fmt.Sprintf("Neue %s Folge in Plex: %s", t.ClearNameShow, t.EpisodeName)
			if err := sendNotification(l, twilioClient, *twilioPhoneFromNumber, *twilioPhoneToNumber, message); err != nil {
				level.Error(l).Log("err", errors.Wrap(err, "sending notification"))
				return
			}
		}
	}
}

// extractSeasonNumber extracts the season number from ZDF website as the feed doesn't have the correct one and always uses season 0.
func extractSeasonEpisodeNumber(url string) (int, int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return 0, 0, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return 0, 0, err
	}

	var seasonNumber, episodeNumber int
	doc.Find("div.details > span.teaser-cat").Each(func(i int, s *goquery.Selection) {
		matches := reSeasonNumber.FindStringSubmatch(strings.ReplaceAll(s.Text(), "\n", ""))
		if len(matches) == 3 {
			sn, err := strconv.Atoi(matches[1])
			if err != nil {
				return
			}
			en, err := strconv.Atoi(matches[2])
			if err != nil {
				return
			}
			seasonNumber = sn
			episodeNumber = en
		}
	})
	if seasonNumber != 0 {
		return seasonNumber, episodeNumber, nil
	} else {
		return 0, 0, errors.New("couldn't extract season number")
	}
}

func sendNotification(l log.Logger, client *twilio.RestClient, fromNumber string, toNumber string, message string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(fromNumber)
	params.SetBody(message)

	respAPI, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		level.Error(l).Log("err", err)
		return err
	}
	level.Info(l).Log("msg", "send notification via sms", "to", respAPI.To, "from", fromNumber, "message", message)
	return nil
}

func hashURL(url string) string {
	binHash := md5.Sum([]byte(url))
	return hex.EncodeToString(binHash[:])
}

// task is an internal download task
type task struct {
	// OutputFile is the final name of the file
	OutputFile string
	// ClearName is the clear name of the episode in the format: SxxExx - Episode Clear Name
	ClearName string
	// ClearNameShow is the clear name of the show
	ClearNameShow string
	// EpisodeName is the name of the episode, without any season or episode number information
	EpisodeName string
	// URL is the download URL of the episode
	URL string
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
