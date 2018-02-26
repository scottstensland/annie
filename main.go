package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/iawia002/annie/extractors"
	"github.com/iawia002/annie/utils"
)

var (
	debug   bool
	version bool
)

func init() {
	flag.BoolVar(&debug, "d", false, "Debug mode")
	flag.BoolVar(&version, "v", false, "Show version")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("error")
		return
	}
	videoURL := args[0]
	u, err := url.ParseRequestURI(videoURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	domain := utils.Domain(u.Host)
	switch domain {
	case "douyin":
		extractors.Douyin(videoURL)
	default:
		fmt.Println("unsupported URL")
	}
}