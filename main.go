package main

import (
	"fmt"
	"github.com/codegangsta/cli"
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
			//Usage:     "list all assets",
			Subcommands: []cli.Command{
				{
					Name:   "assets",
					Usage:  "list all assets",
					Action: listAssets,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "class",
							Value: "class",
							Usage: "filter results by class",
						},
						cli.StringFlag{
							Name:  "instance",
							Value: "inst",
							Usage: "filter results by instance",
						},
						cli.StringFlag{
							Name:  "product",
							Value: "prod",
							Usage: "filter results by product",
						},
						cli.StringFlag{
							Name:  "cluster",
							Value: "cluster",
							Usage: "filter results by cluster",
						},
						cli.StringFlag{
							Name:  "business-unit",
							Value: "bu",
							Usage: "filter results by instance",
						},
						cli.StringFlag{
							Name:  "type",
							Value: "type",
							Usage: "filter by asset type",
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
	}

	app.Run(os.Args)

}

func listAssets(ctx *cli.Context) {
	fmt.Println("this will list all assets by asset id")
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
