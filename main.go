package main

import (
	"flag"
	"fmt"
	wiskeyConf "github.com/45cali/wiskey/config"
	"github.com/45cali/wiskey/fqdn"
	"github.com/45cali/wiskey/help"
	"github.com/vindalu/go-vindalu-client"
	"os"
)

func main() {
	help.Help()
	//fmt.Println("main args")
	args := os.Args

	//fmt.Println(len(args))
	//os.Exit(0)
	//fmt.Println("main flags")

	search := flag.String("search", "", "search stuff")
	atype := flag.String("type", "", "type of asset")
	count := flag.Int64("count", 100, "default results returned")
	hfqdn := flag.String("fqdn", "", "filter results by class, product, cluster, business unit or domain")
	//list := flag.String("list", "", "type of asset")
	//	fmt.Println("main parse flags")

	flag.Parse()

	//if len(*atype) == 0 {
	//	fmt.Println("-type cannot be blank")
	//	os.Exit(0)
	//}

	//fmt.Printf("searching for %s's that have fields %s and returning the first %d and applying the following filters %s\n", *atype, *search, *count, *hfqdn)
	//fmt.Println("main vindalu client")

	c, _ := vindalu.NewClient(wiskeyConf.Server())
	//fmt.Println("main switch case")

	switch {
	case len(args) == 3 && len(*atype) > 0:
		//	fmt.Println("switch only -type")

		searchAssets(*atype, *search, *hfqdn, *count, c)
	case len(*hfqdn) > 0 && len(*atype) > 0:
		searchAssets(*atype, *search, *hfqdn, *count, c)

	case len(*search) > 0 && len(*atype) > 0:
		//	fmt.Println("switch -search and -type")

		searchAssets(*atype, *search, *hfqdn, *count, c)
	default:
		//	fmt.Println("switch default help")

		fmt.Print(help.Usage)
	}

}

func searchAssets(atype, search, hfqdn string, count int64, c *vindalu.Client) {

	qb := fqdn.ParseFlags(search)
	fmt.Println(qb)
	items, err := c.List(atype, nil, qb)

	if err != nil {
		fmt.Println("wiskey was unable to connect to server")
		os.Exit(0)
	}
	fmt.Println("total results: ", len(items))
	hosts := []string{}
	for _, item := range items {
		hosts = append(hosts, item.Id)

	}

	filtered := fqdn.Filter(hosts, hfqdn)
	fmt.Println("total not filtered: ", len(filtered))
	//	for _, f := range filtered {
	//		fmt.Println(f)
	//	}

}
