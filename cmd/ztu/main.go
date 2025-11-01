package main

import (
	"context"
	"fmt"
	"os"

	"github.com/r4f4ss/ztu/params"
	"github.com/urfave/cli/v3"
)

const descritptionStr = `Examples of use:

- to compress a file 
$ztu -o fileCompressed -i dictionaryCID -c file

- to decompress a file 
$ztu -o fileDecompressed -d file

ztu stands for Zeta Tucanae, which is a solar-type star in the constellation Tucana.`

func main() {
	app := &cli.Command{
		Name:        "ztu",
		Usage:       "An implementation of the ZIPFS compression/decompression specification",
		ArgsUsage:   "file",
		Description: descritptionStr,
		Flags:       params.ZtuFlags,
		Action:      ztu,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err != nil {
			fmt.Println("Failed to write error to stderr", "err", err)
		}
		os.Exit(1)
	}
}

func ztu(_ context.Context, c *cli.Command) error {
	config, err := params.GetConfig(c)
	if err != nil {
		return err
	}

	if config.IsCompression {
		// compression.Compress(config)
		return fmt.Errorf("compression not implemented (wip)")
	} else {
		return fmt.Errorf("decompression not implemented")
	}

	return nil
}
