package asset_manager

func sqlEscape(s string) string {
	res := ""
	for _, r := range s {
		if r == '"' {
			res += string('"')
		}
		res += string(r)
	}
	return res
}

func sqlNullString(s *string) string {
	if s == nil {
		return "NULL"
	} else {
		return `"` + *s + `"`
	}
}
