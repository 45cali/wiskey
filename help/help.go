package help

import (
	"fmt"
	"os"
)

const Usage = `
wiskey [options] [filters]

options:
	-h, --help, help      Show this message
	-type <asset type>    list results of specified type
	-list types           list avaliable asset types
	-list fields          list searchable fields of specified asset type (must be used
	                      with the -type flag)

filters:
	-search               only returns assets if specified type with specified field values
	-fqdn                 can filter 'server' type by: class, superClass,instance,
	                      superInstance, product, cluster, businessUnit, domain

`

func Help() {
	args := os.Args
	for _, a := range args {
		if a == "help" || a == "-h" || a == "--help" {
			fmt.Println(Usage)
			os.Exit(0)
		}
	}
}
