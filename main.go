package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type XKCDComic struct {
	Url         string `json:"img"`
	Description string `json:"transcript"`
	Index       uint   `json:"num"`
}

func getXKCDInfo(index uint) XKCDComic {
	var url = fmt.Sprintf("https://xkcd.com/%d/info.0.json", index)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error : %s", err)
	}
	var result XKCDComic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Println("Error : %s", err)
	}
	return result
}

func download() {
	var files []XKCDComic

	for i := uint(1); i < 2530; i++ {
		newComic := getXKCDInfo(i)
		files = append(files, newComic)
	}

	file, _ := json.MarshalIndent(files, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
}

func search(keyword string) {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		defer jsonFile.Close()
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var comics []XKCDComic
	json.Unmarshal(byteValue, &comics)

	var results []XKCDComic
	for _, comic := range comics {
		if strings.Contains(
			strings.ToLower(comic.Description),
			strings.ToLower(keyword)) {
			results = append(results, comic)
		}
	}

	for _, result := range results {
		fmt.Println(result.Index)
		fmt.Println(result.Url)
		fmt.Println("")
		fmt.Println(result.Description)
		fmt.Println("-------------------------------------------------")
	}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func main() {
	if !Exists("data.json") {
		fmt.Println("Datajson not found, let's quickly download it")
		download()
	}
	wordPtr := flag.String("s", "android", "What word do you want to (s)earch for in the vast realm of XKCD")
	flag.Parse()
	search(*wordPtr)
}
