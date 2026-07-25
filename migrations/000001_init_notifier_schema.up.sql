CREATE TABLE notifications (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user ON notifications (user_id);

-- Basic Chat Messages schema for active orders
CREATE TABLE chat_messages (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    sender_id VARCHAR(36) NOT NULL,
    recipient_id VARCHAR(36) NOT NULL,
    message TEXT NOT NULL,
    sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_order_chat ON chat_messages (order_id);

-- File metadata (e.g. proof of delivery images or restaurant banners)
CREATE TABLE files (
    id VARCHAR(36) PRIMARY KEY,
    owner_id VARCHAR(36) NOT NULL,
    file_url VARCHAR(512) NOT NULL,
    file_type VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);