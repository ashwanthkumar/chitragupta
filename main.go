package main

import (
	"database/sql"
	"fmt"
	"os"

	getter "github.com/hashicorp/go-getter"
	_ "github.com/mattn/go-sqlite3"
	homedir "github.com/mitchellh/go-homedir"
)

// Constants used for the app
const (
	DBName = "chitragupta.db"
	// TODO - Consider making this configurable and it can default to User's HOME directory if not specified
	DefaultProfileLocation = ".chitragupta"
	InitDBDownloadURL      = "https://github.com/ashwanthkumar/chitragupta/raw/master/chitragupta.init.db"
)

func main() {
	// TODO - Move the following section into a separate file along with the constants
	// Start -- Checking if an existing DB is present in the home directory, if not download the default one from our repo
	homeDir, err := homedir.Dir()
	panicOnErr(err)

	defaultDB := fmt.Sprintf("%s%c%s%c%s", homeDir, os.PathSeparator, DefaultProfileLocation, os.PathSeparator, "chitragupta.db")
	appBase, err := homedir.Expand(fmt.Sprintf("~/%s/", DefaultProfileLocation))
	panicOnErr(err)
	err = os.MkdirAll(appBase, os.ModePerm)
	panicOnErr(err)

	// check if a DB is already present, if not download it from Github
	_, err = os.Stat(defaultDB)
	if err != nil {
		fmt.Printf("%s is not present, downloading the default db from %s\n", defaultDB, InitDBDownloadURL)
		// path is not present, download the Init DB
		err = getter.GetFile(defaultDB, InitDBDownloadURL)
		panicOnErr(err)
		fmt.Printf("Successfully downloaded from %s to %s\n", InitDBDownloadURL, defaultDB)
	}
	// End -- Checking if an existing DB is present in the home directory, if not download the default one from our repo

	db, err := sql.Open("sqlite3", defaultDB)
	panicOnErr(err)
	defer db.Close()

	// stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS userinfo (
	// 	username VARCHAR(64) NULL,
	// 	departname VARCHAR(64) NULL,
	// 	created DATE NULL
	//   );`)
	// panicOnErr(err)

	// _, err = stmt.Exec()
	// panicOnErr(err)

	// // insert
	// stmt, err = db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	// panicOnErr(err)

	// _, err = stmt.Exec("astaxie", "研发部门", "2012-12-09")
	// panicOnErr(err)

}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
