package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	epochtypes "github.com/osmosis-labs/osmosis/x/epochs/types"
)

// Parameter store keys
var (
	KeyRefreshEpochIdentifier = []byte("RefreshEpochIdentifier")
	KeyMinimumRiskFactor      = []byte("MinimumRiskFactor")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(refreshEpochIdentifier string, minimumRiskFactor sdk.Dec) Params {
	return Params{
		RefreshEpochIdentifier: refreshEpochIdentifier,
		MinimumRiskFactor:      minimumRiskFactor,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		RefreshEpochIdentifier: "day",
		MinimumRiskFactor:      sdk.NewDecWithPrec(5, 2), // 5%
	}
}

// validate params
func (p Params) Validate() error {
	if err := epochtypes.ValidateEpochIdentifierInterface(p.RefreshEpochIdentifier); err != nil {
		return err
	}
	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRefreshEpochIdentifier, &p.RefreshEpochIdentifier, epochtypes.ValidateEpochIdentifierInterface),
		paramtypes.NewParamSetPair(KeyMinimumRiskFactor, &p.MinimumRiskFactor, ValidateMinimumRiskFactor),
	}
}

func ValidateMinimumRiskFactor(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) || v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("minimum risk factor should be between 0 - 100: %s", v.String())
	}

	return nil
}
