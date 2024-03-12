package user_mission

import (
// "elichika/enum"
)

// provide 2 types of update:
// - addition: used for things that get added to
// - maximum: maximum is used for things that we must be getting the max of
// the rule is that maximum can be updated with a lower value (i.e. max voltage), while addition can only get non negative update
// - so things like rank or event point can be updated using max, but they never go down, so it make more sense to use addition
// maybe we should provide some otherway to choose the type of update, if necessary

var (
	isConditionCountSum = map[int32]bool{}
)

func init() {
}

func updateCountByConditionType(conditionType int32, currentCount *int32, newCount int32) {

}
