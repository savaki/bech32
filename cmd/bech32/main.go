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
	HRP    string
}

func main() {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:    "decode",
			Aliases: []string{"d"},
			Usage:   "decode a bech32 string to a human readable part and hex encoded data",
			Action:  decode,
		},
		{
			Name:    "encode",
			Aliases: []string{"e"},
			Usage:   "encode a human readable part and hex encoded data to bech32",
			Action:  encode,
		},
	}
	app.Usage = "parse bech32 encoded string"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "format",
			Usage:       "output format: text, json",
			Value:       "text",
			Destination: &opts.Format,
		},
	}
	app.Action = infer
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func infer(c *cli.Context) error {
	first := c.Args().First()
	_, _, err := bech32.Decode(first)
	if err != nil {
		return encode(c)
	} else {
		return decode(c)
	}
}

func decode(c *cli.Context) error {
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

func encode(c *cli.Context) error {
	var records []interface{}
	args := c.Args().Slice()
	for i := 0; i < len(args); i += 2 {
		hrp := args[i]
		data, err := hex.DecodeString(args[i+1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "bech32: failed to encode data, %v: %v", data, err)
			os.Exit(1)
		}
		addr, err := bech32.Encode(hrp, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bech32: failed to encode data, %v: %v", data, err)
			os.Exit(1)
		}

		switch opts.Format {
		case "json":
			records = append(records, map[string]interface{}{
				"hrp":  hrp,
				"data": hex.EncodeToString(data),
				"addr": addr,
			})
		default:
			fmt.Printf("%v %v: %v\n", hrp, hex.EncodeToString(data), addr)
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
