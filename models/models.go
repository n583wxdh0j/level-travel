package models

import (
	"time"
)

type Librarieser interface {
	Has() bool
}

type Libraries []Library

type Library struct {
	ID             int
	Repo           string
	Owner          string
	Link           string
	Stars          int
	LastCommitDate time.Time
	Followers      int
	Description    string
	LastUpdate     time.Time
}

func (libs *Libraries) Has() bool {
	return len(*libs) == 0
}

type Config struct {
	Database struct {
		User    string
		Name    string
		SSLMode string
	}
	GithubAPIKey string
}
