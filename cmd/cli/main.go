package main

import (
	"encoding/json"
	"flag"
	"golambda/internal/logger"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alexflint/go-arg"
	"gopkg.in/yaml.v3"
)

type SendCmd struct {
	Path string `arg:"positional"`
	Dest bool   `arg:"-d"`
}

var args struct {
	Send *SendCmd `arg:"subcommand:send" help:"can also use -d to provide the path to file"`
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

	var logLevel = flag.String("loglevel", "info", "Log level")
	log, err := logger.ConfigureDevelopmentLogger(*logLevel)
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
		logger.Infof("ss", data)
		fileData, _ := ioutil.ReadAll(data)

		inputJSON := isJSON(string(fileData))
		inputYAML := isYAML(string(fileData))

		if inputJSON || inputYAML {
			url := "https://o8cnchwjji.execute-api.ap-southeast-2.amazonaws.com/v1/post_json"
			logger.Infof("URL:>", url)
			logger.Infof("ss", data)
			data, _ := os.Open(args.Send.Path)
			req, err := http.NewRequest("POST", url, data)

			if err != nil {
				panic(err)
			}

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
		} else {
			logger.Error("Please provide either JSON or YAML")
		}

	case args.Send == nil || args.Send.Path == "":
		p := arg.MustParse(&args)
		p.WriteHelp(os.Stdout)
	}

}
