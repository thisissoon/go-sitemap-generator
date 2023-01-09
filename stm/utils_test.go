package stm

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeMap(t *testing.T) {
	var src, dst, expect [][]interface{}
	src = [][]interface{}{{"loc", "1"}, {"changefreq", "2"}, {"mobile", true}, {"host", "http://google.com"}}
	dst = [][]interface{}{{"host", "http://example.com"}}
	expect = [][]interface{}{{"loc", "1"}, {"changefreq", "2"}, {"mobile", true}, {"host", "http://google.com"}}

	src = MergeMap(src, dst)

	if !reflect.DeepEqual(src, expect) {
		t.Fatalf("Failed to maps merge: deferrent map \n%#v\n and \n%#v\n", src, expect)
	}
}

func TestURLJoin(t *testing.T) {
	type args struct {
		src   string
		joins []string
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"two join last is empty": {
			args: args{
				src:   "",
				joins: []string{"http://example.com", ""},
			},
			want: "http://example.com",
		},
		"two joins": {
			args: args{
				src:   "",
				joins: []string{"http://example.com", "men"},
			},
			want: "http://example.com/men",
		},
		"three joins": {
			args: args{
				src:   "",
				joins: []string{"http://example.com", "men", "a"},
			},
			want: "http://example.com/men/a",
		},
		"has slash already": {
			args: args{
				src:   "",
				joins: []string{"http://example.com", "men/", "a"},
			},
			want: "http://example.com/men/a",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := URLJoin(tt.args.src, tt.args.joins...)
			assert.Equal(t, tt.want, got)
		})
	}
}
