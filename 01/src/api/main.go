package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Article struct {
	Title       string
	Description string
	Body        string
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article{
		{
			Title:       "テスト記事01",
			Description: "見出し",
			Body:        "本文01",
		},
		{
			Title:       "テスト記事02",
			Description: "見出し02",
			Body:        "ラーメンはおいしい",
		},
	}

	if isCustomHostname(r.Host) {
		articles = []Article{
			{
				Title:       fmt.Sprintf("%sの記事\n", getCustomHostName()),
				Description: "見出し",
				Body:        "本文01",
			},
			{
				Title:       "独自ドメインのテスト記事02",
				Description: "見出し02",
				Body:        "ラーメンはおいしい",
			},
		}
	}

	str, _ := json.Marshal(articles)

	fmt.Fprintf(w, string(str))
}

func isCustomHostname(hostHeader string) bool {
	customHostname := getCustomHostName()
	if customHostname == "" {
		return false
	}

	if hostHeader == customHostname {
		return true
	}

	return false
}

func getCustomHostName() string {
	return os.Getenv("CUSTOM_HOSTNAME")
}

func main() {
	http.HandleFunc("/articles", GetArticles)
	http.ListenAndServe(":8080", nil)
}
