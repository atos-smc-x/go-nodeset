package nodeset

import (
	"fmt"
	"strings"

	"github.com/yourbasic/bit"
)

type bitset struct {
	*bit.Set
	padding uint
}

// writeRange appends either "", "a", "a,b" or "a-b" to buf.
func writeRange(buf *strings.Builder, a, b int, padding uint) {
	d := "%d"
	if padding > 0 {
		d = fmt.Sprintf("%%0%dd", padding+1)
	}
	switch {
	case a > b:
		return // Append nothing.
	case a == b:
		fmt.Fprintf(buf, d, a)
		//	case a+1 == b:
		//		fmt.Fprintf(buf, "%d,%d", a, b)
	default:
		fmt.Fprintf(buf, d+"-"+d, a, b)
	}
}

func (bs *bitset) String() string {
	buf := new(strings.Builder)
	if bs.Empty() {
		return ""
	}
	if bs.Size() == 1 {
		i := bs.Next(-1)
		writeRange(buf, i, i, bs.padding)
		return buf.String()
	}
	buf.WriteByte('[')
	a, b := -1, -2 // Keep track of a range a-b of elements.
	first := true
	bs.Visit(func(n int) (skip bool) {
		if n == b+1 {
			b++ // Increase current range from a-b to a-b+1.
			return
		}
		if first && a <= b {
			first = false
		} else if a <= b {
			buf.WriteByte(',')
		}
		writeRange(buf, a, b, bs.padding)
		a, b = n, n // Start new range.
		return
	})
	if !first && a <= b {
		buf.WriteByte(',')
	}
	writeRange(buf, a, b, bs.padding)
	buf.WriteByte(']')
	return buf.String()
}

// A NodeSet represent a
type NodeSet struct {
	pattern string
	sets    []*bitset

	// TODO(loicalbertin) padding
	/*
		padding is comming from the first padded number
		nodeset -r node[4,05,00010]
		node[04-05,10]
	*/
}

func (ns *NodeSet) String() string {
	s := make([]interface{}, len(ns.sets))
	for i := range ns.sets {
		s[i] = ns.sets[i]
	}
	return fmt.Sprintf(ns.pattern, s...)
}
