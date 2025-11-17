package fileio

import (
	"os"
)

func WriteDataToFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
