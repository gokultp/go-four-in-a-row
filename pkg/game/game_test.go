package game

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/nsf/termbox-go"
	"github.com/smartystreets/goconvey/convey"
)

func TestGetPlayerOneDisplayProps(t *testing.T) {
	convey.Convey("Given the player one", t, func() {
		var player int = 1
		convey.Convey("When get the player one display", func() {
			rune, colorPlayer, colorBg := getplayerDisplayProps(player)
			convey.Convey("Should return rune with 'X' , Color 'Red' and Background Color 'Default' ", func() {
				var err error
				if rune != 'X' || colorPlayer != termbox.ColorRed || colorBg != termbox.ColorDefault {
					err = errors.New(fmt.Sprintf("Doesn't returned correct value to player one display, returned rune %v, colorPlayer %v, colorBg %v and expected  rune='X', colorPlayer=termbox.ColorRed and colorBg=termbox.ColorDefault", rune, colorPlayer, colorBg))
				}

				convey.So(err, convey.ShouldBeNil)
			})

		})
	})

}

func TestGame_isWon(t *testing.T) {
	tests := map[string]struct {
		name   string
		game   *Game
		col    int
		row    int
		player int
		state []string
		want   bool
	}{
		"5x4 game vertical 2 item": {
			name: "",
			game: NewGame(5, 4, 0, 0),
			want: false,
			col: 2,
			row: 1,
			player: 1,
			state: []string{
				`00000`,
				`00000`,
				`00100`,
				`00100`,
			},
		},
		"5x4 game vertical 4 item win": {
			name: "",
			game: NewGame(5, 4, 0, 0),
			want: true,
			col: 2,
			row: 0,
			player: 1,
			state: []string{
				`00100`,
				`00100`,
				`00100`,
				`00100`,
			},
		},
		"5x4 game vertical 4 item not win": {
			name: "",
			game: NewGame(5, 4, 0, 0),
			want: false,
			col: 1,
			row: 0,
			player: 1,
			state: []string{
				`00100`,
				`00100`,
				`00100`,
				`00100`,
			},
		},
		"6x4 game horizontal 4 from left": {
			name: "",
			game: NewGame(6, 4, 0, 0),
			want: true,
			col: 3,
			row: 1,
			player: 1,
			state: []string{
				`000000`,
				`111100`,
				`000000`,
				`000000`,
			},
		},
		"6x4 game horizontal 4 from right": {
			name: "",
			game: NewGame(6, 4, 0, 0),
			want: true,
			col: 4,
			row: 1,
			player: 1,
			state: []string{
				`000000`,
				`001111`,
				`000000`,
				`000000`,
			},
		},
		"6x4 game horizontal 4 middle": {
			name: "",
			game: NewGame(6, 4, 0, 0),
			want: true,
			col: 1,
			row: 1,
			player: 1,
			state: []string{
				`000000`,
				`011110`,
				`000000`,
				`000000`,
			},
		},
		"6x4 game horizontal 3 left": {
			name: "",
			game: NewGame(6, 4, 0, 0),
			want: false,
			col: 1,
			row: 1,
			player: 1,
			state: []string{
				`000000`,
				`111000`,
				`000000`,
				`000000`,
			},
		},
		"6x4 game horizontal 3 right": {
			name: "",
			game: NewGame(6, 4, 0, 0),
			want: false,
			col: 1,
			row: 1,
			player: 1,
			state: []string{
				`000000`,
				`000111`,
				`000000`,
				`000000`,
			},
		},
		"6x4 game horizontal 3 middle": {
			name: "",
			game: NewGame(6, 4, 0, 0),
			want: false,
			col: 1,
			row: 1,
			player: 1,
			state: []string{
				`000000`,
				`001110`,
				`000000`,
				`000000`,
			},
		},
		"6x5 game negative diagonal 3 top left": {
			name: "",
			game: NewGame(6, 5, 0, 0),
			want: false,
			col: 3,
			row: 3,
			player: 1,
			state: []string{
				`100000`,
				`010000`,
				`001000`,
				`000000`,
				`000000`,
			},
		},
		"6x5 game negative diagonal 4 top left": {
			name: "",
			game: NewGame(6, 5, 0, 0),
			want: true,
			col: 3,
			row: 3,
			player: 1,
			state: []string{
				`100000`,
				`010000`,
				`001000`,
				`000100`,
				`000000`,
			},
		},
		"6x5 game negative diagonal 4 bottom right": {
			name: "",
			game: NewGame(6, 5, 0, 0),
			want: true,
			col: 5,
			row: 4,
			player: 1,
			state: []string{
				`000000`,
				`001000`,
				`000100`,
				`000010`,
				`000001`,
			},
		},
		"6x6 game negative diagonal 4 middle": {
			name: "",
			game: NewGame(6, 6, 0, 0),
			want: true,
			col: 1,
			row: 1,
			player: 1,
			state: []string{
				`000000`,
				`010000`,
				`001000`,
				`000100`,
				`000010`,
				`000000`,
			},
		},
		"6x5 game positive diagonal 3 top right": {
			name: "",
			game: NewGame(6, 5, 0, 0),
			want: false,
			col: 5,
			row: 0,
			player: 1,
			state: []string{
				`000001`,
				`000010`,
				`000100`,
				`000000`,
				`000000`,
			},
		},
		"6x5 game positive diagonal 4 top right": {
			name: "",
			game: NewGame(6, 5, 0, 0),
			want: true,
			col: 5,
			row: 0,
			player: 1,
			state: []string{
				`000001`,
				`000010`,
				`000100`,
				`001000`,
				`000000`,
			},
		},
		"6x5 game positive diagonal 4 bottom left": {
			name: "",
			game: NewGame(6, 5, 0, 0),
			want: true,
			col: 0,
			row: 4,
			player: 1,
			state: []string{
				`000000`,
				`000100`,
				`001000`,
				`010000`,
				`100000`,
			},
		},
		"6x6 game positive diagonal 4 middle": {
			name: "",
			game: NewGame(6, 6, 0, 0),
			want: true,
			col: 1,
			row: 3,
			player: 1,
			state: []string{
				`000010`,
				`000100`,
				`001000`,
				`010000`,
				`000000`,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := tt.game

			err := applyState(tt.game, tt.state)
			if err != nil {
			    t.Fatalf("failed to apply state: %v", err)
			}

			if got := g.isWon(tt.col, tt.row, tt.player); got != tt.want {
				t.Errorf("Game.isWon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func applyState(g *Game, state []string) error {
	if len(state) > g.Height {
		return fmt.Errorf("got height %d while game height is %d", len(state), g.Height)
	}
	for row, rowStr := range state {
		if len(rowStr) > g.Width {
			return fmt.Errorf("got width %d while game width is %d", len(rowStr), g.Width)
		}

		for col, char := range rowStr {
			player, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse %q at %d:%d: %v", char, row, col, err)
			}
			g.State[row][col] = int(player)
		}
	}

	return nil
}
