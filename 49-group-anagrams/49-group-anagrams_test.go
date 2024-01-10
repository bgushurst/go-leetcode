package leetcode

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
				{name: "test1", args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, want: [][]string{{"eat","tea","ate"},{"tan","nat"},{"bat"}}},
				{name: "test2", args: args{strs: []string{""}}, want: [][]string{{""}}},
				{name: "test3", args: args{strs: []string{"a"}}, want: [][]string{{"a"}}},
				{name: "test4", args: args{strs: []string{"bdddddddddd", "bbbbbbbbbbc"}}, want: [][]string{{"bbbbbbbbbbc"},{"bdddddddddd"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
