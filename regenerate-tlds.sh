#!/bin/sh

list=tlds-alpha-by-domain.txt

rm -f $list
wget https://data.iana.org/TLD/$list

header=$(head -1 $list)

cat > tlds.go <<EOF
package isdomain

// $header
var TLDs = map[string]bool{
EOF

grep -v '^#' $list | while read tld; do
    echo -e "\t\"$tld\": true,"
done >> tlds.go

cat >> tlds.go <<EOF
}

// ExtendedTLDs is a set of additional "TLDs", allowing decentralized name
// systems, like TOR and Namecoin.
var ExtendedTLDs = map[string]bool{
	"BIT":   true,
	"ONION": true,
}
EOF

gofmt -w tlds.go

