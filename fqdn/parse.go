package fqdn

import (
	"fmt"
	"regexp"
	"strings"
)

// Filter takes in a []string and uses hfqdn to filter it and then return a []string
func Filter(hosts []string, hfqdn string) (fhosts []string) {
	if len(hfqdn) == 0 {
		fhosts = hosts
		return
	}
	fqdnOps, _ := parseFqdnFlag(hfqdn)
	for _, host := range hosts {
		h, _ := parseHostName(host)
		isValid := evaluate(h, fqdnOps)
		if isValid {
			fhosts = append(fhosts, host)
		}
	}
	return
}

// parseParams parses the search params by the comma delimiter and then the equal delimiter
// returns a map[string]string and error
func parseFqdnFlag(s string) (map[string]string, error) {
	m := make(map[string]string)
	var err error

	splitByComma := strings.Split(s, ",")

	for _, sc := range splitByComma {

		hasEqualSign := strings.Contains(sc, "=")

		if hasEqualSign {
			splitByEqual := strings.Split(sc, "=")

			if len(splitByEqual) == 2 {

				m[splitByEqual[0]] = splitByEqual[1]
			} else {
				err = fmt.Errorf("%s does not have a length of 2", splitByEqual)
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

// ParseSearchFlag takes in a string and splits it by the comma (,) delimiter
// and then splits each sting by the equal (=) delimiter. ParseSearchFlag then
// maps the equalSplit []string to map[string]interface{}
func ParseSearchFlag(s string) map[string]interface{} {

	m := make(map[string]interface{})
	commaSplit := strings.Split(s, ",")
	for _, c := range commaSplit {

		equalSplit := strings.Split(c, "=")
		if len(equalSplit) == 2 {
			m[equalSplit[0]] = equalSplit[1]
		}
	}
	return m
}

func evaluate(h, f map[string]string) (b bool) {

	for k, v := range f {

		if h[k] == v {
			b = true
		} else {
			b = false
			break
		}
	}
	return
}

// ParseFieldsFlag spilts a string by comma (,) delimiter and returns []string and
// a bool that id true in len([]string) > 0
func ParseFieldsFlag(s string) (f []string, b bool) {
	f = strings.Split(s, ",")
	if len(f) >= 1 {
		b = true
	}
	return
}
