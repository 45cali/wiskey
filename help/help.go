package help

import (
	"fmt"
	"os"
)

// Usage holds the help infomation
const Usage = `
wiskey [options] [filters]

options:
	-h, --help, help                           show this message
	-type <asset type>                         list results of specified type
	-list types                                list avaliable asset types
	-list fields                               list searchable fields of specified
	                                           asset type (must be used with the -type flag)


filters:
	-asset  <asset id>                          return all fields for specified asset id,
	                                            used with -type flag

	-fields <field>,<field>...                  used to filter the results of the -type
	                                            flag when used with the -asset flag

	-search <field>=<value>,<field>=<value>...  only returns assets if specified type
	                                            with specified field values

	-fqdn   class=<value>,instance=<vlaue>...   can filter 'server' type by: class,
	                                            superClass,instance, superInstance, product,
					            cluster, businessUnit, domain


	-asset-version 	<value>                     used with -type and -asset flags to
	                                            specify a specific version of asset data

examples:

	wiskey -h
	   diplay help information

	wiskey -type server
	   return a list of all assets of type server

	wiskey -type server -fqdn class=app,instance=1
	   return a list of all assets of type server and the first instance of the app class

	wiskey -type server -search os=windows
	   return a list of servers running windows

	wiskey -type server -search os=windows -fqdn class=app,instance=1
	   return a list of servers running windows and the first instance of the app class

	wiskey -type server -asset <asset id>
	   return all fields of a specific asset

	wiskey -type server -asset <asset id> -asset-version 0
	   return all fields of a specific asset of specified version

	wiskey -type server -asset <asset id> -fields id,uptime
	   return just the id and uptime of specified asset

	wiskey -list types
	   return a list of available assets

	wiskey -list fields -type server
	   return a list a available fields of asset type server
`

// Help looks for command line args help, -h and --help. If found, the help information is displayed
func Help() {
	args := os.Args
	for _, a := range args {
		if a == "help" || a == "-h" || a == "--help" {
			fmt.Println(Usage)
			os.Exit(0)
		}
	}
}
