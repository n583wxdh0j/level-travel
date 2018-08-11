package controllers

import (
	"context"
	"fmt"
	"level-travel/config"
	"level-travel/models"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// func UpdateLibrariesData(timeUpdate time.Duration) {
// 	// TODO: check last update data, if new udate countinue, else return
// 	libs := models.Libraries{}
// 	for {
// 		// API
// 		client, context := getGithubClient()
// 		fillStarsCommitDateAndSubscribers(*libs, client, context)

// 		// write new data to DB
// 		database.UpdateLibrariesData()
// 		time.Sleep(timeUpdate)
// 	}
// }

func getGithubClient() (*github.Client, *context.Context) {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GetConfig().GithubAPIKey},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	return github.NewClient(tokenClient), &context
}

func fillStarsCommitDateAndSubscribers(libs *models.Libraries, client *github.Client, context *context.Context, owner, repo string) {
	listCommits, _, err := client.Repositories.ListCommits(nil, owner, repo, nil)
	if err != nil {
		log.Println("Get LastCommits Erorr: ", err)
	}
	fmt.Printf("%+v\n", listCommits[0].Commit.Committer.Date)

	repository, _, err := client.Repositories.Get(*context, owner, repo)
	if err != nil {
		log.Println("Get Repositories Erorr: ", err)
	}
	fmt.Printf("%+v\n", *repository.StargazersCount)
	fmt.Printf("%+v\n", *repository.SubscribersCount)

}
