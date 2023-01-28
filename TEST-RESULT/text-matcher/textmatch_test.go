package textmatcher_test

import (
	textmatcher "go-test/go-leetcode/text-matcher"
	"testing"
)

func TestMatcher(t *testing.T) {
	type args struct {
		textone string
		texttwo string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TEST 1",
			args: args{
				textone: "abcd",
				texttwo: "abcd",
			},
			want: true,
		},
		{
			name: "TEST 2",
			args: args{
				textone: "abcdef",
				texttwo: "abcd",
			},
			want: false,
		},
		{
			name: "TEST 3",
			args: args{
				textone: "abcde",
				texttwo: "abcd",
			},
			want: true,
		},
		{
			name: "TEST 4",
			args: args{
				textone: "abcd",
				texttwo: "abbd",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := textmatcher.Matcher(tt.args.textone, tt.args.texttwo); got != tt.want {
				t.Errorf("Matcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
