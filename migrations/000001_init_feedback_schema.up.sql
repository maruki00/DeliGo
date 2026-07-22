CREATE DATABASE IF NOT EXISTS feedback_db;
USE feedback_db;

CREATE TABLE feedbacks (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    customer_id VARCHAR(36) NOT NULL,
    
    -- Split rating domains
    product_rating INT CHECK (product_rating BETWEEN 1 AND 5),
    product_review TEXT,
    
    courier_rating INT CHECK (courier_rating BETWEEN 1 AND 5),
    courier_review TEXT,
    
    reported_issue TEXT, -- Holds user reports of mistakes/issues
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_order (order_id)
);