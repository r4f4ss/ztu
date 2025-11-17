package fileio

import (
	"container/list"
	"os"
)

type FileNode struct {
	Data     *byte
	Position uint
}

func GetListFromFile(input string) (*list.List, error) {
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
		node := &FileNode{
			Data:     &b,
			Position: pos,
		}
		fileList.PushBack(node)
		pos++
	}
	return fileList, nil
}
