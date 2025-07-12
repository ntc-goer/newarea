package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func Connect() (*sql.DB, error) {
	client, err := sql.Open("sqlite3", "./../../db.sqlite")
	if err != nil {
		logrus.Errorf("connect err %v", err)
		return nil, err
	}
	if err := client.Ping(); err != nil {
		logrus.Errorf("ping err %v", err)
		return nil, err
	}
	return client, nil
}
