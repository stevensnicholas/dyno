package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "github.com/alexflint/go-arg"
)

func main() {
	type SendCmd struct {
		Path string `arg:"positional"`
		Dest bool   `arg:"-d"`
	}
	var args struct {
		Send *SendCmd `arg:"subcommand:send" help:"can also use -d to provide the path to file"`
	}
	
	arg.MustParse(&args)
	
	switch {
	case args.Send != nil:
		fmt.Printf("Getting file from location %s\n", args.Send.Path)

		data, err := os.Open(args.Send.Path)
		if err != nil {
				log.Fatal(err)
		}
		// will be replaced with actual api-endpoint,
		url := "https://o8cnchwjji.execute-api.ap-southeast-2.amazonaws.com/v1/post_json"
		fmt.Println("URL:>", url)

		req, err := http.NewRequest("POST", url, data)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		//log, err := logger.ConfigureProductionLogger(*logLevel)
		if err != nil {
				panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	case args.Send == nil:
			fmt.Println("Invalid command, use -h to get help")
	}

}