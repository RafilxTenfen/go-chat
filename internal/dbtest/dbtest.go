package dbtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // needed to load sqlite3 driver
	"github.com/rhizomplatform/fs"
)

// DBHandler is a handler funcion that accepts a db instance
type DBHandler func(db *gorm.DB)

// WithDB run the specified handler providing a valid sqlite3 database.
func WithDB(handler DBHandler) {
	// Please see https://www.sqlite.org/inmemorydb.html
	//
	// Look for the section "In-memory Databases And Shared Cache" to understand why
	// the url cannot be a simple ":memory:"
	db, err := gorm.Open("sqlite3", "file:memdb?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// run!
	handler(db)
}

// Require executes the specified script dependency into the provided
// database instance.
func Require(db *gorm.DB, deps ...string) {
	solved := make(map[string]struct{})
	for _, dep := range deps {
		if err := solve(solved, db, dep); err != nil {
			panic(err)
		}
	}
}

func solve(state map[string]struct{}, db *gorm.DB, dep string) error {
	// already solved? nothing to do!
	if _, ok := state[dep]; ok {
		return nil
	}

	// solve our dependencies first
	deps := dependencies[dep]
	for _, d := range deps {
		if err := solve(state, db, d); err != nil {
			return err
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := fs.Path(pwd).Parent().Join("/internal/dbtest/scripts/" + dep + ".sql")
	file, err := path.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	script := string(b)
	commands := strings.Split(script, ";")

	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)

		if cmd == "" {
			continue
		}

		if err := db.Exec(cmd).Error; err != nil {
			return fmt.Errorf("error resolving '%s', on file '%s': %v\nCommand: '%s'", dep, file.Name(), err, cmd)
		}
	}

	state[dep] = struct{}{}
	return nil
}
