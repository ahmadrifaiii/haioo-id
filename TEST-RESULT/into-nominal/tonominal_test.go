package intonominal_test

import (
	intonominal "go-test/go-leetcode/into-nominal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInToNominal(t *testing.T) {
	type args struct {
		num int
	}
	testcases := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "TEST 4000",
			args: args{
				num: 4000,
			},
			want: map[string]int{
				"Rp. 2.000": 2,
			},
		},
		{
			name: "TEST 175200",
			args: args{
				num: 175200,
			},
			want: map[string]int{
				"Rp. 100.000": 1,
				"Rp. 50.000":  1,
				"Rp. 20.000":  1,
				"Rp. 5.000":   1,
				"Rp. 200":     1,
			},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			got := intonominal.InToNominal(test.args.num)
			assert.Equal(t, test.want, got)
		})
	}
}
