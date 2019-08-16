package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string // url
}

func GetMaxNumber() (int, error) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("can't get main page: %s", resp.Status)
	}
	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}
	return comic.Num, nil
}

func GetDescription(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.Status)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}
