package duck

import (
	"bytes"
	"fmt"
)

// Farmer
type FarmerI interface {
	Breed() string
}

type farmer struct {
	ducks []DuckI
}

func NewFarmer(ducks ...DuckI) FarmerI {
	return &farmer{ducks}
}

func (f *farmer) Breed() string {
	var b bytes.Buffer
	for i, d := range f.ducks {
		fmt.Fprintf(&b, "#%d %s\n", i, d.quack())
	}
	return b.String()
}
