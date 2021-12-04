package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/savaki/bech32"
	"github.com/urfave/cli/v2"
)

var opts struct {
	Format string
}

func main() {
	app := cli.NewApp()
	app.Usage = "parse bech32 encoded string"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "format",
			Usage:       "output format: text, json",
			Value:       "text",
			Destination: &opts.Format,
		},
	}
	app.Action = action
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var records []interface{}
	for _, addr := range c.Args().Slice() {
		hrp, data, err := bech32.Decode(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bech32: failed to decode addr, %v: %v", addr, err)
			os.Exit(1)
		}

		switch opts.Format {
		case "json":
			records = append(records, map[string]interface{}{
				"addr": addr,
				"hrp":  hrp,
				"data": hex.EncodeToString(data),
			})
		default:
			fmt.Printf("%v: %v %v\n", addr, hrp, hex.EncodeToString(data))
		}
	}

	switch opts.Format {
	case "json":
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(records)
	default:
		return nil
	}
}
