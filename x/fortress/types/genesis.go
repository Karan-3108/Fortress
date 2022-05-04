package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FortressList: []Fortress{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in fortress
	fortressIdMap := make(map[uint64]bool)
	fortressCount := gs.GetFortressCount()
	for _, elem := range gs.FortressList {
		if _, ok := fortressIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for fortress")
		}
		if elem.Id >= fortressCount {
			return fmt.Errorf("fortress id should be lower or equal than the last id")
		}
		fortressIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
