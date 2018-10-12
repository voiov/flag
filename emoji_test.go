package flag

import "testing"

const (
	OK    = "✅ PASS"
	FAILD = "❗️ FAILD"
)

func TestEmoji(t *testing.T) {
	cn := Emoji("cn")
	if cn == "🇨🇳" && m[Emoji("cn")][0] == "China" && m[Emoji("cn")][1] == "中国" {
		t.Log(OK)
	} else {
		t.Error(FAILD)
	}

}
