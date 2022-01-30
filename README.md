# carbonplayer Framework

Framework for game simulation with minimax and genetic algorithm.

## Features

### minimax

The main algorithm for choosing the next move is [minimax](https://en.wikipedia.org/wiki/Minimax) 

### Genetic Algorithm

## Limitations in simulated games

### 2 players only

You can simulate chess or tic-tac-toe, and 2 players version of Poker.

You **cannot** simulate "Incan Gold". This is due to the actual minimax algorithm implementation.

## Roadmap

* [ ] Control Max Plies value and then update interfaces with error return.
* [ ] Context to cancel scoring with a time limit.
* [ ] Test memoization of Nodes.
* [ ] Having a HTML GUI for Santorini with the REST API Server.
* [ ] Implement a Decision for alpha-beta pruning.
* [ ] Re implement the Tournament for GA Player (already done for Santorini in another project).
* [ ] Once Interfaces are stable, extract each game to its own repository.
