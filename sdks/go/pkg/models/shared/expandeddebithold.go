// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"math/big"
)

type ExpandedDebitHold struct {
	Description string   `json:"description"`
	Destination *Subject `json:"destination,omitempty"`
	// The unique ID of the hold.
	ID string `json:"id"`
	// Metadata associated with the hold.
	Metadata map[string]string `json:"metadata"`
	// Original amount on hold
	OriginalAmount *big.Int `json:"originalAmount"`
	// Remaining amount on hold
	Remaining *big.Int `json:"remaining"`
	// The ID of the wallet the hold is associated with.
	WalletID string `json:"walletID"`
}
