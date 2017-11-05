package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SearchContent(title string) map[string]string {

	var movielist map[string]*json.RawMessage
	var searches = make([]*json.RawMessage, 2)
	var target map[string]string

	req, err := http.NewRequest("GET", "http://omdbapi.com", nil)
	if err != nil {
		panic(err)
	}

	//API for searching movie
	q := req.URL.Query()
	q.Add("s", title)
	q.Add("apikey", "98bba0d3")
	req.URL.RawQuery = q.Encode()

	resp, _ := http.Get(req.URL.String())

	l_body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(l_body, &movielist)

	json.Unmarshal(*movielist["Search"], &searches)

	//Top search
	json.Unmarshal(*searches[0], &target)

	imdbID := target["imdbID"] //IMDb id for metadata

	req, err = http.NewRequest("GET", "http://omdbapi.com", nil)
	if err != nil {
		panic(err)
	}

	//API for getting metadata
	q = req.URL.Query()
	q.Add("i", imdbID)
	q.Add("apikey", "98bba0d3")
	req.URL.RawQuery = q.Encode()

	resp, _ = http.Get(req.URL.String())
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &target)
	fmt.Printf("%v", target)

	return target
}
