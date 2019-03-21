package iso

import (
	"fmt"
)

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
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


func IterateOverASCIIStringLiteral(sl string) {
	for i:= 0; i < len(asciiConst); i++ {
		fmt.Printf("%X %q %b ", ascii[i], ascii[i], ascii[i])

	}
}

// GreetingExtendedASCII returnerer en tekst-streng i 
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {
	return ""
}
