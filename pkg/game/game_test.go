package game

import (
	"errors"
	"fmt"
	"testing"

	termbox "github.com/nsf/termbox-go"
	"github.com/smartystreets/goconvey/convey"
)

func TestGetPlayerOneDisplayProps(t *testing.T) {
	//termbox.ColorDefault, termbox.ColorBlack
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
