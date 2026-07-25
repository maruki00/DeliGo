#!/usr/bin/env bash
set -euo pipefail

# System parameters configuration points
API_URL="http://localhost:8080/api/v1"

echo "=== 1. Creating Users (via Admin Context Authorization) ==="

# Register standard customer
CUSTOMER_RESP=$(curl -s -X POST "$API_URL/users" \
  -H "Content-Type: application/json" \
  -H "X-User-Role: admin" \
  -d '{
    "email": "customer@delivery.com",
    "password": "securepassword123",
    "phone": "+1555019922",
    "first_name": "John",
    "last_name": "Doe",
    "role": "customer"
  }')
echo "Customer Creation JSON Response Output:"
echo "$CUSTOMER_RESP" | grep -o '"id":"[^"]*' | grep -o '[^"]*$' || true
CUSTOMER_ID=$(echo "$CUSTOMER_RESP" | grep -o '"id":"[^"]*' | grep -o '[^"]*$' || echo "FAILED")

# Register standard courier
COURIER_RESP=$(curl -s -X POST "$API_URL/users" \
  -H "Content-Type: application/json" \
  -H "X-User-Role: admin" \
  -d '{
    "email": "courier@delivery.com",
    "password": "fastdelivery456",
    "phone": "+1555014488",
    "first_name": "Jane",
    "last_name": "Speedy",
    "role": "courier"
  }')
COURIER_ID=$(echo "$COURIER_RESP" | grep -o '"id":"[^"]*' | grep -o '[^"]*$' || echo "FAILED")

echo -e "\n=== 2. Authorized Operation: Get Profile details by Admin Context ==="
curl -i -X GET "$API_URL/users/$CUSTOMER_ID" \
  -H "X-User-Role: admin"

echo -e "\n=== 3. Unauthorized Operation: Customer attempting to ban Courier (Should receive 403 Forbidden) ==="
curl -i -X POST "$API_URL/users/$COURIER_ID/ban" \
  -H "Content-Type: application/json" \
  -H "X-User-Role: customer" \
  -d '{"status": "banned"}'

echo -e "\n=== 4. Authorized Operation: Admin Executing Ban Request against Courier (Should receive 200 OK + Publish RMQ Event) ==="
curl -i -X POST "$API_URL/users/$COURIER_ID/ban" \
  -H "Content-Type: application/json" \
  -H "X-User-Role: admin" \
  -d '{"status": "banned"}'