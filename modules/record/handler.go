package record

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Handle all "record" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgRecord:
			return handlerRecord(ctx, msg, keeper)
		default:
			errMsg := "Unrecognized Upgrade Msg type: " + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handlerRecord(ctx sdk.Context, msg sdk.Msg, k Keeper) sdk.Result {

	MsgRecord, ok := msg.(MsgRecord)
	if !ok {
		errMsg := "Unrecognized record msg type"
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	k.Record(ctx, MsgRecord)

	return sdk.Result{
		Code: 0,
		Log:  fmt.Sprintf("Record for %s", MsgRecord.fileHash),
	}
}
