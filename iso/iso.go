package iso

import (
	"bytes"
	"fmt"
)

const asciiConst = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f" +
	"€ ‚ƒ„…†‡ˆ‰Š‹Œ Ž  ‘’“”•–—˜™š›œ žŸ " +
	"¡¢£¤¥¦§¨©ª«¬ ®¯°±²³´µ¶·¸¹º»¼½¾¿À" +
	"ÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßà" +
	"áâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿ"

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

//€ ‚ƒ„…†‡ˆ‰Š‹Œ Ž  ‘’“”•–—˜™š
// IterateOverASCIIStringLiteral tar en "string literal" som INN-data

// GenerateExtendedASCIIStringLiteral ...
func GenerateExtendedASCIIStringLiteral() string {
	var buffer bytes.Buffer
	for i := '\u0008'; i <= '\u00FF'; i++ {
		buffer.WriteString(fmt.Sprintf("%c", i))
	}
	return buffer.String()
}

// IterateOverASCIIStringLiterlal...
// Forskjell fra oppgave 1 er at jeg brukte %c for å fjerne unødvendige tegn.
func IterateOverASCIIStringLiteral(sl string) {
	for i := 0; i < len(asciiConst); i++ {
		fmt.Printf("%10X %10c %10b \n", asciiConst[i], asciiConst[i], asciiConst[i])

	}
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {
	greeting := "Salut, ça va °-) Κοστίζει €50"
	for _, c := range []byte(greeting) {
		fmt.Printf("%c", c)
	}
	return greeting
}
