package main
import (
	_ "github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/golang-migrate/migrate"
	"log"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	m, err := migrate.New(
		"file://database/migrations",
		"sqlite3:///Users/maxheckel/Go/src/github.com/maxheckel/meditations/db.sqlite")
	if err != nil{
		log.Fatal(err)
	}
	files,_ := ioutil.ReadDir("/Users/maxheckel/Go/src/github.com/maxheckel/meditations/database/migrations")
	num := len(files) / 2
	if len(os.Args) < 2 {
		fmt.Println("Must pass a method")
		return
	}
	if os.Args[1] == "up" {
		m.Steps(num)
	}
	if os.Args[1] == "down"{
		m.Steps(num * -1)
	}

}