package params

import (
	"fmt"
	"strings"

	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v3"
)

var ZtuFlags = []cli.Flag{
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
	&cli.StringFlag{
		Name:    "cid",
		Usage:   "CID of dictionary for compression",
		Aliases: []string{"i"},
	},
}

type Config struct {
	Output        string
	Input         string
	IsCompression bool
	DictCid       cid.Cid
}

func GetConfig(c *cli.Command) (*Config, error) {

	compress := c.Bool("compress")
	decompress := c.Bool("decompress")
	output := c.String("output")
	dictCidStr := c.String("cid")
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

	var dictCid cid.Cid
	var err error
	if isCompression && strings.Compare(dictCidStr, "") == 0 {
		return nil, fmt.Errorf("required dictionary cid for compress file")
	} else if isCompression {
		dictCid, err = cid.Decode(dictCidStr)
		if err != nil {
			return nil, fmt.Errorf("invalid CID")
		}
	}

	conf := &Config{
		Output:        output,
		Input:         file,
		IsCompression: isCompression,
		DictCid:       dictCid,
	}
	return conf, nil
}
