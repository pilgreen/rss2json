package main

import (
  "encoding/json"
  "flag"
  "os"

  "github.com/mmcdole/gofeed"
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  var url string
  var feed *gofeed.Feed
  var err error

  flag.StringVar(&url, "url", "", "The feed url.")
  flag.Parse()

  fp := gofeed.NewParser()
  fi, _ := os.Stdin.Stat()

  if len(url) > 0 {
    feed, err = fp.ParseURL(url)
    check(err)
  } else if fi.Mode() & os.ModeNamedPipe != 0 {
    feed, err = fp.Parse(os.Stdin)
    check(err)
  }

  b, err := json.Marshal(feed)
  check(err)

  os.Stdout.Write(b)
}
