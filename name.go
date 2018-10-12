package flag

//ChineseName recive 2 Letter which in ISO 3166-1 alpha-2 and return it's Chinese name
func ChineseName(code string) string {
	n := Emoji(code)
	if n == "" {
		return ""
	}
	if v, ok := m[n]; ok {
		return v[1]
	}
	return ""
}

//EnglishName recive 2 Letter which in ISO 3166-1 alpha-2 and return it's English name
func EnglishName(code string) string {
	n := Emoji(code)
	if n == "" {
		return ""
	}
	if v, ok := m[n]; ok {
		return v[0]
	}
	return ""
}
