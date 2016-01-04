package fqdn

import (
	//"fmt"
	"strings"
)

// ParseFlags this is a comment
func ParseFlags(s string) map[string]string {

	m := make(map[string]string)
	commaSplit := strings.Split(s, ",")
	for _, c := range commaSplit {
		//fmt.Println(c)

		equalSplit := strings.Split(c, "=")

		m[equalSplit[0]] = equalSplit[1]

	}

	return m
}
