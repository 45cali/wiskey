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

	fqdnOps, _ := parseFqdnFlag(hfqdn)
	for _, host := range hosts {
		//fmt.Println(h)
		h, _ := parseHostName(host)

		isValid := evaluate(h, fqdnOps)
		//fmt.Println(h)
		//fmt.Println(fqdnOps)
		//fmt.Println(isValid)

		if isValid {
			fhosts = append(fhosts, host)
		}

	}

	return
}

// parseParams parses the search params by the comma delimiter and then the equal delimiter
// returns a
func parseFqdnFlag(s string) (map[string]string, error) {
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

// ParseFlags this is a comment
func ParseSearchFlag(s string) map[string]interface{} {

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

func evaluate(h, f map[string]string) (b bool) {
	//fmt.Println(h)
	//	fmt.Println(f)

	for k, v := range f {
		//fmt.Printf("%s == %s %s \n",h[k],v,h[k] == v)
		if h[k] == v {
			b = true
		} else {
			b = false
			break
		}
	}
	return
}

func ParseFieldsFlag(s string) (f []string, b bool) {
	f = strings.Split(s, ",")
	if len(f) > 0 {
		b = true
	}
	return
}
