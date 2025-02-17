// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"math/big"
	"time"
)

type AccountBalance struct {
	AccountID     string    `json:"accountId"`
	Asset         string    `json:"asset"`
	Balance       *big.Int  `json:"balance"`
	CreatedAt     time.Time `json:"createdAt"`
	Currency      string    `json:"currency"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}
