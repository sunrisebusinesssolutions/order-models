package models

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusApproved  OrderStatus = "APPROVED"
	OrderStatusCompleted OrderStatus = "COMPLETED"
	OrderStatusDenied    OrderStatus = "DENIED"
	OrderStatusDeclined  OrderStatus = "DECLINED"
	OrderStatusDeleted   OrderStatus = "DELETED"
)

func OrderStatusFromLabel(label string) OrderStatus {
	switch label {
	case "Approved":
		return OrderStatusApproved
	case "Completed":
		return OrderStatusCompleted
	case "Denied":
		return OrderStatusDenied
	case "Declined":
		return OrderStatusDeclined
	case "Deleted":
		return OrderStatusDeleted
	case "Pending", "":
		fallthrough
	default:
		return OrderStatusPending
	}
}

func (os OrderStatus) String() string {
	switch os {
	case OrderStatusApproved:
		return "Approved"
	case OrderStatusCompleted:
		return "Completed"
	case OrderStatusDenied:
		return "Denied"
	case OrderStatusDeclined:
		return "Declined"
	case OrderStatusDeleted:
		return "Deleted"
	case OrderStatusPending:
		fallthrough
	default:
		return "Pending"
	}
}
