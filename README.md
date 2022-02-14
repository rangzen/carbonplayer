# Carbon Player Framework

The Carbon Player Framework is a personal tool for ideas exploration about 2 players board games simulation with some decision algorithm like minimax and genetic algorithm.  
Currently, the GA part is **not** in this project, I have to backport it from a private repository about Santorini.

Please use it for testing and fun.  
Even if it's not intended to be a professional thing (it's a personal toy), I still value remarks on quality or good practices.
Please contact me, or create an issue, for any subject.

## Features

### minimax

Actually, the main algorithm for choosing the next move is [minimax](https://en.wikipedia.org/wiki/Minimax) 

### Genetic Algorithm

The GA part will be backported from a previous Santorini implementation.

## Games

### Tic Tac Toe

Implemented to test minimax implementation.  

There is only two CLI tools, one to play against CP, one to watch two CP play against each other.

### Santorini

One of my favourite game. The first implementation was with GA because I wanted to find which metrics influence the game.

You can use an HTML/SVG GUI with the REST API or a CLI tool to test the actual Carbon Player.  
Spoiler: he is dumb, but I already sometimes lose…

## Limitations in simulated games

### 2 players only

You can simulate chess or tic-tac-toe, and 2 players version of Poker.

You **cannot** simulate “Incan Gold” for example. This is due to the actual minimax algorithm implementation.

## Road map

* [ ] Control Max Plies value and then update interfaces with error return.
* [ ] Context to cancel scoring with a time limit.
* [ ] Test memoization of Nodes.
* [ ] Having an HTML GUI for Santorini with the REST API Server.
* [ ] Implement a Decision for alpha-beta pruning.
* [ ] Re implement Mutation, Tournament, etc. for GA Player (already done for Santorini in another project).
* [ ] Once Interfaces are stable, extract each game to theirs own repositories.
