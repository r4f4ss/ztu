package fileio

import (
	"os"
)

func WriteDataToFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
