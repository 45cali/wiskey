package main

import (
	"flag"
	"fmt"
	"github.com/45cali/wiskey/fqdn"
	//	"github.com/codegangsta/cli"
	//"github.com/45cali/wiskey/help"
	"github.com/vindalu/go-vindalu-client"
	"os"
)

func main() {

	search := flag.String("search", "", "search stuff")
	atype := flag.String("type", "", "type of asset")
	count := flag.Int64("count", 100, "default results returned")
	fqdn := flag.String("fqdn", "", "filter results by class, product, cluster, business unit or domain")

	flag.Parse()

	if len(*atype) == 0 {
		fmt.Println("-type cannot be blank")
		os.Exit(0)
	}

	fmt.Printf("searching for %s's that have fields %s and returning the first %d and applying the following filters %s\n", *atype, *search, *count, *fqdn)

	c, _ := vindalu.NewClient("http://vindalu.cloudsys.tmcs/")

	searchAssets(*atype, *search, *count, c)

}

func searchAssets(atype, search string, count int64, c *vindalu.Client) {
	//c, _ := vindalu.NewClient("http://vindalu.cloudsys.tmcs/")

	q := fqdn.ParseFlagsOld(search)
	//fmt.Println(q)
	//items, err := c.List(atype, nil, q)

	if err != nil {
		fmt.Println("err")
		os.Exit(1)
	}
	fmt.Println("total results: ", len(items))
}
