package flag

import "testing"

func TestChineseName(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"中国",
			args{code: "cn"},
			"中国",
		},
		{"乌兹别克斯坦",
			args{code: "uz"},
			"乌兹别克斯坦",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChineseName(tt.args.code); got != tt.want {
				t.Errorf("ChineseName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnglishName(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "China",
			args: args{code: "cn"},
			want: "China",
		},
		{"Uzbekistan",
			args{code: "uz"},
			"Uzbekistan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnglishName(tt.args.code); got != tt.want {
				t.Errorf("EnglishName() = %v, want %v", got, tt.want)
			}
		})
	}
}
