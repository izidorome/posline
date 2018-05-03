package posline_test

import (
	"testing"

	"github.com/noverde/posline"
)

func TestMarshal(t *testing.T) {
	type Person struct {
		Name string `posline:"5"`
		Age  int    `posline:"2"`
	}

	scenarios := []struct {
		description string
		in          []Person
		expected    string
	}{
		{
			description: "it encodes single struct",
			in: []Person{
				Person{
					Name: "John",
					Age:  20,
				},
			},
			expected: "John 20",
		},
		{
			description: "it encodes multiple structs",
			in: []Person{
				Person{
					Name: "John",
					Age:  20,
				},
				Person{
					Name: "Doe",
					Age:  40,
				},
				Person{
					Name: "Baron",
					Age:  25,
				},
			},
			expected: "John 20\nDoe  40\nBaron25",
		},
		{
			description: "it limits to tag size",
			in: []Person{
				Person{
					Name: "Elizabeth Taylor",
					Age:  69,
				},
			},
			expected: "Eliza69",
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := posline.Marshal(tt.in)

			if tt.expected != result {
				t.Errorf("exp: %s got: %s ", tt.expected, result)
			}
		})
	}

}

func TestMarshalWithOptions(t *testing.T) {
	type Bank struct {
		Owner           string  `posline:"10,padleft"`
		Account         int     `posline:"10,padleft,zerofill"`
		Balance         float32 `posline:"9,padleft,zerofill,nofp"`
		LastTransaction float32 `posline:"9"`
	}

	scenarios := []struct {
		description string
		in          []Bank
		expected    string
	}{
		{
			description: "it modifies output based on tag options",
			in: []Bank{
				Bank{
					Owner:           "John",
					Account:         123,
					Balance:         32.12,
					LastTransaction: 10.19,
				},
			},
			expected: "      John000000012300000321210.19    ",
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := posline.Marshal(tt.in)

			if tt.expected != result {
				t.Errorf("exp: %s got: %s ", tt.expected, result)
			}
		})
	}
}

func TestMarshalErrors(t *testing.T) {
	type Wrong struct {
		Name string `posline:"notnum"`
	}

	w := Wrong{
		Name: "Joseph",
	}

	_, err := posline.Marshal(w)

	if err != posline.ErrInvalidSize {
		t.Errorf("does not raised ErrInvalidSize")
	}
}
