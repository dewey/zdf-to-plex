# ZDF to Plex

Downloads a fixed list of freely available shows from ZDF Mediathek and stores them in a directory,
correctly named to be picked up by Plex. It wraps yt-dlp to implement the actual downloading.

Optional: If `TWILIO_*` environment variables are set it sends a SMS notification to the provided number once
a new episode is available.

It is supposed to run in a cron job to pick up new episodes, the successful downloads are cached so that we don't redownload them.

# Configuration

The following environment variables can be set to modify the behaviour of the tool.

- `DEBUG`: Print yt-dlp output and other debug messages
- `HTTP_PROXY_ADDRESS`: A http proxy that is used for the download. Example: username:password@10.0.0.1:10000
- `ENTRYPOINTS`: Comma separated list of shows listed on the ZDF Mediathek website
- `YTDL_BIN`: Path to the yt-dlp binary on your system. Run `which yt-dlp` after installing to see where it's located.
- `CACHE_PATH`: Path where to store the cache file containing a list of URLs that were already downloaded.
- `OUTPUT_BASE_PATH`: Path where the downloaded show will be stored. A directory containing the show name will be created if it doesn't exist at the base path. 
- `TWILIO_SUBACCOUNT_ID`: The Twilio account ID
- `TWILIO_AUTH_TOKEN`: The Twilio auth token
- `TWILIO_FROM_NUMBER`: The phone number SMS notifications from Twilio are sent from
- `TWILIO_TO_NUMBER`: The phone number SMS notifications from Twilio are sent to

# Usage

```
DEBUG=TRUE HTTP_PROXY_ADDRESS=username:password@10.0.0.1:10000 ENTRYPOINTS=https://www.zdf.de/serien/inspector-barnaby YTDL_BIN=/usr/local/bin/yt-dlp CACHE_PATH=zdf_to_plex.cache OUTPUT_BASE_PATH=/home/ubuntu/media/TV-Misc ./zdf-to-plex-amd64-linux
```

# Development

- Set up environment variables in `develop.env`
- Run `./run_develop.sh` to run the tool