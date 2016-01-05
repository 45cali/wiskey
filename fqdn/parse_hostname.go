package fqdn

import (
	"fmt"
	"os"
	"strings"
)

func Filter(hostnames []string, fqdn string) []string {
	pass

}

func parseParams(s string) (m map[string]string, err error) {
	commaSplit := strings.Split(s, ",")
	for _, c := range commaSplit {
		equalSplit := strings.Split(c, "=")
		if len(commaSplit) != 2 {
			err = fmt.Errorf("could not split %s with a = delimiter", equalSplit)
		}
		m[equalSplit[0]] = equalSplit[1]
	}
	return
}
