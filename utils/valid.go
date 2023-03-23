package utils

const (
	payment_initiator = "payment_initiator"
	payment_approver  = "payment_approver"
	admin             = "admin"
)

func IsSupportedRole(role string) bool {
	switch role {
	case payment_initiator, payment_approver, admin:
		return true
	default:
		return false
	}
}
