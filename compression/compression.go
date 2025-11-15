package compression

import (
	"container/list"
	"fmt"
	"os"

	"github.com/r4f4ss/ztu/bitpackage"
	"github.com/r4f4ss/ztu/ipfs"
	"github.com/r4f4ss/ztu/params"
)

type fileNode struct {
	data     *byte
	position uint
}

func Compress(config *params.Config) error {

	if !config.IsCompression {
		return fmt.Errorf("not to compress")
	}

	dictionary, err := ipfs.GetDictionaryByCID(config.DictCid)
	if err != nil {
		return err
	}
	pack := bitpackage.NewPack(len(dictionary.Segments), nil)

	fileList, err := getListFromFile(config.Input)
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
				fileFull[e.Value.(*fileNode).position] = i
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
		if e.Value.(*fileNode).data == nil {
			pack.Packing(&fileFull[i], nil)
		} else {
			pack.Packing(nil, e.Value.(*fileNode).data)
		}
		i++
	}

	return nil
}

func setNElementsNill(e *list.Element, n int) {
	for i := 0; i < n; i++ {
		if e == nil {
			break
		}
		e.Value.(*fileNode).data = nil
		e = e.Next()
	}
}

func getListFromFile(input string) (*list.List, error) {
	originalFile, err := os.OpenFile(input, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer originalFile.Close()

	stat, err := originalFile.Stat()
	if err != nil {
		return nil, err
	}

	fileBytes := make([]byte, stat.Size())
	_, err = originalFile.Read(fileBytes)
	if err != nil {
		return nil, err
	}

	fileList := list.New()
	var pos uint = 0
	for _, b := range fileBytes {
		node := &fileNode{
			data:     &b,
			position: pos,
		}
		fileList.PushBack(node)
		pos++
	}
	return fileList, nil
}

func isSameSegment(segment []byte, e *list.Element) bool {
	for i := 0; i < len(segment); i++ {
		if e == nil || e.Value.(*fileNode).data == nil {
			return false
		}
		if segment[i] != *e.Value.(*fileNode).data {
			return false
		}
		e = e.Next()
	}

	return true
}
