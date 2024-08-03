package info_trigger

import (
	"elichika/router"
)

func init() {
	// the request / response is exactly the same, there's only different in the end point
	router.AddHandler("/", "POST", "/infoTrigger/readEventMarathonResult", read)
}
