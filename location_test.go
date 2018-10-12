package flag

import (
	"testing"
)

func TestNew(t *testing.T) {
	us, err := New("us")
	if err != nil {
		t.Error(FAILD)
	} else {
		t.Log(OK, us.Emoji(), us.EnglishName(), us.ChineseName())
	}
}

func TestAddNew(t *testing.T) {
	type args struct {
		code  string
		names [2]string
	}
	tests := []struct {
		name string
		args args
	}{{"Vircual Location", args{"XX", [2]string{"XX location", "XX 地区"}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddNew(tt.args.code, tt.args.names)
			if f, _ := New(tt.args.code); f.EnglishName() == tt.args.names[0] {
				t.Log(OK, "Add a new item:", tt.args.code, tt.args.names)
			} else {
				t.Error(FAILD)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{"Vircual Location", args{"XX"}, "XX"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.code); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
			if _, ok := m[Emoji(tt.args.code)]; ok {
				t.Error(FAILD)
			}

		})
	}
}
