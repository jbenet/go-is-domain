package isdomain

// ExtendedTLDs is a set of additional "TLDs", allowing decentralized name
// systems, like TOR and Namecoin.
var ExtendedTLDs = map[string]bool{
	"BIT":   true,
	"ONION": true,
	"ETH":   true,
	"BBS":   true,
	"CHAN":  true,
	"CYB":   true,
	"DYN":   true,
	"EPIC":  true,
	"GEEK":  true,
	"GOPHER":true,
	"INDY":  true,
	"LIBRE": true,
	"NEO":   true,
	"NULL":  true,
	"O":     true,
	"OSS":   true,
	"OZ":    true,
	"PARODY":true,
	"PIRATE":true,
	"CRYPTO": true, // unstoppabledomains.com
	"ZIL":    true, // unstoppabledomains.com
}
