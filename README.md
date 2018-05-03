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
    FirstName string `posline:"10"`
    SurName   string `posline:"10"`
    Age       int    `posline:"2"`
    City      string `posline:"15"`
    Country   string `posline:"2"
}
```

Now, to generate the previous flat line, we just need to marshal it:

```golang
  john := Person{
      FirstName: "JOHN",
      SurName:   "DOE",
      Age:       23,
      City:      "São Paulo",
      Country:   "BR",
  }

  flat, err := posline.Marshal(john)

  if err != nil {
      // handle error
  }

  fmt.Println(flat) // => JOHN      DOE       23São Paulo      BR
```

## Modifiers

To help with common tasks when dealing with positional flat files, this lib includes some `modifiers`.

The modifiers should be passed as tag options.

### leftpad
This modifier pads the text to the left to fill the blank space.


```golang
type Person struct {
    FirstName string `posline:"10,leftpad"`
    Country   string `posline:"2"`
}
```

This will pad blank spaces to the left of string JOHN:

```
      JOHNBR
```

### zerofill
This modifier fills the empty space with 0 instead of blank space(" ")


```golang
type Person struct {
    Age       string `posline:"10,zerofill"`
    Country   string `posline:"2"`
}
```

Will produce:

```
2300000000BR
```
### nofp
The nofp (No Floating Point) modifier removes the "." when dealing with float32 or float64 values.
Example:

```golang
type Bank struct {
    Owner   string   `posline:"10"`
    Balance float32  `posline:"10,nofp"`
}

func main() {
	bank := Bank{
		Owner:   "JOHN",
		Balance: 123.45,
	}

	flat, err := posline.Marshal(bank)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

    fmt.Println(flat)
}

```
This example will produce:

```
JOHN      12345
```


**Note:** The nofp modifier will round the float number to 2 decimal places. 

## Multiple Modifiers

You can use multiple modifier as the following example:

```golang
type Bank struct {
    Owner   string   `posline:"10"`
    Balance float32  `posline:"10,leftpad,zerofill,nofp"`
}
```

Using the previous example, will output:

```
JOHN      0000012345
```