package main

import (
	"dyno/internal/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alexflint/go-arg"
	"gopkg.in/yaml.v3"
)

type SendCmd struct {
	Path string `arg:"positional"`
}

var args struct {
	Send    *SendCmd `arg:"subcommand:send" help:"can also use -d to provide the path to file"`
	Verbose bool     `arg:"-v" help:"enable verbose logging"`
	Debug   bool     `arg:"-d" help:"enable debug logging"`
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func isYAML(s string) bool {
	var js map[string]interface{}
	return yaml.Unmarshal([]byte(s), &js) == nil
}

func main() {
	arg.MustParse(&args)

	// Configure logger
	logLevel := "warn"
	if args.Verbose {
		logLevel = "info"
	}
	if args.Debug {
		logLevel = "debug"
	}
	log, err := logger.ConfigureDevelopmentLogger(logLevel)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	switch {
	case args.Send != nil && args.Send.Path != "":
		logger.Infow("Getting file from location %s\n", args.Send.Path)

		data, err := os.Open(args.Send.Path)
		if err != nil {
			logger.Fatalf("Error", err)
		}
		fileData, _ := ioutil.ReadAll(data)

		url := "https://o8cnchwjji.execute-api.ap-southeast-2.amazonaws.com/v1/post_json"
		logger.Debugf("URL:>", url)
		logger.Debugf("ss", data)

		req, err := http.NewRequest("POST", url, data)
		if err != nil {
			panic(err)
		}

		if isJSON(string(fileData)) {
			req.Header.Set("Content-Type", "application/json")
		} else if isYAML(string(fileData)) {
			req.Header.Set("Content-Type", "application/x-yaml")
		} else {
			logger.Fatal("Please provide either a JSON or YAML file")
		}

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

	default:
		p := arg.MustParse(&args)
		p.WriteHelp(os.Stdout)
	}
}
