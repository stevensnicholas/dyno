package main

import (
	"dyno/internal/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

type content struct {
	Data_yaml []byte `yaml:"result"`
	Data_json []byte `json:"result"`
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func isYAML(s string) bool {
	var js map[string]interface{}
	return yaml.Unmarshal([]byte(s), &js) == nil
}

func sendRequest(requestBody []byte, url string, contentType string) {
	con:=strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", url, con)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", contentType)
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

		url := "http://localhost:8080/openapi"
		logger.Debugf("URL:>", url)
		logger.Debugf("ss", data)
		
		var readFileContent []byte
		var requestBody []byte
		readFileContent,_=ioutil.ReadFile(args.Send.Path)

		if isJSON(string(fileData)) || isYAML(string(fileData)) {
			request_payload:=content{Data_json : readFileContent}
			requestBody,_=json.Marshal(request_payload)
			sendRequest(requestBody, url, "application/json")
		} else {
			logger.Fatal("Please provide either a JSON or YAML file")
		}

	default:
		p := arg.MustParse(&args)
		p.WriteHelp(os.Stdout)
	}
}
