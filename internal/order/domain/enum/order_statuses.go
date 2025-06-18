package enum

type OrderStatus int

const (
	OrderStatusUnknown OrderStatus = iota
	OrderStatusCreated
	OrderStatusPaid
	OrderStatusPending
	OrderStatusProcessing
	OrderStatusShipped
	OrderStatusDelivered
	OrderStatusCancelled
	OrderStatusRefunded
)

func (_this OrderStatus) String() string {
	switch _this {
	case OrderStatusUnknown:
		return "Unknown"
	case OrderStatusCreated:
		return "Created"
	case OrderStatusPaid:
		return "Paid"
	case OrderStatusPending:
		return "Pending"
	case OrderStatusProcessing:
		return "Processing"
	case OrderStatusShipped:
		return "Shipped"
	case OrderStatusDelivered:
		return "Delivered"
	case OrderStatusCancelled:
		return "Cancelled"
	case OrderStatusRefunded:
		return "Refunded"
	}
	return "Unknown"
}
