package entities

type OrderProductEntity interface {
	GetId() int
	GetOrderId() int
	GetProductId() int
	GetQty() int
}

// id VARCHAR(36) PRIMARY KEY NOT NULL,
//     customer_id VARCHAR(36) NOT NULL,
//     order_status VARCHAR(50) NOT NULL,
//     total_amount DECIMAL(10, 2) NOT NULL,
//     currency VARCHAR(3) NOT NULL,
//     order_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     shipping_address JSONB NULL,
//     billing_address JSONB NULL,
//     payment_id VARCHAR(36) NULL,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
