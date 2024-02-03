package live

import (
	"elichika/router"
)

func init() {
	router.AddHandler("/live/finishTutorial", finish) // this is correct
}
