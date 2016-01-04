package main

import (
	"fmt"
	"github.com/45cali/wiskey/fqdn"
	"github.com/codegangsta/cli"
	"github.com/vindalu/go-vindalu-client"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "wiskey"
	app.Usage = "cli tool for vindalu"
	app.Version = "0.0.1"
	app.Author = "Nick Colbert"
	app.Commands = []cli.Command{
		{
			Name: "list",
			//ShortName: "",
			Usage: "list all assets",
			Subcommands: []cli.Command{
				{
					Name:   "assets",
					Usage:  "list all assets",
					Action: listAssets,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "field",
							//Value: "class",
							Usage: "flag to pass in field types",
						},

						// FQDN filters
						cli.StringFlag{
							Name: "class, c",
							//Value: "class",
							Usage: "filter results by class",
						},
						cli.StringFlag{
							Name: "instance, i",
							//Value: "inst",
							Usage: "filter results by instance",
						},
						cli.StringFlag{
							Name: "product, p",
							//Value: "prod",
							Usage: "filter results by product",
						},
						cli.StringFlag{
							Name: "cluster, cl",
							//Value: "cluster",
							Usage: "filter results by cluster",
						},
						cli.StringFlag{
							Name: "business-unit, b",
							//Value: "bu",
							Usage: "filter results by instance",
						},
						cli.StringFlag{
							Name: "type, t",
							//Value: "type",
							Usage: "filter by asset type ",
						},
					},
				},
				{
					Name:   "types",
					Usage:  "list types of assets",
					Action: listTypes,
				},
				{
					Name:   "properties",
					Usage:  "list properties of an asset",
					Action: listProperties,
				},
				{
					Name:   "detail",
					Usage:  "list detail about specific asset in json",
					Action: listDetail,
				},
			},
		},
		{
			Name:   "search",
			Usage:  "search assets by passing in a field type and value",
			Action: search,
		},
	}

	app.Run(os.Args)

}

func listAssets(ctx *cli.Context) {
	c, _ := vindalu.NewClient("http://vindalu.cloudsys.tmcs")

	q := fqdn.ParseFlags(ctx.String("field"))

	item, err := c.List("server", q, 500)

	if err != nil {
		fmt.Println("err")
		os.Exit(1)
	}
	fmt.Println("total results: ", len(item))
}

func listTypes(ctx *cli.Context) {
	fmt.Println("this will list available asset types")
}

func listProperties(ctx *cli.Context) {
	fmt.Println("this will list properties of an asset type")
}

func listDetail(ctx *cli.Context) {
	fmt.Println("this will list the details of a specific asset")
}

func search(ctx *cli.Context) {
	fmt.Println("this will search by params")
}
