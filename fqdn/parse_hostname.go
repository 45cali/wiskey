package fqdn

import (
	"fmt"
	//	"os"
	"regexp"
	"strings"
)

func Filter(hosts []string, hfqdn string) (fhosts []string) {
	if len(hfqdn) == 0 {
		fhosts = hosts
		return
	}

	fqdnOps, _ := parseFqdn(hfqdn)
	for _, host := range hosts {
		//fmt.Println(h)
		h, _ := parseHostName(host)
		isValid := evaluate(h, fqdnOps)
		if isValid {
			fhosts = append(fhosts, host)
		}

	}

	return
}

// parseParams parses the search params by the comma delimiter and then the equal delimiter
// returns a
func parseFqdn(s string) (map[string]string, error) {
	m := make(map[string]string)
	var err error
	// split strings by ','
	splitByComma := strings.Split(s, ",")

	for _, sc := range splitByComma {

		hasEqualSign := strings.Contains(sc, "=")

		if hasEqualSign {
			splitByEqual := strings.Split(sc, "=")

			if len(splitByEqual) == 2 {

				m[splitByEqual[0]] = splitByEqual[1]
			}
		}
	}

	return m, err
}

func parseHostName(s string) (m map[string]string, err error) {
	m = make(map[string]string)
	host := strings.Split(s, ".")
	switch len(host) {
	case 5:
		//m["class"] = strings.Trim(host[0], "1234567890")
		class := strings.Split(host[0], "-")
		if len(class) == 2 {
			m["superClass"] = strings.Trim(class[0], "1234567890")
			m["class"] = strings.Trim(class[1], "1234567890")

			re := regexp.MustCompile("[0-9]+")
			instance := re.FindAllString(host[0], -1)
			if len(instance) > 0 {
				m["superInstance"] = instance[0]
			}
			if len(instance) == 2 {
				m["instance"] = instance[1]
			}
		} else {
			m["class"] = strings.Trim(class[0], "1234567890")

			re := regexp.MustCompile("[0-9]+")
			instance := re.FindAllString(host[0], -1)
			if len(instance) == 1 {
				m["instance"] = instance[0]
			}
		}

		m["product"] = host[1]
		m["cluster"] = host[2]
		m["businessUnit"] = host[3]
		m["domain"] = host[4]
	case 3:
		m["class"] = host[0]
		m["businessUnit"] = host[1]
		m["domain"] = host[2]
	default:
		err = fmt.Errorf("%s is not a valid fqdn hostname", host)
	}
	return
}

func evaluate(h, f map[string]string) bool {
	return true
}
