package main

import (
	"fmt"

	"github.com/noverde/posline"
)

// Proposal reutns ok
type Proposal struct {
	Name     string  `posline:"10,padleft,zerofill"`
	Address  string  `posline:"25"`
	Money    int     `posline:"12,padleft,zerofill"`
	Interest float32 `posline:"10,zerofill,nofp"`
}

func main() {
	p := Proposal{
		Name:    "32",
		Address: "Avenida Ariçaba 58",
		Money:   30,
	}

	p2 := Proposal{
		Name:     "Maria",
		Address:  "Avenida OOOO8",
		Money:    45,
		Interest: 12.234,
	}

	a := []Proposal{p, p2}

	// posline.MarshalStruct(p)
	line, err := posline.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(line)
}
