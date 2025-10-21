package ipfs

import (
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
)

const stDictionaryStr = `
		type Dictionary struct {
			Name                 String
			Description optional String
			Segments             [Bytes]
		}
	`

type Dictionary struct {
	Name        string
	Description *string
	Segments    [][]byte
}

func generateSchemaTypeDictionary() (schema.Type, error) {
	ts, err := ipld.LoadSchemaBytes([]byte(stDictionaryStr))
	if err != nil {
		return nil, err
	}
	schemaType := ts.TypeByName("Dictionary")

	return schemaType, nil
}
