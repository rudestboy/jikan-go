package main

import (
	"fmt"

	"github.com/rudestboy/jikan-go/pkg/jikan"

	pg "github.com/go-pg/pg/v9"
)

type Anime struct {
	MalID int
	Title string
}

func main() {
	client := jikan.New(nil)

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "",
	})
	defer db.Close()

	if err := db.Insert(&Anime{MalID: 40, Title: "not real lol"}); err != nil {
		panic(err)
	}

	anime, err := client.GetAnime(4983)
	if err != nil {
		panic(err)
	}

	fmt.Println(anime)
}
