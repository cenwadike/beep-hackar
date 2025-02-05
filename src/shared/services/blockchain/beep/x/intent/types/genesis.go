package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:      PortID,
		IntentsList: []Intents{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in intents
	intentsIdMap := make(map[uint64]bool)
	intentsCount := gs.GetIntentsCount()
	for _, elem := range gs.IntentsList {
		if _, ok := intentsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for intents")
		}
		if elem.Id >= intentsCount {
			return fmt.Errorf("intents id should be lower or equal than the last id")
		}
		intentsIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
