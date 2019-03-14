package isdomain

import (
	"fmt"
	"strings"
)

// IsICANNTLD returns whether the given string is a TLD (Top Level Domain),
// according to ICANN. Well, really according to the TLDs listed in this
// package.
func IsICANNTLD(s string) bool {
	s = strings.ToUpper(s)
	_, found := TLDs[s]
	return found
}

// IsExtendedTLD returns whether the given string is a TLD (Top Level Domain),
// extended with a few other "TLDs": .bit, .onion
func IsExtendedTLD(s string) bool {
	s = strings.ToUpper(s)
	_, found := ExtendedTLDs[s]
	return found
}

// IsTLD returns whether the given string is a TLD (according to ICANN, or
// in the set of ExtendedTLDs listed in this package.
func IsTLD(s string) bool {
	return IsICANNTLD(s) || IsExtendedTLD(s)
}

// IsDomain returns whether given string is a domain.
// It first checks the TLD, and then uses a regular expression.
func IsDomain(s string) bool {
	fmt.Println("checking suffix")
	if strings.HasSuffix(s, ".") {
		s = s[:len(s)-1]
	}
	fmt.Println("splitting string")
	split := strings.Split(s, ".")
	tld := split[len(split)-1]
	fmt.Println("checking if tld")
	fmt.Println("tld to check", tld)
	if !IsTLD(tld) {
		return false
	}
	fmt.Println("forcing lowercase")
	s = strings.ToLower(s)
	s = strings.Trim(s, "_dnslink.")
	return domainRegexp.MatchString(s)
}
