package stringutil

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type args struct{ slice []string }
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{"1", args{[]string{"1", "1"}}, []string{"1"}},
		{"2", args{[]string{"1", "1", "2"}}, []string{"1", "2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Unique(tt.args.slice); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Unique() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestFindString(t *testing.T) {
	type args struct {
		array []string
		str   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"1", "2"}, "0"}, -1},
		{"2", args{[]string{"1", "2"}, "1"}, 0},
		{"3", args{[]string{"1", "2"}, "3"}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindString(tt.args.array, tt.args.str); got != tt.want {
				t.Errorf("FindString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringIn(t *testing.T) {
	type args struct {
		array []string
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]string{"1", "2"}, "0"}, false},
		{"2", args{[]string{"1", "2"}, "1"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringIn(tt.args.array, tt.args.str); got != tt.want {
				t.Errorf("StringIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"123456"}, "654321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaNumberic(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"test"}, true},
		{"2", args{"123456"}, true},
		{"3", args{"test123456"}, true},
		{"4", args{"123456_"}, false},
		{"5", args{"123456~"}, false},
		{"5", args{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNumberic(tt.args.s); got != tt.want {
				t.Errorf("IsAlphaNumberic() = %v, want %v", got, tt.want)
			}
		})
	}
}
