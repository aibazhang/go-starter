package duck

import (
	"bytes"
	"fmt"
	"reflect"
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
		fmt.Fprintf(&b, "#%d %s %T\n", i, d.quack(), d)
		fmt.Println(reflect.TypeOf(d).Elem())
	}
	return b.String()
}
