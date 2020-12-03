package nodeset

import (
	"testing"

	"github.com/yourbasic/bit"
)

func Test_bitset_String(t *testing.T) {
	type fields struct {
		Set     *bit.Set
		Padding uint
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"SingleSet", fields{bit.New(1), 0}, "1"},
		{"BasicRange", fields{bit.New(0, 1, 2, 3, 4), 0}, "[0-4]"},
		{"Ranges", fields{bit.New(0, 1, 2, 10, 11, 12, 13), 0}, "[0-2,10-13]"},
		{"Diffs", fields{bit.New(2, 10, 12, 130), 0}, "[2,10,12,130]"},
		{"BasicRangePadding", fields{bit.New(0, 1, 2, 3, 4), 3}, "[0000-0004]"},
		{"RangesPadding", fields{bit.New(0, 1, 2, 10, 11, 12, 13), 1}, "[00-02,10-13]"},
		{"DiffsPadding", fields{bit.New(2, 10, 12, 1300), 2}, "[002,010,012,1300]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &bitset{
				Set:     tt.fields.Set,
				padding: tt.fields.Padding,
			}
			if got := bs.String(); got != tt.want {
				t.Errorf("bitset.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSet_String(t *testing.T) {
	type fields struct {
		pattern string
		sets    []*bitset
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"SingleSet", fields{"node%s", []*bitset{{bit.New(3), 0}}}, "node3"},
		{"SingleSetPadding", fields{"node%s", []*bitset{{bit.New(3), 3}}}, "node0003"},
		{"BasicRange", fields{"node%s", []*bitset{{bit.New(0, 1, 2, 3), 0}}}, "node[0-3]"},
		{"MultiRanges", fields{"node%s-mpi%s", []*bitset{{bit.New(0, 1, 2, 3), 0}, {bit.New(10, 16, 17, 18, 19, 20), 1}}}, "node[0-3]-mpi[10,16-20]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &NodeSet{
				pattern: tt.fields.pattern,
				sets:    tt.fields.sets,
			}
			if got := ns.String(); got != tt.want {
				t.Errorf("NodeSet.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
