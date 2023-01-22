package shortener

import (
	"testing"
)

func TestShortener_Short(t *testing.T) {
	type args struct {
		link string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{},
			want: "0",
		},
		{
			name: "Good string",
			args: args{
				link: "LINK",
			},
			want: "4",
		},
		{
			name: "String with symbols",
			args: args{
				link: " LINK $%\n #\tF ",
			},
			want: "14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shortener{}
			if got := s.Short(tt.args.link); got != tt.want {
				t.Errorf("Short() = %v, want %v", got, tt.want)
			}
		})
	}
}
