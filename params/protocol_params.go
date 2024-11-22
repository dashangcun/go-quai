// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math"
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
)

const (
	GasLimitBoundDivisor    uint64 = 1024    // The bound divisor of the gas limit, used in update calculations.
	PercentGasUsedThreshold uint64 = 90      // Percent Gas used threshold at which the gas limit adjusts
	GenesisGasLimit         uint64 = 5000000 // Gas limit of the Genesis block.

	StateCeil                 uint64 = 20000000 // Maximum the StateCeil may ever be
	StateLimitBoundDivisor    uint64 = 1024     // The bound divisor of the gas limit, used in update calculations.
	PercentStateUsedThreshold uint64 = 90       // Percent Gas used threshold at which the gas limit adjusts

	EtxStateUsed uint64 = 10000 // state used by a simple etx

	MaximumExtraDataSize  uint64 = 32    // Maximum size extra data may be after Genesis.
	CallValueTransferGas  uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	TxGas                 uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxGasContractCreation uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroGas         uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv          uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	LogDataGas            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend           uint64 = 2300  // Free gas given at beginning of call.
	ETXGas                uint64 = 21000 // Per ETX generated by opETX or normal cross-chain transfer.
	//  The etx fractions  should be based on the current  expansion number
	ETXRegionMaxFraction int    = 1   // The maximum fraction of transactions for cross-region ETXs
	ETXPrimeMaxFraction  int    = 1   // The maximum fraction of transactions for cross-prime ETXs
	ETXRLimitMin         int    = 10  // Minimum possible cross-region ETX limit
	ETXPLimitMin         int    = 10  // Minimum possible cross-prime ETX limit
	EtxExpirationAge     uint64 = 100 // Number of blocks an ETX may wait for inclusion at the destination

	Sha3Gas     uint64 = 30 // Once per SHA3 operation.
	Sha3WordGas uint64 = 6  // Once per word of the SHA3 operation's data.

	JumpdestGas   uint64 = 1     // Once per JUMPDEST operation.
	EpochDuration uint64 = 30000 // Duration between proof-of-work epochs.

	CreateDataGas         uint64 = 200   //
	CallCreateDepth       uint64 = 1024  // Maximum depth of call/create stack.
	ExpGas                uint64 = 10    // Once per EXP instruction
	LogGas                uint64 = 375   // Per LOG* operation.
	CopyGas               uint64 = 3     //
	StackLimit            uint64 = 1024  // Maximum size of VM stack allowed.
	TierStepGas           uint64 = 0     // Once per operation, for a selection of them.
	LogTopicGas           uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	CreateGas             uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	Create2Gas            uint64 = 32000 // Once per CREATE2 operation
	SelfdestructRefundGas uint64 = 24000 // Refunded following a selfdestruct operation.
	MemoryGas             uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.

	TxDataNonZeroGas          uint64 = 16   // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.
	TxAccessListAddressGas    uint64 = 2400 // Per address specified in access list
	TxAccessListStorageKeyGas uint64 = 1900 // Per storage key specified in access list

	// These have been changed during the course of the chain
	SloadGas        uint64 = 800  // This is only used in the Qi tx processing
	SelfdestructGas uint64 = 5000 // Cost of SELFDESTRUCT

	// EXP has a dynamic portion depending on the size of the exponent
	ExpByte uint64 = 50 // was raised to 50

	// Extcodecopy has a dynamic AND a static cost. This represents only the
	// static portion of the gas.
	ExtcodeCopyBase uint64 = 700

	// CreateBySelfdestructGas is used when the refunded account is one that does
	// not exist. This logic is similar to call.
	CreateBySelfdestructGas uint64 = 25000

	BaseFeeChangeDenominator = 8          // Bounds the amount the base fee can change between blocks.
	ElasticityMultiplier     = 2          // Bounds the maximum gas limit a block may have.
	InitialBaseFee           = 1 * Wei    // Initial base fee for blocks.
	MaxBaseFee               = 100 * GWei // Maximum base fee for blocks.
	InitialStateLimit        = 5000000    // Initial state fee for blocks.

	MaxCodeSize = 24576 // Maximum bytecode to permit for a contract

	// Precompiled contract gas prices

	EcrecoverGas        uint64 = 3000 // Elliptic curve sender recovery gas price
	Sha256BaseGas       uint64 = 60   // Base price for a SHA256 operation
	Sha256PerWordGas    uint64 = 12   // Per-word price for a SHA256 operation
	Ripemd160BaseGas    uint64 = 600  // Base price for a RIPEMD160 operation
	Ripemd160PerWordGas uint64 = 120  // Per-word price for a RIPEMD160 operation
	IdentityBaseGas     uint64 = 15   // Base price for a data copy operation
	IdentityPerWordGas  uint64 = 3    // Per-work price for a data copy operation

	Bn256AddGas             uint64 = 150   // Gas needed for an elliptic curve addition
	Bn256ScalarMulGas       uint64 = 6000  // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseGas     uint64 = 45000 // Base price for an elliptic curve pairing check
	Bn256PairingPerPointGas uint64 = 34000 // Per-point price for an elliptic curve pairing check

	// The Refund Quotient is the cap on how much of the used gas can be refunded
	RefundQuotient uint64 = 5

	MaxAddressGrindAttempts int = 1000 // Maximum number of attempts to grind an address to a valid one
	MinimumEtxGasDivisor        = 5    // The divisor for the minimum gas for inbound ETXs (Block gas limit / MinimumEtxGasDivisor)
	MaximumEtxGasMultiplier     = 2    // Multiplied with the minimum ETX gas for inbound ETXs (Block gas limit / MinimumEtxGasDivisor) * MaximumEtxGasMultiplier
	MinEtxCount                 = 50   // These counts are used in the case where tx is not eligible to be started
	MaxEtxCount                 = 100

	// Dynamic Expansion parameters

	//  This is the threshold (range 0-100) above which the
	// score will begin the tree expansion decision process. This threshold should be
	// 	chosen high enough to not be easily triggered by minor changes in node
	// 	operating behavior, but not so high that the security efficiency becomes
	// 	unacceptably low.
	TREE_EXPANSION_THRESHOLD uint16 = math.MaxUint16

	// This is the smoothing factor (range 0-10) used by each zone in its low-pass
	// filter to gather a long running average of the zone's security efficiency
	// score. Choosing a larger will make the filter less responsive; the tree
	// expansion algorithm will be less susceptible to short term variations in the
	// efficiency score, but will take longer to decide to trigger an expansion when
	// one becomes necessary.
	TREE_EXPANSION_FILTER_ALPHA uint16 = 9

	//  Once all chains have confirmed above TREE_EXPANSION_THRESHOLD, this is
	//  the number of consecutive prime blocks that must remain above the
	//  threshold to confirm the decision to expand the tree.
	TREE_EXPANSION_TRIGGER_WINDOW uint16 = 144

	// Once the network has confirmed the decision to expand the tree, this is
	// the number of prime blocks to wait until the expansion is activated. This
	// should be chosen to give node operators some time to adjust their
	// infrastructure, if needed, to account for the upcoming network change.
	TREE_EXPANSION_WAIT_COUNT = 1024

	OldConversionLockPeriod       uint64 = 10 // The number of zone blocks that a conversion output is locked for
	NewConversionLockPeriod       uint64 = 7200
	MinQiConversionDenomination          = 10
	ConversionConfirmationContext        = common.PRIME_CTX // A conversion requires a single coincident Dom confirmation
	QiToQuaiConversionGas                = 100000           // The gas used to convert Qi to Quai
	DefaultCoinbaseLockup                = 0                // The default lockup byte for coinbase rewards
)

var (
	MaxGossipsubPacketSize            = 3 << 20
	GasCeil                    uint64 = 30000000
	ColosseumGasCeil           uint64 = 30000000
	GardenGasCeil              uint64 = 30000000
	OrchardGasCeil             uint64 = 30000000
	LighthouseGasCeil          uint64 = 30000000
	LocalGasCeil               uint64 = 30000000
	DurationLimit                     = big.NewInt(5) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	GardenDurationLimit               = big.NewInt(5) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	OrchardDurationLimit              = big.NewInt(5) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	LighthouseDurationLimit           = big.NewInt(5) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	LocalDurationLimit                = big.NewInt(1) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	TimeToStartTx              uint64 = 5 * BlocksPerDay
	BlocksPerDay               uint64 = new(big.Int).Div(big.NewInt(86400), DurationLimit).Uint64() // BlocksPerDay is the number of blocks per day assuming 12 second block time
	DifficultyAdjustmentPeriod        = big.NewInt(720)                                             // This is the number of blocks over which the average has to be taken
	DifficultyAdjustmentFactor int64  = 40                                                          // This is the factor that divides the log of the change in the difficulty
	MinQuaiConversionAmount           = new(big.Int).Mul(big.NewInt(10000000000), big.NewInt(GWei)) // 0.000000001 Quai
	MaxWorkShareCount                 = 16
	OldWorkSharesThresholdDiff        = 3 // Number of bits lower than the target that the default consensus engine uses
	NewWorkSharesThresholdDiff        = 4 // Number of bits lower than the target that the default consensus engine uses
	WorkSharesInclusionDepth          = 3 // Number of blocks upto which the work shares can be referenced and this is protocol enforced
	LockupByteToBlockDepth            = make(map[uint8]uint64)
	LockupByteToRewardsRatio          = make(map[uint8]*big.Int)
	ExchangeRate                      = big.NewInt(86196385918997143) // This is the initial exchange rate in Qi per Quai in Its/Qit // Garden = big.NewInt(166666666666666667)
	// These numbers should be "equivalent" to the initial conversion rate
	QuaiToQiConversionBase          = big.NewInt(10000000) // UNUSED Is the starting "historical conversion" in Qits for 10,000 Quai we need 10,000*1e3
	QiToQuaiConversionBase          = big.NewInt(10000000) // UNUSED Is the starting "historical conversion" in Qits for 10,000 Qi we need 10,000*1e3
	OneOverKqi                      = big.NewInt(30000000) // This is the number of hashes need to get 1 Qit. 3e9 is ~$0.001 // = big.NewInt(500)
	MaxTimeDiffBetweenBlocks int64  = 100                  // Max time difference between the blocks to 100 secs
	OneOverAlpha                    = big.NewInt(200)      // The alpha value for the quai to qi conversion
	ControllerKickInBlock    uint64 = 1000000000

	MinBaseFeeInQits              = big.NewInt(5)
	OneOverBaseFeeControllerAlpha = big.NewInt(100)
	BaseFeeMultiplier             = big.NewInt(50)
)

const (
	GoldenAgeForkNumberV2       = 588000
	GoldenAgeForkGraceNumber    = 100
	GoldenAgeGracePaymentPeriod = 6000
)

func init() {
	LockupByteToBlockDepth[0] = OldConversionLockPeriod
	LockupByteToBlockDepth[1] = 30240  // 1.75 days
	LockupByteToBlockDepth[2] = 60480  // 3.5 days
	LockupByteToBlockDepth[3] = 120960 // 7 days

	LockupByteToRewardsRatio[1] = big.NewInt(24) // additional 16% WPY
	LockupByteToRewardsRatio[2] = big.NewInt(10) // additional 20% WPY
	LockupByteToRewardsRatio[3] = big.NewInt(4)  // additional 25% WPY
}

// This is TimeFactor*TimeFactor*common.NumZonesInRegion*common.NumRegionsInPrime
func PrimeEntropyTarget(expansionNum uint8) *big.Int {
	numRegions, numZones := common.GetHierarchySizeForExpansionNumber(expansionNum)
	regionTimeFactor := int64(max(numRegions, 2))
	zoneTimeFactor := int64(max(numZones, 2))
	timeFactorMultiplied := new(big.Int).Mul(big.NewInt(regionTimeFactor), big.NewInt(zoneTimeFactor))
	return new(big.Int).Mul(timeFactorMultiplied, new(big.Int).SetUint64(numZones*numRegions))
}

// This is TimeFactor*common.NumZonesInRegion
func RegionEntropyTarget(expansionNum uint8) *big.Int {
	_, numZones := common.GetHierarchySizeForExpansionNumber(expansionNum)
	timeFactor := int64(max(numZones, 2))
	return new(big.Int).Mul(big.NewInt(timeFactor), new(big.Int).SetUint64(numZones))
}

func MinGasLimit(number uint64) uint64 {
	if number < GoldenAgeForkNumberV2 {
		return 10000000
	} else {
		return 12000000
	}
}

// Gas calculation functions

func SstoreSetGas(stateSize, contractSize *big.Int) uint64 {
	return CalculateGasWithStateScaling(stateSize, contractSize, 20000) // Once per SSTORE operation from clean zero to non-zero
}

func SstoreSentryGas(stateSize, contractSize *big.Int) uint64 {
	return CalculateGasWithStateScaling(stateSize, contractSize, 2300) // Minimum gas required to be present for an SSTORE call, not consumed
}

func ColdAccountAccessCost(stateSize, contractSize *big.Int) uint64 {
	return CalculateGasWithStateScaling(stateSize, contractSize, 2600) // COLD_ACCOUNT_ACCESS_COST
}

func WarmStorageReadCost(stateSize, contractSize *big.Int) uint64 {
	return CalculateGasWithStateScaling(stateSize, contractSize, 100) // WARM_STORAGE_READ_COST
}

func ColdSloadCost(stateSize, contractSize *big.Int) uint64 {
	return CalculateGasWithStateScaling(stateSize, contractSize, 2100) // COLD_SLOAD_COST
}

func SstoreResetGas(stateSize, contractSize *big.Int) uint64 {
	return CalculateGasWithStateScaling(stateSize, contractSize, 5000) // Once per SSTORE operation from clean non-zero to something else
}

func CallNewAccountGas(stateSize *big.Int) uint64 {
	return CalculateGasWithStateScaling(stateSize, big.NewInt(0), 25000) // Paid for CALL when the destination address didn't exist prior.
}

// SSTORE_CLEARS_SCHEDULE is defined as SSTORE_RESET_GAS + ACCESS_LIST_STORAGE_KEY_COST
// Which becomes: 5000 - 2100 + 1900 = 4800
func SstoreClearsScheduleRefund(stateSize, contractSize *big.Int) uint64 {
	return SstoreResetGas(stateSize, contractSize) - ColdSloadCost(stateSize, contractSize) + TxAccessListStorageKeyGas // Once per SSTORE operation for clearing an originally existing storage slot
}

func CalculateGasWithStateScaling(stateSize, contractSize *big.Int, baseRate uint64) uint64 {
	var scalingFactor *big.Int
	if stateSize.Sign() != 0 {
		scalingFactor = common.LogBig(stateSize)
	}
	if contractSize.Sign() != 0 {
		logContractSize := common.LogBig(contractSize)
		scalingFactor = new(big.Int).Add(scalingFactor, logContractSize)
	}
	// If we can assume that the gas price constants is correct for level 4 trie
	num := new(big.Int).Mul(scalingFactor, big.NewInt(int64(baseRate)))
	den := new(big.Int).Mul(big.NewInt(4), common.Big2e64)
	return new(big.Int).Div(num, den).Uint64()
}

func CalculateCoinbaseValueWithLockup(value *big.Int, lockupByte uint8) *big.Int {
	if lockupByte == 0 {
		return value
	}
	return new(big.Int).Add(value, new(big.Int).Div(value, LockupByteToRewardsRatio[lockupByte]))
}

func CalculateQiGasWithUTXOSetSizeScalingFactor(scalingFactor float64, baseRate uint64) uint64 {
	if scalingFactor < 15 {
		return baseRate
	}
	return uint64(scalingFactor*float64(baseRate)) / 15
}
