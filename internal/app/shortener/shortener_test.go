package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("simple create", func(t *testing.T) {
		s := &shortener{logger: nil}
		newS := New(nil)
		assert.Equal(t, s, newS)
	})
}

func Test_shortener_Short(t *testing.T) {
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
			s := &shortener{}
			if got := s.Short(tt.args.link); got != tt.want {
				t.Errorf("Short() = %v, want %v", got, tt.want)
			}
		})
	}
}
