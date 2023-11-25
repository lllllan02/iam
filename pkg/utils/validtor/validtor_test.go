package validtor

import "testing"

func TestIsEmailValid(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"lllllan@test.com"}, true},
		{"2", args{"test@example"}, false},
		{"3", args{"test@com."}, false},
		{"4", args{"test.com"}, false},
		{"5", args{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailValid(tt.args.email); got != tt.want {
				t.Errorf("IsEmailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
