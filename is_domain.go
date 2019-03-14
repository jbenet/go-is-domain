package isdomain

import (
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
	if strings.HasSuffix(s, ".") {
		s = s[:len(s)-1]
	}
	split := strings.Split(s, ".")
	tld := split[len(split)-1]
	if !IsTLD(tld) {
		return false
	}
	s = strings.ToLower(s)
	// allow matching of dnslink domains
	// the following only removes `_dnslink.``
	// if it is the leading value.
	// it will remove from _dnslink.example.com
	// but will not remove from example._dnslink.com
	if strings.Contains(s, "_dnslink.") {
		s = strings.TrimPrefix(s, "_dnslink.")
	}
	return domainRegexp.MatchString(s)
}
