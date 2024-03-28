package notification

type Event string

const (
	EventPaymentWaitingForCapture Event = "payment.waiting_for_capture"
	EventPaymentSucceeded         Event = "payment.succeeded"
	EventPaymentCanceled          Event = "payment.canceled"

	EventRefundSucceeded Event = "refund.succeeded"

	EventPayoutSucceeded Event = "payout.succeeded"
	EventPayoutCanceled  Event = "payout.canceled"

	EventDealClosed Event = "deal.closed"
)
