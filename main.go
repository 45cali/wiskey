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
	args := os.Args

	search := flag.String("search", "", "search stuff")
	atype := flag.String("type", "", "type of asset")
	count := flag.Int64("count", 100, "default results returned")
	hfqdn := flag.String("fqdn", "", "filter results by class, product, cluster, business unit or domain")
	list := flag.String("list", "", "type of asset")
	asset := flag.String("asset", "", "expects asset id")
	fields := flag.String("fields", "", "retrieve only specified fields of an asset")
	aVersion := flag.Int64("asset-version", 0, "version of specified asset default is 0")

	flag.Parse()

	c, _ := vindalu.NewClient(wiskeyConf.Server())

	switch {
	case len(args) == 3 && len(*atype) > 0:
		searchAssets(*atype, *search, *hfqdn, *count, c)

	case len(*hfqdn) > 0 && len(*atype) > 0:
		searchAssets(*atype, *search, *hfqdn, *count, c)

	case len(*search) > 0 && len(*atype) > 0:
		searchAssets(*atype, *search, *hfqdn, *count, c)

	case len(args) == 3 && *list == "types":
		listTypes(c)

	case len(args) == 5 && *list == "fields" && len(*atype) > 0:
		listTypeProperties(*atype, c)

	case len(*asset) > 0 && len(*atype) > 0:
		getAssetFields(*atype, *asset, *fields, *aVersion, c)

	default:
		fmt.Print(help.Usage)
	}
}

func searchAssets(atype, search, hfqdn string, count int64, c *vindalu.Client) {
	atypeSplit := strings.Split(atype, ",")
	qb := fqdn.ParseSearchFlag(search)
	hosts := []string{}

	for _, a := range atypeSplit {

		items, err := c.List(a, nil, qb)
		if err != nil {
			fmt.Println("wiskey was unable to connect to server")
			os.Exit(0)
		}
		for _, item := range items {
			hosts = append(hosts, item.Id)
		}
	}

	filtered := fqdn.Filter(hosts, hfqdn)
	for _, f := range filtered {
		fmt.Println(f)
	}
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

	f, b := fqdn.ParseFieldsFlag(fields)
	if b == true {
		for _, i := range f {
			fmt.Println(i, ": ", items.Data[i])
		}
	} else {

		for k, v := range items.Data {
			fmt.Println(k, ": ", v)
		}
	}
}
