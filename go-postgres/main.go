package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"golang.org/x/sync/errgroup"

	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		return errors.New("POSTGRES_URL is empty")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	g := new(errgroup.Group)
	g.Go(func() error {
		return doUpsert(db, "MSFT", 100, 1000)
	})
	g.Go(func() error {
		return doUpsert(db, "GOOG", 100, 1000)
	})
	g.Go(func() error {
		return doUpsert(db, "AMZN", 100, 1000)
	})
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func doUpsert(db *sql.DB, key string, size int, waitMs int) error {
	i := 0
	for {
		if i >= size {
			i = 0
		}
		value := time.Now()
		log.Printf("| %v | %v | %v |\n", i, key, value)
		_, err := db.Exec("insert into gopg.kv (id,key,value) values ($1,$2,$3) on conflict on constraint pk do update set key=$2, value=$3;", i, key, value)
		if err != nil {
			return err
		}
		i++
		time.Sleep(time.Duration(waitMs) * time.Millisecond)
	}
	return nil
}
