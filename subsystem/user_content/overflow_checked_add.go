package user_content

func OverflowCheckedAdd(current *int32, added *int32) {
	res := int64(*current)
	res += int64(*added)
	if res >= (1 << 31) {
		if res >= (1 << 31) {
			*added = int32(res - (1<<31 - 1))
			*current = 1<<31 - 1
		}
	} else if res > 0 {
		*current = int32(res)
		*added = 0
	} else {
		// TODO(resource): Maybe this should panic or put user in debt but let's not think about it now
		*current = 0
		*added = 0
	}
}
