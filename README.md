# rss2json
A Go program to convert an RSS feed into a JSON object. The program uses the [gofeed](https://github.com/mmcdole/gofeed) parser, which can take either an `RSS` or `Atom` feed. You can either pass the data in through `stdin`, or by setting a url flag. The flag will take precendence over `stdin` if both are used for some reason.

# Command-line options

### -url string

The program will fetch the file and parse it using the `gofeed.ParseUrl()` function.
