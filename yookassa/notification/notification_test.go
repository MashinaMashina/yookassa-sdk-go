package yoonotify_test

import (
	"testing"

	"github.com/MashinaMashina/yookassa-sdk-go/yookassa/notification"
	yoopayment "github.com/MashinaMashina/yookassa-sdk-go/yookassa/payment"
	"github.com/stretchr/testify/require"
)

const PaymentBody = `{
  "type": "notification",
  "event": "payment.waiting_for_capture",
  "object": {
    "id": "22d6d597-000f-5000-9000-145f6df21d6f",
    "status": "waiting_for_capture",
    "paid": true,
    "amount": {
      "value": "2.00",
      "currency": "RUB"
    },
    "authorization_details": {
      "rrn": "10000000000",
      "auth_code": "000000",
      "three_d_secure": {
        "applied": true
      }
    },
    "created_at": "2018-07-10T14:27:54.691Z",
    "description": "Заказ №72",
    "expires_at": "2018-07-17T14:28:32.484Z",
    "metadata": {},
    "payment_method": {
      "type": "bank_card",
      "id": "22d6d597-000f-5000-9000-145f6df21d6f",
      "saved": false,
      "card": {
        "first6": "555555",
        "last4": "4444",
        "expiry_month": "07",
        "expiry_year": "2021",
        "card_type": "MasterCard",
      "issuer_country": "RU",
      "issuer_name": "Sberbank"
      },
      "title": "Bank card *4444"
    },
    "refundable": false,
    "test": false
  }
}`

func TestParseNotificationBody(t *testing.T) {
	notify, err := yoonotify.ParseNotificationBody([]byte(PaymentBody))
	require.NoError(t, err)
	require.True(t, notify.IsPayment())
	require.False(t, notify.IsPayout())
	require.False(t, notify.IsRefund())

	payment, err := notify.GetPayment()
	require.NoError(t, err)

	require.Equal(t, yoopayment.WaitingForCapture, payment.Status)
}
