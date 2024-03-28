package yoopayout

import (
	"time"

	yoocommon "github.com/MashinaMashina/yookassa-sdk-go/yookassa/common"
)

type PayoutStatus string

const (
	PayoutStatusPending   PayoutStatus = "pending"
	PayoutStatusSucceeded PayoutStatus = "succeeded"
	PayoutStatusCanceled  PayoutStatus = "canceled"
)

type Payout struct {
	Id                  string                     `json:"id,omitempty"`
	Amount              *yoocommon.Amount          `json:"amount,omitempty"`
	Status              PayoutStatus               `json:"status,omitempty"`
	PayoutDestination   *PayoutDestination         `json:"payout_destination,omitempty"`
	Description         string                     `json:"description,omitempty"`
	CreatedAt           time.Time                  `json:"created_at,omitempty"`
	Deal                *PayoutDeal                `json:"deal,omitempty"`
	SelfEmployed        *PayoutSelfEmployed        `json:"self_employed,omitempty"`
	Receipt             *PayoutReceipt             `json:"receipt,omitempty"`
	CancellationDetails *PayoutCancellationDetails `json:"cancellation_details,omitempty"`
	Metadata            interface{}                `json:"metadata,omitempty"`
	Test                bool                       `json:"test,omitempty"`
}
