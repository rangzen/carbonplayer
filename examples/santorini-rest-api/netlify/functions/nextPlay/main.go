package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	stdr.SetVerbosity(3)
	logger := stdr.New(nil)
	g := santorini.NewGame()
	p := santorini.NewPlayer()
	c := Command{}
	err := json.Unmarshal([]byte(request.Body), &c)
	if err != nil {
		logger.Error(err, "unmarshalling body", "body", request.Body)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       JSONErrorMessage("unmarshalling body"),
		}, err
	}
	cp, n := fromCommand(c)
	logger.Info("confReceived", "decision", cp.Decision, "maxPlies", cp.MaxPlies)
	logger.Info("moveReceived", "santorini.Node", n.String())
	d := decision.NewMinimax(logger, cp.MaxPlies)
	nextMove := d.NextMove(g, p, &n)
	logger.Info("moveCalculated", "santorini.Node", nextMove.String(), "score", nextMove.Score())
	nextGs := toCommand(cp, *(nextMove.(*santorini.Node)))
	j, err := json.Marshal(&nextGs)
	if err != nil {
		logger.Error(err, "marshalling response", "body", nextGs)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       JSONErrorMessage("marshalling response"),
		}, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(j),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}

func JSONErrorMessage(message string) string {
	errorMessage := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}
	j, _ := json.Marshal(&errorMessage)
	return string(j)
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameState struct {
	TurnOf  int         `json:"turn_of"`
	Workers [4]Position `json:"workers"`
	Board   [5][5]int   `json:"board"`
}

type CarbonPlayer struct {
	Decision string `json:"decision"`
	MaxPlies int    `json:"max_plies"`
}

type Command struct {
	CarbonPlayer CarbonPlayer `json:"carbon_player"`
	GameState    GameState    `json:"game_state"`
}

func fromCommand(command Command) (CarbonPlayer, santorini.Node) {
	cp := CarbonPlayer{
		Decision: command.CarbonPlayer.Decision,
		MaxPlies: command.CarbonPlayer.MaxPlies,
	}
	node := santorini.Node{
		TurnOf: command.GameState.TurnOf - 1,
		Worker: [4]santorini.Position{
			{command.GameState.Workers[0].X - 1, command.GameState.Workers[0].Y - 1},
			{command.GameState.Workers[1].X - 1, command.GameState.Workers[1].Y - 1},
			{command.GameState.Workers[2].X - 1, command.GameState.Workers[2].Y - 1},
			{command.GameState.Workers[3].X - 1, command.GameState.Workers[3].Y - 1},
		},
		Board: command.GameState.Board,
	}
	return cp, node
}

func toCommand(cp CarbonPlayer, n santorini.Node) Command {
	gs := GameState{
		TurnOf: n.TurnOf + 1,
		Workers: [4]Position{
			{X: n.Worker[0][0] + 1, Y: n.Worker[0][1] + 1},
			{X: n.Worker[1][0] + 1, Y: n.Worker[1][1] + 1},
			{X: n.Worker[2][0] + 1, Y: n.Worker[2][1] + 1},
			{X: n.Worker[3][0] + 1, Y: n.Worker[3][1] + 1},
		},
		Board: n.Board,
	}
	return Command{
		CarbonPlayer: cp,
		GameState:    gs,
	}
}
