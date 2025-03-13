package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hackermanpeter/github-activity/config"
	"github.com/hackermanpeter/github-activity/internal"
)

var conf = config.New()

func main() {
	ctx := context.Background()

	// get user input from cmd
	args, err := internal.ReceiveArguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fetch data from github
	req := internal.BuildRequest(ctx, conf, args)
	resp, err := internal.Client.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Unable to reach github ~ Status Code: %v", resp.StatusCode)
	}

	// decode data
	var githubResponse []internal.GithubResponse

	err = json.NewDecoder(resp.Body).Decode(&githubResponse)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	// parse data
	response := internal.ParseGithubResponse(githubResponse)

	result := internal.FormatOutput(args.Username, response)
	fmt.Print(result)

}
