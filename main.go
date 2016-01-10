package main

import (
	"flag"
	"fmt"
	wiskeyConf "github.com/45cali/wiskey/config"
	"github.com/45cali/wiskey/fqdn"
	"github.com/45cali/wiskey/help"
	"github.com/vindalu/go-vindalu-client"
	"os"
	"strings"
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
	list := flag.String("list", "", "type of asset")
	asset := flag.String("asset", "", "expects asset id")
	fields := flag.String("fields", "", "retrieve only specified fields of an asset")
	aVersion := flag.Int64("asset-version", 0, "version of specified asset")
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

	case len(args) == 3 && *list == "types":
		//fmt.Println("case list types")
		listTypes(c)

	case len(args) == 5 && *list == "fields" && len(*atype) > 0:
		//fmt.Println("case list fields type: ", *atype)
		listTypeProperties(*atype, c)
	case len(*asset) > 0 && len(*atype) > 0:
		//fmt.Println("case -type ... -asset id fields  ", *fields, *aVersion)
		getAssetFields(*atype, *asset, *fields, *aVersion, c)
	default:
		//	fmt.Println("switch default help")

		fmt.Print(help.Usage)
	}

}

func searchAssets(atype, search, hfqdn string, count int64, c *vindalu.Client) {
	atypeSplit := strings.Split(atype, ",")

	qb := fqdn.ParseSearchFlag(search)
	//fmt.Println(qb)
	hosts := []string{}

	for _, a := range atypeSplit {

		items, err := c.List(a, nil, qb)

		if err != nil {
			fmt.Println("wiskey was unable to connect to server")
			os.Exit(0)
		}

		//fmt.Println("total results: ", len(items))

		for _, item := range items {
			hosts = append(hosts, item.Id)

		}

	}
	fmt.Println("total assets: ", len(hosts))
	filtered := fqdn.Filter(hosts, hfqdn)
	fmt.Println("total assets not filtered: ", len(filtered))
	//	for _, f := range filtered {
	//		fmt.Println(f)
	//	}

}

func listTypes(c *vindalu.Client) {
	items, err := c.GetTypes()
	if err != nil {
		fmt.Println("wiskey was unable to connect to server")
		os.Exit(0)
	}
	for _, item := range items {
		fmt.Println(item.Name)
	}
}

func listTypeProperties(s string, c *vindalu.Client) {
	items, err := c.ListTypeProperties(s)

	if err != nil {
		fmt.Println("wiskey was unable to connect to server", err)
		os.Exit(0)
	}

	for _, item := range items {
		fmt.Println(item)
	}

}

func getAssetFields(atype, asset, fields string, version int64, c *vindalu.Client) {
	items, err := c.Get(atype, asset, version)
	if err != nil {
		fmt.Println("wiskey was unable to connect to server", err)
		os.Exit(0)
	}
	for k, v := range items.Data {
		fmt.Println(k, ": ", v)
	}
}
