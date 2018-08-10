package models

import (
	"database/sql"
	"time"
)

type Config struct {
	Database struct {
		User    string
		Name    string
		SSLMode string
	}
	GithubAPIKey string
}

type LibrariesDTO []LibraryDTO

type LibraryDTO struct {
	ID             sql.NullInt64
	Name           sql.NullString
	Link           sql.NullString
	Stars          sql.NullInt64
	LastCommitDate sql.NullString
	Followers      sql.NullInt64
	Description    sql.NullString
}

type Libraries []Library

type Library struct {
	ID             int
	Name           string
	Link           string
	Stars          int
	LastCommitDate time.Time
	Followers      int
	Description    string
}

func (libs *Libraries) Has() bool {
	return len(*libs) == 0
}
