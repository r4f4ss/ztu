package ipfs

import (
	"context"
	"time"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/node/bindnode"
)

const timeoutSeconds = 5

func GetDictionaryByCID(cid cid.Cid) (*Dictionary, error) {
	var duration time.Duration = time.Duration(timeoutSeconds * float64(time.Second))
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	api, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}

	cidPath := path.FromCid(cid)
	rc, err := api.Block().Get(ctx, cidPath)
	if err != nil {
		return nil, err
	}

	dicInt := &Dictionary{}
	schemaDictionary, err := generateSchemaTypeDictionary()
	if err != nil {
		return nil, err
	}

	prototype := bindnode.Prototype(dicInt, schemaDictionary)
	nb := prototype.NewBuilder()
	if err := dagcbor.Decode(nb, rc); err != nil {
		return nil, err
	}

	dictionary := bindnode.Unwrap(nb.Build()).(*Dictionary)

	return dictionary, nil
}
