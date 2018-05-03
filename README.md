# posline

The posline package aims to provide easy serialization of Go structs into positional fixed-width strings. 
These kind of structures are useful when integrating with old system which depends on positional flat files

## Installation

```
import "github.com/noverde/posline"
```

## Usage

Consider the following fixed-width string line:


```
JOHN      DOE       23São Paulo      BR
```

Looking at this file, we can see the following structure:

* `firstname` with 10 fixed columns;
* `surname` with 10 fixed columns;
* `age` with 2 fixed columns;
* `city` with 15 fixed columns;
* `country` with 2 fixed columns.

Let's map it into a Go Struct:

```golang
type Person struct {
    FirstName string `posline:10"`
    SurName   string `posline:10"`
    Age       int    `posline:"2"`
    City      string `posline:"15"`
    Country   string `posline:"2"
}
```

Now, to generate the previous flat line, we just need to marshal it:

```golang

  john := Person{
      FirstName: "JOHN",
      SurName: "DOE",
      Age: 23,
      City: "São Paulo",
      Country: "BR"
  }

  flat, err := posline.Marshal(john)

  if err != nil {
      // handle error
  }

  fmt.Println(flat) // => JOHN      DOE       23São Paulo      BR
```
