package yoopayout

import (
	yoocommon "github.com/MashinaMashina/yookassa-sdk-go/yookassa/common"
)

type PayoutReceipt struct {
	ServiceName  string            `json:"service_name,omitempty"`
	NpdReceiptId string            `json:"npd_receipt_id,omitempty"`
	URL          string            `json:"url,omitempty"`
	Amount       *yoocommon.Amount `json:"amount,omitempty"`
}
