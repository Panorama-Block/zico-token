package seqs

import (
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evmMcms "github.com/smartcontractkit/mcms/sdk/evm"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink/deployment/common/changeset/internal/ops"
	"github.com/smartcontractkit/chainlink/deployment/common/view/v1_0"
)

type SeqGrantRolesTimelockDeps struct {
	Chain cldf_evm.Chain
}

type RolesAndAddresses struct {
	Role      common.Hash
	Name      string
	Addresses []common.Address
}

type SeqGrantRolesTimelockInput struct {
	ContractType       cldf.ContractType   `json:"contractType"`
	ChainSelector      uint64              `json:"chainSelector"`
	Timelock           common.Address      `json:"timelock"`
	RolesAndAddresses  []RolesAndAddresses `json:"rolesAndAddresses"`
	IsDeployerKeyAdmin bool                `json:"isDeployerKeyAdmin"`
}

type SeqGrantRolesTimelockOutput struct {
	McmsTxs []mcmsTypes.Transaction `json:"mcmsTxs"`
}

var SeqGrantRolesTimelock = operations.NewSequence(
	"seq-grant-role-with-config",
	semver.MustParse("1.0.0"),
	"Grants appropriate roles to MCMS contracts in the EVM Timelock contract",
	func(b operations.Bundle, deps SeqGrantRolesTimelockDeps, in SeqGrantRolesTimelockInput) (SeqGrantRolesTimelockOutput, error) {
		var (
			mcmsTxs              []mcmsTypes.Transaction
			addressesInInspector []string
			err2                 error
		)

		timelockInspector := evmMcms.NewTimelockInspector(deps.Chain.Client)

		for _, roleAndAddress := range in.RolesAndAddresses {
			switch roleAndAddress.Role {
			case v1_0.PROPOSER_ROLE.ID:
				addressesInInspector, err2 = timelockInspector.GetProposers(b.GetContext(), in.Timelock.Hex())
			case v1_0.CANCELLER_ROLE.ID:
				addressesInInspector, err2 = timelockInspector.GetCancellers(b.GetContext(), in.Timelock.Hex())
			case v1_0.BYPASSER_ROLE.ID:
				addressesInInspector, err2 = timelockInspector.GetBypassers(b.GetContext(), in.Timelock.Hex())
			case v1_0.EXECUTOR_ROLE.ID:
				addressesInInspector, err2 = timelockInspector.GetExecutors(b.GetContext(), in.Timelock.Hex())
			case v1_0.ADMIN_ROLE.ID:
				addressesInInspector = []string{}
			}
			if err2 != nil {
				b.Logger.Errorw("Failed to get addresses from Timelock Inspector",
					"chainSelector", deps.Chain.ChainSelector(),
					"chainName", deps.Chain.Name(),
					"Timelock Address", in.Timelock.Hex(),
					"Role", roleAndAddress.Name,
					"Error", err2,
				)
				return SeqGrantRolesTimelockOutput{}, err2
			}
			for _, addressToGrantRole := range roleAndAddress.Addresses {
				if !slices.Contains(addressesInInspector, addressToGrantRole.Hex()) {
					opReport, err := operations.ExecuteOperation(b, ops.OpEVMGrantRole,
						ops.OpEVMGrantRoleDeps{
							Chain: deps.Chain,
						},
						ops.OpEVMGrantRoleInput{
							TimelockAddress:    in.Timelock,
							Address:            addressToGrantRole,
							IsDeployerKeyAdmin: in.IsDeployerKeyAdmin,
							RoleID:             roleAndAddress.Role,
						},
					)
					if err != nil {
						b.Logger.Errorw("Failed to grant role",
							"chainSelector", deps.Chain.ChainSelector(),
							"chainName", deps.Chain.Name(),
							"Timelock Address", in.Timelock.Hex(),
							"Role Name", roleAndAddress.Name,
							"Address", addressToGrantRole.Hex(),
						)
						return SeqGrantRolesTimelockOutput{}, err
					}

					if !in.IsDeployerKeyAdmin {
						mcmsTxs = append(mcmsTxs, opReport.Output.MCMSTx)
					} else {
						b.Logger.Infow("Role granted",
							"Role Name", roleAndAddress.Name,
							"chainSelector", deps.Chain.ChainSelector(),
							"chainName", deps.Chain.Name(),
							"Timelock Address", in.Timelock.Hex(),
							"Address", addressToGrantRole.Hex(),
						)
					}
				}
			}
		}

		return SeqGrantRolesTimelockOutput{McmsTxs: mcmsTxs}, nil
	},
)
