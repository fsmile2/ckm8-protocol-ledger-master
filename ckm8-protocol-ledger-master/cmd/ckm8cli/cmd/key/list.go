package key

import (
	"fmt"

	"github.com/spf13/cobra"
	"https://github.com/fsmile2/ckm8/cmd/ckm8cli/cmd/utils"
	"https://github.com/fsmile2/ckm8/wallet"
	wtypes "https://github.com/fsmile2/ckm8/wallet/types"
)

// listCmd lists all the stored keys
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all keys",
	Long:    `List all keys.`,
	Example: "ckm8cli key list",
	Run: func(cmd *cobra.Command, args []string) {
		cfgPath := cmd.Flag("config").Value.String()
		wallet, err := wallet.OpenWallet(cfgPath, wtypes.WalletTypeSoft, true)
		if err != nil {
			utils.Error("Failed to open wallet: %v\n", err)
		}

		keyAddresses, err := wallet.List()
		if err != nil {
			utils.Error("Failed to list keys: %v\n", err)
		}

		for _, keyAddress := range keyAddresses {
			fmt.Printf("%s\n", keyAddress.Hex())
		}
	},
}
