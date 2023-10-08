package storage

import (
	"database/sql"
	"urlShortener/internal/link"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(pathToDB string) error {
	var err error
	db, err = sql.Open("sqlite3", "database/db.sqlite")
	if err != nil {
		return err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY,
		short TEXT NOT NULL,
		url TEXT NOT NULL
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func Put(link *link.Link) error {
	id, err := getId(db)
	if err != nil {
		return err
	}

	link.Initialize(id)

	insertSQL := `
	INSERT INTO links (id, short, url) VALUES (?, ?, ?);
	`

	_, err = db.Exec(insertSQL, link.Id, link.ShortUrl, link.Url)
	if err != nil {
		return err
	}

	return nil
}

func GetAll() ([]link.Link, error) {
	selectSQL := `
	SELECT id, short, url FROM links;
	`

	rows, err := db.Query(selectSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []link.Link

	for rows.Next() {
		newLink := link.Link{}
		err = rows.Scan(&newLink.Id, &newLink.ShortUrl, &newLink.Url)

		if err != nil {
			return nil, err
		}

		links = append(links, newLink)
	}

	return links, nil
}

func GetFullUrl(shortPath string) (string, error) {
	selectSQL := `
	SELECT url FROM links WHERE short = ?;
	`

	rows, err := db.Query(selectSQL, shortPath)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var fullUrl string

	rows.Next()
	err = rows.Scan(&fullUrl)
	if err != nil {
		return "", err
	}

	return fullUrl, nil
}

func getId(db *sql.DB) (int, error) {
	var maxID int
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM links").Scan(&maxID)
	if err != nil {
		return 0, err
	}

	return maxID + 1, nil
}
