package isdomain

// ExtendedTLDs is a set of additional "TLDs", allowing decentralized name
// systems, like TOR and Namecoin.
var ExtendedTLDs = map[string]bool{
	"BIT":    true, // namecoin.org
	"ONION":  true, // torproject.org
	"ETH":    true, // ens.domains
	"CRYPTO": true, // unstoppabledomains.com
	"ZIL":    true, // unstoppabledomains.com
}
