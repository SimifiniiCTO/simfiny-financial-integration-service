package helper

import (
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// DefaultPockets returns the default pockets for a bank accout.
// These pockets are used to fascilitate the abstraction of sub-bank accounts better enabling proper goal management
// IDEALLY pockets should partition ONLY all BANK ACCOUNTS a given user utilizes.
// TODO: fascilitate this later
func DefaultPockets() []*proto.Pocket {
	return []*proto.Pocket{
		{
			Type: proto.PocketType_POCKET_TYPE_DEBT_REDUCTION,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_EMERGENCY_FUND,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_DISCRETIONARY_SPENDING,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_FUN_MONEY,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_LONG_TERM_SAVINGS,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_SHORT_TERM_SAVINGS,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_INVESTMENT,
		},
	}
}
