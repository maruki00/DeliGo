-- Create custom enum types for roles and statuses
CREATE TYPE user_role AS ENUM ('admin', 'customer', 'restaurant_owner', 'courier');
CREATE TYPE user_status AS ENUM ('active', 'suspended', 'banned');

-- Core Users table (Acts as auth credentials + profile source of truth)
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role user_role NOT NULL,
    status user_status DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Casbin Rules table (Managed via GORM adapter, consumed by API Gateway or internal logic)
CREATE TABLE casbin_rule (
    id SERIAL PRIMARY KEY,
    ptype VARCHAR(100) NOT NULL,
    v0 VARCHAR(100), -- Subject (e.g., role or user_id)
    v1 VARCHAR(100), -- Object (e.g., URI path)
    v2 VARCHAR(100), -- Action (e.g., GET, POST)
    v3 VARCHAR(100),
    v4 VARCHAR(100),
    v5 VARCHAR(100)
);

CREATE INDEX idx_casbin_rule ON casbin_rule (ptype, v0, v1, v2);

-- Automated update trigger for handling updated_at mutation
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();