package main

import (
	"testing"
	"unicode"
)

const greeting = "Salut, ça va °-) Κοστίζει €50"

// For å få testen til å bli rett, bruk if-stening og husk fra 129 og til 250 elns er extended.
func TestGreetingExtendedASCII(t *testing.T) {
	for i := 0; i < len(greeting); i++ {
		if greeting[i] <= unicode.MaxASCII {
			t.Errorf("%q is not extended ASCII", greeting[i])
		}
	}
}
