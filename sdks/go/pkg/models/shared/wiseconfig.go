// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WiseConfig struct {
	APIKey string `json:"apiKey"`
	// The frequency at which the connector will try to fetch new BalanceTransaction objects from Stripe API.
	//
	PollingPeriod *string `json:"pollingPeriod,omitempty"`
}
