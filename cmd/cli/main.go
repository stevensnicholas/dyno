package main

import (
	"flag"
	"fmt"
	"golambda/internal/logger"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alexflint/go-arg"
)

type SendCmd struct {
	Path string `arg:"positional"`
	Dest bool   `arg:"-d"`
}

var args struct {
	Send *SendCmd `arg:"subcommand:send" help:"can also use -d to provide the path to file"`
}

func main() {
	arg.MustParse(&args)

	var logLevel = flag.String("loglevel", "info", "Log level")
	log, err := logger.ConfigureDevelopmentLogger(*logLevel)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	switch {
	case args.Send != nil:
		logger.Infow("Getting file from location %s\n", args.Send.Path)

		data, err := os.Open(args.Send.Path)
		if err != nil {
			logger.Fatalf("Error", err)
		}
		// will be replaced with actual api-endpoint,
		url := "https://o8cnchwjji.execute-api.ap-southeast-2.amazonaws.com/v1/post_json"
		fmt.Println("URL:>", url)

		req, err := http.NewRequest("POST", url, data)

		if err != nil {
			panic(err)
		}

		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		logger.Infof(
			"response Status:", resp.Status,
			"response Headers:", resp.Header,
			"response Body:", string(body),
		)
	case args.Send == nil:
		logger.Error("Invalid command, use -h to get help")
	}

}
