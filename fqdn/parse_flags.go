package fqdn

import (
	//"fmt"
	//	"os"
	"strings"
)

// ParseFlags this is a comment
func ParseFlags(s string) map[string]interface{} {

	m := make(map[string]interface{})
	commaSplit := strings.Split(s, ",")
	for _, c := range commaSplit {
		//fmt.Println(c)

		equalSplit := strings.Split(c, "=")
		if len(equalSplit) == 2 {
			m[equalSplit[0]] = equalSplit[1]
		}
	}

	return m
}
