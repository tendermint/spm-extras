package wasmcmd

import (
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmcli "github.com/CosmWasm/wasmd/x/wasm/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GenesisWasmMsgCmd returns add-wasm-genesis-message cobra Command.
func GenesisWasmMsgCmd(defaultNodeHome string) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "add-wasm-genesis-message",
		Short:                      "Wasm genesis subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	genesisIO := wasmcli.NewDefaultGenesisIO()

	txCmd.AddCommand(
		wasmcli.GenesisStoreCodeCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisInstantiateContractCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisExecuteContractCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisListContractsCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisListCodesCmd(defaultNodeHome, genesisIO),
	)

	return txCmd
}

// GetEnabledProposals parses the ProposalsEnabled / EnableSpecificProposals values to
// produce a list of enabled proposals to pass into wasmd app.
func GetEnabledProposals(proposalsEnabled, enableSpecificProposals string) []wasm.ProposalType {
	if enableSpecificProposals == "" {
		if proposalsEnabled == "true" {
			return wasm.EnableAllProposals
		}
		return wasm.DisableAllProposals
	}
	chunks := strings.Split(enableSpecificProposals, ",")
	proposals, err := wasm.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}
	return proposals
}

// AddModuleInitFlags adds module init flags to the start command.
func AddModuleInitFlags(startCmd *cobra.Command) {
	wasm.AddModuleInitFlags(startCmd)
}
