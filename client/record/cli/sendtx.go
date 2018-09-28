package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/irisnet/irishub/client/context"
	"github.com/irisnet/irishub/client/tendermint/tx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const (
	FlagTxHash     = "tx-hash"
	FlagFileName   = "file-name"
	FlagTargetPath = "target-path"
)

func QueryTxHash(cdc *wire.Codec, cliCtx context.CLIContext, hashHexStr string, trustNode bool) ([]byte, error) {

	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	res, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return nil, err
	}

	if !cliCtx.TrustNode {
		err := ValidateTxResult(cliCtx, res)
		if err != nil {
			return nil, err
		}
	}

	info, err := formatTxResult(cdc, res)
	if err != nil {
		return nil, err
	}

	return cdc.MarshalJSONIndent(info, "", "  ")
}

// ValidateTxResult performs transaction verification
func ValidateTxResult(cliCtx context.CLIContext, res *ctypes.ResultTx) error {
	check, err := cliCtx.Certify(res.Height)
	if err != nil {
		return err
	}

	err = res.Proof.Validate(check.Header.DataHash)
	if err != nil {
		return err
	}
	return nil
}

func GetCmdSubmit(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "submit specified file",
		RunE: func(cmd *cobra.Command, args []string) error {

			// _ = viper.GetString(FlagTargetPath)
			// fileNameStr := viper.GetString(FlagFileName)

			// // build and sign the transaction, then broadcast to Tendermint
			// clictx := context.NewCLIContext().WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

			// // get the from/to address
			// from, err := clictx.GetFromAddress()
			// if err != nil {
			// 	return err
			// }

			// msg := record.NewMsgRecord("this should be ipfs hash", fileNameStr, from)

			// err = msg.ValidateBasic()
			// if err != nil {
			// 	return err
			// }

			// err = clictx.EnsureSignBuildBroadcast(clictx.FromAddressName, []sdk.Msg{msg}, cdc)
			// if err != nil {
			// 	return err
			// }
			return nil
		},
	}

	cmd.Flags().String(FlagFileName, "", "")
	cmd.Flags().String(FlagTargetPath, "", "tx hash")

	return cmd
}

func queryTx(cdc *wire.Codec, cliCtx context.CLIContext, hashHexStr string) ([]byte, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	res, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return nil, err
	}

	if !cliCtx.TrustNode {
		err := tx.ValidateTxResult(cliCtx, res)
		if err != nil {
			return nil, err
		}
	}

	info, err := formatTxResult(cdc, res)
	if err != nil {
		return nil, err
	}

	return cdc.MarshalJSONIndent(info, "", "  ")
}

func GetCmdQureyHash(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query --tx-hash [hash]",
		Short: "query metadata tx hash",
		RunE: func(cmd *cobra.Command, args []string) error {
			// find the key to look up the account
			hashHexStr := viper.GetString(FlagTxHash)

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			output, err := queryTx(cdc, cliCtx, hashHexStr)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil

		},
	}

	cmd.Flags().String(FlagTxHash, "", "tx hash")

	return cmd
}

func GetCmdDownload(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download --tx-hash [hash]",
		Short: "download specified file with tx hash",
		RunE: func(cmd *cobra.Command, args []string) error {
			// find the key to look up the account
			// hashHexStr := viper.GetString(client.FlagTxHash)
			// trustNode := viper.GetBool(client.FlagTrustNode)

			// output, err := QueryTxHash(cdc, context.NewCoreContextFromViper(), hashHexStr, trustNode)
			// if err != nil {
			// 	return err
			// }
			// fmt.Println(string(output))

			return nil
		},
	}

	cmd.Flags().String(FlagTxHash, "", "tx hash")
	cmd.Flags().String(FlagFileName, "", "")
	cmd.Flags().String(FlagTargetPath, "", "tx hash")

	return cmd
}
