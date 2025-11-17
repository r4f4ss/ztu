package fileio

import (
	"os"

	"github.com/fxamacker/cbor/v2"
)

type CompressedFile struct {
	Dictionary string
	Data       []byte
}

func NewCompressedFile(dictionary string, data []byte) *CompressedFile {
	return &CompressedFile{
		Dictionary: dictionary,
		Data:       data,
	}
}

func NewCompressedFileFromFile(filePath string) (*CompressedFile, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := cbor.NewDecoder(file)
	var cp CompressedFile
	err = dec.Decode(&cp)
	if err != nil {
		return nil, err
	}

	return &cp, nil
}

func (cp *CompressedFile) WriteToFile(filePath string) error {
	encodedBytes, err := cbor.Marshal(cp)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, encodedBytes, 0640)
	if err != nil {
		return err
	}

	return nil
}
