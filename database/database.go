package database

import (
	"database/sql"
	"fmt"
	"level-travel/config"
	"level-travel/models"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

var space = " "

func UpdateLibraries(libs *models.Libraries) error {
	connStr := createConnStrToDB()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Open DB error", err)
		return err
	}
	defer db.Close()
	updQuery := createUpdQuery(libs)
	_, err = db.Query(updQuery)
	if err != nil {
		log.Println("Query error", err)
		return err
	}
	return err
}

func GetLibraries() (*models.Libraries, error) {
	connStr := createConnStrToDB()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Open DB error", err)
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM libraries")
	if err != nil {
		log.Println("Get libraries from DB error", err)
		return nil, err
	}
	var libs models.Libraries
	for rows.Next() {
		var lib models.Library
		err := rows.Scan(
			&lib.ID,
			&lib.Repo,
			&lib.Owner,
			&lib.Link,
			&lib.Stars,
			&lib.LastCommitDate,
			&lib.Followers,
			&lib.Description,
			&lib.LastUpdate,
		)
		if err != nil {
			log.Println("Get libaries row error", err)
			return nil, err
		}
		libs = append(libs, lib)
	}
	return &libs, err
}

func createConnStrToDB() string {
	return fmt.Sprintf(
		"user=%s dbname=%s sslmode=%s",
		config.GetConfig().Database.User,
		config.GetConfig().Database.Name,
		config.GetConfig().Database.SSLMode,
	)
}

func createUpdQuery(libs *models.Libraries) string {
	var query strings.Builder
	query.WriteString(
		fmt.Sprint(
			"INSERT INTO libraries",
			space,
			"(",
			"id,",
			"repo,",
			"owner,",
			"link,",
			"stars,",
			"last_commit_date,",
			"followers,",
			"description,",
			"last_update",
			")",
			space,
			"VALUES",
			space,
		),
	)
	for i, val := range *libs {
		query.WriteString(
			fmt.Sprintf(
				"('%s', '%s', '%s', %d, '%s', %d, '%s', '%s')",
				val.ID,
				val.Repo,
				val.Owner,
				val.Link,
				val.Stars,
				val.LastCommitDate,
				val.Followers,
				val.Description,
				val.LastUpdate,
			),
		)
		if i != len(*libs)-1 {
			query.WriteString(",")
		}
	}
	query.WriteString(";")
	return query.String()
}
