package flag

const (
	//emojiA uint32 =  0x1F1E6  OR like follow
	//emojiA uint32 =  '\U0001F1E6'
	//emojiA uint32 =  '🇦'
	emojiA uint32 = '🇦'
)

// Emoji recive 2 Letter which in ISO 3166-1 alpha-2 and return a emoji flag
func Emoji(code string) string {
	if len(code) != 2 {
		return ""
	}

	for _, v := range code {
		if v < 'A' || v > 'z' {
			return ""
		}
		if v > 'Z' && v < 'a' {
			return ""
		}

	}
	ej := string(emojiA+uint32(code[0]|0x20-byte('a'))) + string(emojiA+uint32(code[1]|0x20-byte('a')))
	return ej
}
