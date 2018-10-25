package game

type Manager struct {
	playerOneWins int
	playerTwoWins int
}

func NewManager() *Manager {
	return &Manager{}
}

// NewGame instantiates a new game using the current count of previous wins.
func (m *Manager) NewGame(game *Game) *Game {
	game.Cancel()
	if game.Winner == 1 {
		m.playerOneWins++
	} else if game.Winner == 2 {
		m.playerTwoWins++
	}

	return NewGame(game.Width, game.Height, m.playerOneWins, m.playerTwoWins)
}
