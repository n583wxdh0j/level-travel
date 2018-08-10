package database

import (
	"database/sql"
	"fmt"
	"level-travel/config"
	"level-travel/models"
	"log"

	_ "github.com/lib/pq"
)

func UpdateLibrariesData(libraries models.Libraries) error {
	connStr := createConnStrToDB()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(
		fmt.Sprintf(
			"INSERT INTO library_list" +
			"(name, link, stars, last_commit_date, followers, description)" +
			"VALUES" +
			// TODO: TIME
			"('%s', '%s', %d, '%v', %d, '%s');",
		)
		
	)
}

func createConnStrToDB() string {
	return fmt.Sprintf(
		"user=%s dbname=%s sslmode=%s",
		config.GetConfig().Database.User,
		config.GetConfig().Database.Name,
		config.GetConfig().Database.SSLMode,
	)
}