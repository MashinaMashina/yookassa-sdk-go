package notification

import (
	"encoding/json"
	"fmt"

	yoopayment "github.com/MashinaMashina/yookassa-sdk-go/yookassa/payment"
	yoopayout "github.com/MashinaMashina/yookassa-sdk-go/yookassa/payout"
	yoorefund "github.com/MashinaMashina/yookassa-sdk-go/yookassa/refund"
)

var ErrUnknownType = fmt.Errorf("unknown type")
var ErrUnknownEvent = fmt.Errorf("unknown event")
var ErrNotSupportedForObject = fmt.Errorf("not supported for object")

type Notification struct {
	typo      Type
	event     Event
	isPayment bool
	isRefund  bool
	isPayout  bool
	object    any
}

type intermediateNotification struct {
	Typo   Type            `json:"type"`
	Event  Event           `json:"event"`
	Object json.RawMessage `json:"object"`
}

func ParseNotificationBody(body []byte) (*Notification, error) {
	var tmp intermediateNotification
	err := json.Unmarshal(body, &tmp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal into intermediate: %w", err)
	}

	notify := &Notification{
		typo:  tmp.Typo,
		event: tmp.Event,
	}
	if notify.typo != TypeNotification {
		return nil, ErrUnknownType
	}

	switch tmp.Event {
	case EventPaymentWaitingForCapture, EventPaymentSucceeded, EventPaymentCanceled:
		notify.isPayment = true
		notify.object = &yoopayment.Payment{}
	case EventRefundSucceeded:
		notify.isRefund = true
		notify.object = &yoorefund.Refund{}
	case EventPayoutSucceeded, EventPayoutCanceled:
		notify.isPayout = true
		notify.object = &yoopayout.Payout{}
	default:
		return nil, ErrUnknownEvent
	}

	err = json.Unmarshal(tmp.Object, &notify.object)
	if err != nil {
		return nil, fmt.Errorf("unmarshal object: %w", err)
	}

	return notify, nil
}

func (r *Notification) IsPayment() bool {
	return r.isPayment
}

func (r *Notification) GetPayment() (*yoopayment.Payment, error) {
	if !r.isPayment {
		return nil, ErrNotSupportedForObject
	}

	return r.object.(*yoopayment.Payment), nil
}

func (r *Notification) IsRefund() bool {
	return r.isRefund
}

func (r *Notification) GetRefund() (*yoorefund.Refund, error) {
	if !r.isRefund {
		return nil, ErrNotSupportedForObject
	}

	return r.object.(*yoorefund.Refund), nil
}

func (r *Notification) IsPayout() bool {
	return r.isPayout
}

func (r *Notification) GetPayout() (*yoopayout.Payout, error) {
	if !r.isPayout {
		return nil, ErrNotSupportedForObject
	}

	return r.object.(*yoopayout.Payout), nil
}
