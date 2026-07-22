CREATE TABLE IF NOT EXISTS couriers (
    id VARCHAR(36) PRIMARY KEY,
    vehicle_type VARCHAR(20) NOT NULL CHECK (vehicle_type IN ('bicycle', 'scooter', 'car')),
    is_active BOOLEAN DEFAULT FALSE,
    current_latitude DECIMAL(9,6),
    current_longitude DECIMAL(9,6),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders_couriers (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    courier_id VARCHAR(36) REFERENCES couriers(id),
    status VARCHAR(20) DEFAULT 'searching' CHECK (status IN ('searching', 'accepted', 'rejected', 'at_restaurant', 'picked_up', 'delivered')),
    assigned_at TIMESTAMP NULL,
    picked_up_at TIMESTAMP NULL,
    delivered_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_order ON orders_couriers(order_id);
CREATE INDEX IF NOT EXISTS idx_courier ON orders_couriers(courier_id);

CREATE TABLE IF NOT EXISTS orders_tracking (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_order_tracking ON orders_tracking(order_id);