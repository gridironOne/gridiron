// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Furychain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package baseapp

import (
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"

	authprecompile "github.com/gridironOne/gridiron/cosmos/precompile/auth"
	bankprecompile "github.com/gridironOne/gridiron/cosmos/precompile/bank"
	distrprecompile "github.com/gridironOne/gridiron/cosmos/precompile/distribution"
	erc20precompile "github.com/gridironOne/gridiron/cosmos/precompile/erc20"
	govprecompile "github.com/gridironOne/gridiron/cosmos/precompile/governance"
	stakingprecompile "github.com/gridironOne/gridiron/cosmos/precompile/staking"
	ethprecompile "github.com/gridironOne/gridiron/eth/core/precompile"
)

// PrecompilesToInject returns a function that provides the initialization of the standard
// set of precompiles.
func PrecompilesToInject(app *GridironBaseApp, customPcs ...ethprecompile.Registrable) func() *ethprecompile.Injector {
	return func() *ethprecompile.Injector {
		// Create the precompile injector with the standard precompiles.
		pcs := ethprecompile.NewPrecompiles([]ethprecompile.Registrable{
			authprecompile.NewPrecompileContract(app.AuthzKeeper, app.AuthzKeeper),
			bankprecompile.NewPrecompileContract(
				bankkeeper.NewMsgServerImpl(app.BankKeeper),
				app.BankKeeper,
			),
			distrprecompile.NewPrecompileContract(
				distrkeeper.NewMsgServerImpl(app.DistrKeeper),
				distrkeeper.NewQuerier(app.DistrKeeper),
			),
			erc20precompile.NewPrecompileContract(
				app.BankKeeper, app.ERC20Keeper,
			),
			govprecompile.NewPrecompileContract(
				govkeeper.NewMsgServerImpl(app.GovKeeper),
				app.GovKeeper,
			),
			stakingprecompile.NewPrecompileContract(app.StakingKeeper),
		}...)

		// Add the custom precompiles to the injector.
		for _, pc := range customPcs {
			pcs.AddPrecompile(pc)
		}
		return pcs
	}
}
