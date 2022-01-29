package carbonplayer

// Game represents the rules of a game.
type Game interface {
	// Initial returns a node in an initial state to start a game.
	Initial() Node
	// PossibleChildren gives all the possible Node from this Node.
	PossibleChildren(n Node) []Node
	// IsFinal returns true if the game is finished.
	IsFinal(n Node) bool
	// WinnerIndex returns the index of the winner.
	// Some game use 0 for draw.
	WinnerIndex(n Node) int
}

type Decision interface {
	// NextMove calculates the next move from a Node giving a Game and a Player.
	NextMove(g Game, p Player, n Node) Node
}

// Node represents a board position.
type Node interface {
	// Score returns the actual score of this node.
	// The Score is normally given by a Player.
	Score() float64
	// SetScore updates the score.
	SetScore(s float64)
	// String represents a state with a string.
	// Can be used for memoization, communication and import/export.
	String() string
}

// Player represents reactions to a game position.
type Player interface {
	// Score gives a score to a position.
	// source is true if you are the Player that starts the research.
	// source is equivalent to:
	// Are you the one who have "interest" in the actual scoring?
	Score(g Game, s Node, source bool) float64
}

// Metric represent a value from a Node.
// Normally used by Player to calculate a Score.
// E.g., distance between players, numbers of free movements, etc.
type Metric interface {
	// Value returns the value of a certain Node regarding the represented metric.
	Value(node Node, source bool) float64
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
