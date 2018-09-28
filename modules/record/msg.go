package record

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// name to idetify transaction types
const MsgType = "record"

// MsgRecord
type MsgRecord struct {
	txHash       string
	fileHash     string
	ownerAddress sdk.AccAddress
}

func NewMsgRecord(tx string, file string, owner sdk.AccAddress) MsgRecord {
	return MsgRecord{
		txHash:       tx,
		fileHash:     file,
		ownerAddress: owner,
	}
}

func (msg MsgRecord) Type() string { return MsgType }

func (msg MsgRecord) ValidateBasic() sdk.Error {
	// TO DO
	return nil
}

func (msg MsgRecord) String() string {
	return fmt.Sprintf("MsgRecord{%v, %v}", msg.txHash, msg.fileHash)
}

func (msg MsgRecord) Get(key interface{}) (value interface{}) {
	return nil
}

func (msg MsgRecord) GetSignBytes() []byte {
	b, err := msgCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ownerAddress}
}
