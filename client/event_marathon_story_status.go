package client

import (
	"elichika/generic"
)

type EventMarathonStoryStatus struct {
	ReadStoryNumber int32                            `json:"read_story_number"`
	Stories         generic.List[EventMarathonStory] `json:"stories"`
}
