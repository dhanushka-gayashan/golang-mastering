package main

import (
	"sort"
	"strings"
)

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
	}

	var str strings.Builder
	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}


// To make the list sortable, we need to implement all 3 methods in sort interface
func (l list) Len() int           {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].title < l[j].title
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}


//===========================================
// Implement methods for Sort by Release Date
//===========================================

// byRelease sorts by product release dates.
type byRelease struct {
	// byRelease embeds `list` and reuses list's Len and Swap methods.
	list
}

// Less takes priority over the Less method of the `list`.
// `sort.Sort` will first call this method instead of the `list`'s Less method.
func (br byRelease) Less(i, j int) bool {
	// `Before()` accepts a `time.Time` but `released` is not `time.Time`.
	return br.list[i].released.Before(br.list[j].released.Time)
}

/*
Anonymous embedding means auto-forwarding method calls to the embedded value:
func (br byRelease) Len() int      { return br.list.Len() }
func (br byRelease) Swap(i, j int) { br.list.Swap(i, j)   }
*/
func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}

