CREATE TABLE feedbacks (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    customer_id VARCHAR(36) NOT NULL,
    
    product_rating INT CHECK (product_rating BETWEEN 1 AND 5),
    product_review TEXT,
    
    courier_rating INT CHECK (courier_rating BETWEEN 1 AND 5),
    courier_review TEXT,

    reported_issue TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);