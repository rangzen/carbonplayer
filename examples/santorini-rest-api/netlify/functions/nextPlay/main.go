/*
 * Copyright 2022 Cédric L'HOMME
 *
 * This file is part of the Carbon Player Framework.
 *
 * The Carbon Player Framework is free software: you can redistribute it and/or modify it under the terms of
 * the GNU General Public License as published by the Free Software Foundation, either version 3 of the License,
 * or (at your option) any later version.
 *
 * The Carbon Player Framework is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
 * without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with the Carbon Player Framework.
 * If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/rangzen/carbonplayer/games/santorini/transtypage"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	stdr.SetVerbosity(3)
	logger := stdr.New(nil)
	g := santorini.NewGame()
	p := santorini.NewPlayer()
	c := transtypage.Command{}
	err := json.Unmarshal([]byte(request.Body), &c)
	if err != nil {
		logger.Error(err, "unmarshalling body", "body", request.Body)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(transtypage.JSONErrorMessage("unmarshalling body")),
		}, err
	}
	cp, n := transtypage.FromCommand(c)
	logger.Info("confReceived", "decision", cp.Decision, "maxPlies", cp.MaxPlies)
	logger.Info("moveReceived", "santorini.Node", n.String())
	d := decision.NewMinimax(logger, cp.MaxPlies)
	nextMove := d.NextMove(g, p, &n)
	logger.Info("moveCalculated", "santorini.Node", nextMove.String(), "score", nextMove.Score())
	nextGs := transtypage.ToCommand(cp, *(nextMove.(*santorini.Node)))
	j, err := json.Marshal(&nextGs)
	if err != nil {
		logger.Error(err, "marshalling response", "body", nextGs)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(transtypage.JSONErrorMessage("marshalling response")),
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
