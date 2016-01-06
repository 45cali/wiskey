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
	search := flag.String("search", "", "search stuff")
	atype := flag.String("type", "", "type of asset")
	count := flag.Int64("count", 100, "default results returned")
	hfqdn := flag.String("fqdn", "", "filter results by class, product, cluster, business unit or domain")

	flag.Parse()

	if len(*atype) == 0 {
		fmt.Println("-type cannot be blank")
		os.Exit(0)
	}

	fmt.Printf("searching for %s's that have fields %s and returning the first %d and applying the following filters %s\n", *atype, *search, *count, *hfqdn)

	c, _ := vindalu.NewClient(wiskeyConf.Server())

	searchAssets(*atype, *search, *hfqdn, *count, c)

}

func searchAssets(atype, search, hfqdn string, count int64, c *vindalu.Client) {

	qb := fqdn.ParseFlags(search)

	items, err := c.List(atype, nil, qb)

	if err != nil {
		fmt.Println("err")
		os.Exit(1)
	}
	fmt.Println("total results: ", len(items))
	hosts := []string{}
	for _, item := range items {
		hosts = append(hosts, item.Id)

	}

	filtered := fqdn.Filter(hosts, hfqdn)
	fmt.Println("total not filtered: ", len(filtered))
	for _, f := range filtered {
		fmt.Println(f)
	}
}
