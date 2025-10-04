package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

var ztuFlags = []cli.Flag{
	&cli.StringFlag{
		Name:      "output",
		Usage:     "Output file name",
		Aliases:   []string{"o"},
		TakesFile: true,
	},
	&cli.BoolFlag{
		Name:        "compress",
		Usage:       "compress file",
		Aliases:     []string{"c"},
		HideDefault: true,
	},
	&cli.BoolFlag{
		Name:        "decompress",
		Usage:       "decompress file",
		Aliases:     []string{"d"},
		HideDefault: true,
	},
}

type config struct {
	output        string
	input         string
	isCompression bool
}

func getConfig(c *cli.Command) (*config, error) {

	compress := c.Bool("compress")
	decompress := c.Bool("decompress")
	output := c.String("output")
	file := c.Args().First()

	if compress && decompress {
		return nil, fmt.Errorf("can not apply both compression and decompression")
	} else if !compress && !decompress {
		return nil, fmt.Errorf("must choose compression or decompression")
	}

	if strings.Compare(file, "") == 0 {
		return nil, fmt.Errorf("required input file")
	}

	isCompression := false
	if compress {
		isCompression = true
	}

	if strings.Compare(output, "") == 0 {
		if isCompression {
			output = file + ".ztu"
		} else {
			return nil, fmt.Errorf("required output file")
		}
	}

	conf := &config{
		output:        output,
		input:         file,
		isCompression: isCompression,
	}
	return conf, nil
}
