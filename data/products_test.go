package data

import "testing"

func TestCheksValidation(t *testing.T) {
	p := &Product{
		Name:  "water",
		Price: 10,
		SKU:   "abd-asd-asd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
