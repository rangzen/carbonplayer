package carbonplayer

// Game represents the rules of a game.
type Game interface {
	// Name gives the name of the game.
	Name() string
	// Initial returns a node in an initial state to start a game.
	Initial() Node
	// PossibleChildren gives all the possible Node from this Node.
	PossibleChildren(n Node) []Node
	// IsFinal returns true if the game is finished.
	IsFinal(n Node) bool
	// WinnerIndex returns true if the game have a winner.
	WinnerIndex(n Node) int
}

type Decision interface {
	// NextMove calculates the next move from a Node giving a Game and a Player.
	NextMove(g Game, p Player, n Node) Node
}

// Node represents a board position.
type Node interface {
	// Score returns the evaluation score given by the creator.
	Score() float64
	// SetScore updates the score.
	SetScore(s float64)
	// Base64 represents a state with a string.
	// Can be used for memoization, communication and import/export.
	Base64() string
}

// Player represents reactions to a game position.
type Player interface {
	// Score gives a score to a position.
	// For minimax, score should be positive. E.g.: win=100, draw=0 and position=]0, 1]
	Score(g Game, s Node) float64
}

// PlayerGA represents a Genetic Algorithm player.
type PlayerGA interface {
	// Coef returns all the actual coefficients.
	Coef() []float64
}

// Mutator represents a possible mutation on the PlayerGA.Coef.
type Mutator interface {
	// Mutate creates a new PlayerGA mutated from the input PlayerGA
	Mutate(input PlayerGA) PlayerGA
}
