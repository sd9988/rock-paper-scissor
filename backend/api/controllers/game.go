package controllers

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"rps/api/responses"
	"rps/api/utils"
)

type Choices struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	RandomNumber int `json:"random_number"`
}

type GameResult struct {
	Result         string `json:"result"`
	Player         int    `json:"player"`
	PlayerChoice   string `json:"playerChoice"`
	Computer       int    `json:"computer"`
	ComputerChoice string `json:"computerChoice"`
}

type PlayRequest struct {
	Player int `json:"player"`
}

const (
	API_URL = "https://codechallenge.boohma.com/"
	RANDOM  = "random"
	WIN     = "win"
	LOSE   = "lose"
	TIE     = "tie"
)

var choices = []string{"rock", "spock", "paper", "lizard", "scissors"}

func (server *Server) Health(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, 200, map[string]string{
		"health": "ok",
	})
}

func (server *Server) Choices(w http.ResponseWriter, r *http.Request) {
	var allChoices []Choices
	for id, choice := range choices {
		allChoices = append(allChoices, Choices{
			ID:   id,
			Name: choice,
		})
	}
	responses.JSON(w, 200, allChoices)
}

func (server *Server) Choice(w http.ResponseWriter, r *http.Request) {
	randomNumber, choice := server.GetRandomChoice()
	responses.JSON(w, 200, Choices{
		ID:   randomNumber,
		Name: choice,
	})
}

func (server *Server) Play(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	playerRequest := PlayRequest{}

	err = json.Unmarshal(body, &playerRequest)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	randomNumber, choice := server.GetRandomChoice()

	var result string

	difference := int(math.Abs(float64(randomNumber)-float64(playerRequest.Player))) % 5

	if difference == 1 || difference == 2 {
		result = LOSE
	}

	if difference == 4 || difference == 3 {
		result = WIN
	}

	if difference == 0 {
		result = TIE
	}

	responses.JSON(w, 200, GameResult{
		Result:         result,
		Player:         playerRequest.Player,
		Computer:       randomNumber,
		PlayerChoice:   choices[playerRequest.Player],
		ComputerChoice: choice,
	})

}

func (server *Server) GetRandomChoice() (int, string) {
	randomNumber := utils.RandInt(0, 5)
	choice := choices[randomNumber]
	return randomNumber, choice
}
