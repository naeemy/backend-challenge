package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func solution3() {
	menu := make(map[string]int)
	texts := parseAPI("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")

	re := regexp.MustCompile(`(\s|,|\.)+`)
	cleanTexts := re.ReplaceAllString(texts, ` `)

	for _, word := range strings.Split(cleanTexts, " ") {
		if len(word) > 0 {
			if menu[word] == 0 {
				menu[word] = 1
			} else {
				menu[word] += 1
			}
		}
	}

	jsonMenu, _ := json.MarshalIndent(menu, "", "")

	http.HandleFunc("/summary", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(jsonMenu))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func parseAPI(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(content)
}
