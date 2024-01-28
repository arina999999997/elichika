package user_content

func OverflowCheckedAdd(current *int32, added int32) bool {
	res := int64(*current)
	res += int64(added)
	if (res < 0) || res >= (1<<31) {
		return false
	}
	*current = int32(res)
	return true
}
