package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/rangzen/carbonplayer/games/santorini/transtypage"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

func main() {
	stdr.SetVerbosity(3)
	logger := stdr.New(nil)

	http.Handle("/", handleRoot(logger))
	http.Handle("/nextPlay", handleNextPlay(logger))

	logger.Info("Carbon Player - Santorini REST API Server running on localhost:10842")
	log.Fatal(http.ListenAndServe(":10842", nil))
}

func handleRoot(logger logr.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("handling root")
		w.Write([]byte("Welcome to the Carbon Player - Santorini REST API server!"))
	}
}

func handleNextPlay(logger logr.Logger) http.HandlerFunc {
	g := santorini.NewGame()
	p := santorini.NewPlayer()
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			logger.Info("duration", "time", time.Since(start))
		}(time.Now())
		logger.Info("handling nextPlay")
		w.Header().Add("Content-Type", "application/json")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(err, "reading body", "body", string(body))
			w.Write(transtypage.JSONErrorMessage("reading body"))
			return
		}
		c := transtypage.Command{}
		err = json.Unmarshal(body, &c)
		if err != nil {
			logger.Error(err, "unmarshalling body", "body", string(body))
			w.Write(transtypage.JSONErrorMessage("unmarshalling body"))
			return
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
			w.Write(transtypage.JSONErrorMessage("marshalling response"))
			return
		}
		w.Write(j)
	}
}
