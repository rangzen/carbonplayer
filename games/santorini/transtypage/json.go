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

package transtypage

import (
	"encoding/json"

	"github.com/rangzen/carbonplayer/games/santorini"
)

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

func FromCommand(command Command) (CarbonPlayer, santorini.Node) {
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

func ToCommand(cp CarbonPlayer, n santorini.Node) Command {
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

func JSONErrorMessage(message string) []byte {
	errorMessage := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}
	j, _ := json.Marshal(&errorMessage)
	return j
}
