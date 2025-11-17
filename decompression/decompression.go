package decompression

import (
	"github.com/ipfs/go-cid"
	"github.com/r4f4ss/ztu/bitpackage"
	"github.com/r4f4ss/ztu/fileio"
	"github.com/r4f4ss/ztu/ipfs"
	"github.com/r4f4ss/ztu/params"
)

func Decompress(config *params.Config) error {
	cp, err := fileio.NewCompressedFileFromFile(config.Input)
	if err != nil {
		return err
	}

	cid, err := cid.Parse(cp.Dictionary)
	if err != nil {
		return err
	}
	dictionary, err := ipfs.GetDictionaryByCID(cid)
	if err != nil {
		return err
	}
	pack := bitpackage.NewPack(len(dictionary.Segments), cp.Data)

	outputData := make([]byte, 0, len(cp.Data))
	for cod, b := pack.UnpackingNext(); cod != nil; cod, b = pack.UnpackingNext() {
		if *cod == len(dictionary.Segments) {
			outputData = append(outputData, *b)
		} else if *cod >= 0 {
			outputData = append(outputData, dictionary.Segments[*cod]...)
		}
	}

	err = fileio.WriteDataToFile(config.Output, outputData)
	if err != nil {
		return err
	}

	return nil
}
