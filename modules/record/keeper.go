package record

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
)

// Record Keeper
type Keeper struct {
	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding.
	cdc *wire.Codec
}

func NewKeeper(cdc *wire.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// Returns the go-wire codec.
func (keeper Keeper) WireCodec() *wire.Codec {
	return keeper.cdc
}

func (keeper Keeper) NewRecord(ctx sdk.Context, addr string, time string, hash string, size string, node string) Record {

	var record Record = Record{
		ownerAddress: addr,
		submitTime:   time,
		dataHash:     hash,
		dataSize:     size,
		pinedNode:    node,
	}
	//keeper.SetRecord(ctx, record)

	return record
}

func (k Keeper) Record(ctx sdk.Context, msg MsgRecord) {
	kvStore := ctx.KVStore(k.storeKey)
	fileHash := msg.fileHash
	msgBytes, err := k.cdc.MarshalBinary(msg)
	if err != nil {
		panic(err)
	}

	kvStore.Set(GetFileHashKey(fileHash), msgBytes)
}
