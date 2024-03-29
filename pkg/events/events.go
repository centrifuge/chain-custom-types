package events

import (
	"github.com/centrifuge/chain-custom-types/pkg/keystore"
	"github.com/centrifuge/chain-custom-types/pkg/loans"
	"github.com/centrifuge/chain-custom-types/pkg/permissions"
	"github.com/centrifuge/chain-custom-types/pkg/pools"
	"github.com/centrifuge/chain-custom-types/pkg/rewards"
)

type Events struct {
	ChainBridge_FungibleTransfer        []EventChainBridgeFungibleTransfer        //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventChainBridgeNonFungibleTransfer     //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventChainBridgeGenericTransfer         //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainBridgeChainWhitelisted        //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventChainBridgeRelayerAdded            //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventChainBridgeRelayerThresholdChanged //nolint:stylecheck,golint

	Claims_Claimed        []EventClaimsClaimed        //nolint:stylecheck,golint
	Claims_RootHashStored []EventClaimsRootHashStored //nolint:stylecheck,golint

	CrowdloanClaim_ClaimPalletInitialized        []EventCrowdloanClaimClaimPalletInitialized    //nolint:stylecheck,golint
	CrowdloanClaim_RewardClaimed                 []EventCrowdloanClaimRewardClaimed             //nolint:stylecheck,golint
	CrowdloanClaim_ClaimLockedAtUpdated          []EventCrowdloanClaimLockedAtUpdated           //nolint:stylecheck,golint
	CrowdloanClaim_ClaimContributionsRootUpdated []EventCrowdloanClaimContributionsRootUpdated  //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanTrieIndexUpdated     []EventCrowdloanClaimCrowdloanTrieIndexUpdated //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanLeaseStartUpdated    []EventCrowdloanClaimLeaseStartUpdated         //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanLeasePeriodUpdated   []EventCrowdloanClaimLeasePeriodUpdated        //nolint:stylecheck,golint

	CrowdloanReward_RewardClaimed            []EventCrowdloanRewardRewardClaimed            //nolint:stylecheck,golint
	CrowdloanReward_RewardPalletInitialized  []EventCrowdloanRewardRewardClaimed            //nolint:stylecheck,golint
	CrowdloanReward_DirectPayoutRatioUpdated []EventCrowdloanRewardDirectPayoutRatioUpdated //nolint:stylecheck,golint
	CrowdloanReward_VestingPeriodUpdated     []EventCrowdloanRewardVestingPeriodUpdated     //nolint:stylecheck,golint
	CrowdloanReward_VestingStartUpdated      []EventCrowdloanRewardVestingStartUpdated      //nolint:stylecheck,golint

	CollatorAllowlist_CollatorAdded   []EventCollatorAllowlistCollatorAdded   //nolint:stylecheck,golint
	CollatorAllowlist_CollatorRemoved []EventCollatorAllowlistCollatorRemoved //nolint:stylecheck,golint

	Fees_FeeChanged    []EventFeesFeeChanged    //nolint:stylecheck,golint
	Fees_FeeToAuthor   []EventFeesFeeToAuthor   //nolint:stylecheck,golint
	Fees_FeeToBurn     []EventFeesFeeToBurn     //nolint:stylecheck,golint
	Fees_FeeToTreasury []EventFeesFeeToTreasury //nolint:stylecheck,golint

	Keystore_KeyAdded   []keystore.EventKeystoreKeyAdded   //nolint:stylecheck,golint
	Keystore_KeyRevoked []keystore.EventKeystoreKeyRevoked //nolint:stylecheck,golint
	Keystore_DepositSet []keystore.EventKeystoreDepositSet //nolint:stylecheck,golint

	LiquidityRewards_NewEpoch []rewards.EventLiquidityRewardsNewEpoch //nolint:stylecheck,golint

	LiquidityRewardsBase_GroupRewarded    []rewards.EventLiquidityRewardsBaseGroupRewarded    //nolint:stylecheck,golint
	LiquidityRewardsBase_StakeDeposited   []rewards.EventLiquidityRewardsBaseStakeDeposited   //nolint:stylecheck,golint
	LiquidityRewardsBase_StakeWithdrawn   []rewards.EventLiquidityRewardsBaseStakeWithdrawn   //nolint:stylecheck,golint
	LiquidityRewardsBase_RewardClaimed    []rewards.EventLiquidityRewardsBaseRewardClaimed    //nolint:stylecheck,golint
	LiquidityRewardsBase_CurrencyAttached []rewards.EventLiquidityRewardsBaseCurrencyAttached //nolint:stylecheck,golint

	Loans_Created                   []loans.EventLoansCreated                   //nolint:stylecheck,golint
	Loans_Borrowed                  []loans.EventLoansBorrowed                  //nolint:stylecheck,golint
	Loans_Repaid                    []loans.EventLoansRepaid                    //nolint:stylecheck,golint
	Loans_WrittenOff                []loans.EventLoansWrittenOff                //nolint:stylecheck,golint
	Loans_Mutated                   []loans.EventLoansMutated                   //nolint:stylecheck,golint
	Loans_Closed                    []loans.EventLoansClosed                    //nolint:stylecheck,golint
	Loans_PortfolioValuationUpdated []loans.EventLoansPortfolioValuationUpdated //nolint:stylecheck,golint
	Loans_WriteOffPolicyUpdated     []loans.EventLoansWriteOffPolicyUpdated     //nolint:stylecheck,golint

	Nfts_DepositAsset []EventNftsDepositAsset //nolint:stylecheck,golint
	Nft_Transferred   []EventNftTransferred   //nolint:stylecheck,golint

	Permissions_Added   []permissions.EventPermissionsAdded   //nolint:stylecheck,golint
	Permissions_Removed []permissions.EventPermissionsRemoved //nolint:stylecheck,golint
	Permissions_Purged  []permissions.EventPermissionsPurged  //nolint:stylecheck,golint

	PoolSystem_Rebalanced        []pools.EventPoolSystemRebalanced        //nolint:stylecheck,golint
	PoolSystem_MaxReserveSet     []pools.EventPoolSystemMaxReserveSet     //nolint:stylecheck,golint
	PoolSystem_EpochClosed       []pools.EventPoolSystemEpochClosed       //nolint:stylecheck,golint
	PoolSystem_SolutionSubmitted []pools.EventPoolSystemSolutionSubmitted //nolint:stylecheck,golint
	PoolSystem_EpochExecuted     []pools.EventPoolSystemEpochExecuted     //nolint:stylecheck,golint
	PoolSystem_Created           []pools.EventPoolSystemCreated           //nolint:stylecheck,golint
	PoolSystem_Updated           []pools.EventPoolSystemUpdated           //nolint:stylecheck,golint

	PoolRegistry_Registered       []pools.EventPoolRegistryRegistered       //nolint:stylecheck,golint
	PoolRegistry_UpdateRegistered []pools.EventPoolRegistryUpdateRegistered //nolint:stylecheck,golint
	PoolRegistry_UpdateExecuted   []pools.EventPoolRegistryUpdateExecuted   //nolint:stylecheck,golint
	PoolRegistry_UpdateStored     []pools.EventPoolRegistryUpdateStored     //nolint:stylecheck,golint
	PoolRegistry_MetadataSet      []pools.EventPoolRegistryMetadataSet      //nolint:stylecheck,golint

	Registry_RegistryCreated []EventRegistryRegistryCreated //nolint:stylecheck,golint
	Registry_Mint            []EventRegistryNftMint         //nolint:stylecheck,golint
}
