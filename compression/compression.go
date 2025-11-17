package compression

import (
	"container/list"
	"fmt"

	"github.com/r4f4ss/ztu/bitpackage"
	"github.com/r4f4ss/ztu/fileio"
	"github.com/r4f4ss/ztu/ipfs"
	"github.com/r4f4ss/ztu/params"
)

func Compress(config *params.Config) error {

	if !config.IsCompression {
		return fmt.Errorf("not to compress")
	}

	dictionary, err := ipfs.GetDictionaryByCID(config.DictCid)
	if err != nil {
		return err
	}
	pack := bitpackage.NewPack(len(dictionary.Segments), nil)

	fileList, err := fileio.GetListFromFile(config.Input)
	if err != nil {
		return err
	}

	lenFile := fileList.Len()
	fileFull := make([]int, 0, lenFile)
	for range lenFile {
		fileFull = append(fileFull, -1)
	}

	for i := 0; i < len(dictionary.Segments); i++ {
		seg := dictionary.Segments[i]
		for e := fileList.Front(); e != nil; {
			if isSameSegment(seg, e) {
				fileFull[e.Value.(*fileio.FileNode).Position] = i
				setNElementsNill(e, len(seg))
				for range len(seg) {
					e = e.Next()
					if e == nil {
						break
					}
				}
			} else {
				e = e.Next()
				if e == nil {
					break
				}
			}
		}
	}

	i := 0
	for e := fileList.Front(); e != nil; e = e.Next() {
		if e.Value.(*fileio.FileNode).Data == nil {
			pack.Packing(&fileFull[i], nil)
		} else {
			pack.Packing(nil, e.Value.(*fileio.FileNode).Data)
		}
		i++
	}

	fileOut := fileio.NewCompressedFile(config.DictCid.String(), pack.GetData())
	err = fileOut.WriteToFile(config.Output)
	if err != nil {
		return err
	}

	return nil
}

func setNElementsNill(e *list.Element, n int) {
	for i := 0; i < n; i++ {
		if e == nil {
			break
		}
		e.Value.(*fileio.FileNode).Data = nil
		e = e.Next()
	}
}

func isSameSegment(segment []byte, e *list.Element) bool {
	for i := 0; i < len(segment); i++ {
		if e == nil || e.Value.(*fileio.FileNode).Data == nil {
			return false
		}
		if segment[i] != *e.Value.(*fileio.FileNode).Data {
			return false
		}
		e = e.Next()
	}

	return true
}
