package _1_default_implementation

import (
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
