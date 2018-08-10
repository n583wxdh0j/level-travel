package controllers

import (
	"context"
	"fmt"
	"level-travel/config"
	"level-travel/models"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/net/html"
	"golang.org/x/oauth2"
)

func UpdateLibrariesData(timeUpdate time.Duration) {
	libs := models.Libraries{}
	for {
		// parser
		parserData := parser()
		// API
		for _, val := range parserData {
			client, context := getGithubClient()
			fillStarsCommitDateAndSubscribers(*libs, client, context)
		}

		// write new data to DB

		time.Sleep(timeUpdate)
	}
}

func parser() models.Libraries {
	req, err := http.NewRequest("GET", "https://github.com/avelino/awesome-go", nil)
	if err != nil {
		log.Println("Get page error: ", err)
	}

	doc, err := html.Parse(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func getGithubClient() (*github.Client, context.Context) {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GetConfig().GithubAPIKey},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	return github.NewClient(tokenClient), context
}

func fillStarsCommitDateAndSubscribers(libs *models.Libraries, client *github.Client, context context.Context, owner, repo string) {
	listCommits, _, err := client.Repositories.ListCommits(nil, owner, repo, nil)
	if err != nil {
		log.Println("Get LastCommits Erorr: ", err)
	}
	fmt.Printf("%+v\n", listCommits[0].Commit.Committer.Date)

	repository, _, err := client.Repositories.Get(context, owner, repo)
	if err != nil {
		log.Println("Get Repositories Erorr: ", err)
	}
	fmt.Printf("%+v\n", *repository.StargazersCount)
	fmt.Printf("%+v\n", *repository.SubscribersCount)

}
