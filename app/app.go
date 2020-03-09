package app

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"time"
)

type App struct {
	DB           *sql.DB
	Name         string
	HttpClient   *http.Client
	TimeLocation *time.Location
}

func GetFixedTimeZone() *time.Location {
	timeOffset := os.Getenv("TIME_OFFSET")
	if timeOffset == "" {
		timeOffset = "7"
	}
	nTimeOffset, _ := strconv.Atoi(timeOffset)
	return time.FixedZone("", nTimeOffset*60*60)
}
