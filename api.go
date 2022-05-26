package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	ApiOwner ApiOwner `json:"api_owner"`
	Body     Body     `json:"body"`
}

type ApiOwner struct {
	Author    string `json:"author"`
	Web       string `json:"web"`
	Cafecito  string `json:"cafecito"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Linkedin  string `json:"linkedin"`
	Twitter   string `json:"twitter"`
}

type Body struct {
	Word                string    `json:"Word"`
	Definition          string    `json:"Definition"`
	Author              string    `json:"Author"`
	ErrorMessage        string    `json:"ErrorMessage"`
	EncodingWebName     string    `json:"EncodingWebName"`
	WordOrigin          string    `json:"WordOrigin"`
	UrlDefinitionSource string    `json:"UrlDefinitionSource"`
	Types               []string  `json:"Types"`
	Urls                Urls      `json:"urls"`
	DefinitionMD        string    `json:"DefinitionMD"`
	Related             []Related `json:"Related"`
}

type Urls struct {
	ApiUrl     string `json:"api_url"`
	WebUrl     string `json:"web_url"`
	Wiktionary string `json:"wiktionary"`
	Wikipedia  string `json:"wikipedia"`
	Thumbnail  string `json:"thumbnail"`
	Image      string `json:"image"`
}

type Related struct {
	Word string `json:"Word"`
	Urls Urls   `json:"urls"`
}

func generateWordApi() string {

	response, err := http.Get("https://palabras-aleatorias-public-api.herokuapp.com/random")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		log.Fatal(err)
	}

	return strings.ToUpper(responseObject.Body.Word)
}
