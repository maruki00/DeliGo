#!/bin/bash

# Configuration Variables
BASE_URL="http://localhost:8080/orders"
# We will use this variable to capture the generated UUID from the creation step
ORDER_ID=""

echo "================================================================"
echo "STARTING ORDER SERVICE INTEGRATION TESTS"
echo "================================================================"
echo ""

# 1. CREATE A NEW ORDER
# This will calculate the total amount (2 * 1200.00 + 1 * 150.00 = 2550.00),
# save to PostgreSQL, and fire the 'order.created' event to RabbitMQ.
echo "--> 1. Creating a new order..."
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "cust-uuid-1111-2222",
    "restaurant_id": "rest-uuid-3333-4444",
    "delivery_address": "123 Microservice Lane, Go City",
    "items": [
      {
        "product_id": "prod-macbook",
        "product_name": "MacBook Pro M3",
        "price": 1200.00,
        "quantity": 2
      },
      {
        "product_id": "prod-mouse",
        "product_name": "Magic Mouse",
        "price": 150.00,
        "quantity": 1
      }
    ]
  }')

echo "Response: $CREATE_RESPONSE"
echo ""

# Extract the order ID dynamically if running as a script (requires 'jq')
if command -v jq &> /dev/null; then
    ORDER_ID=$(echo "$CREATE_RESPONSE" | jq -r '.data.id')
    echo "Extracted Order ID: $ORDER_ID"
else
    # Fallback placeholder reminder if running manually
    ORDER_ID="PASTE_YOUR_GENERATED_UUID_HERE"
    echo "[Note] 'jq' not found. If running manually, replace '$ORDER_ID' in subsequent commands with the ID returned above."
fi
echo "---"

# 2. FETCH ORDER DETAILS
# Verify that the order exists and has a status of 'pending_payment'.
echo "--> 2. Fetching order details..."
curl -s -X GET "$BASE_URL/$ORDER_ID"
echo -e "\n---"

# 3. CONFIRM ORDER (PAYMENT RECEIVED)
# Transitions status from 'pending_payment' to 'paid' and fires 'order.paid'.
echo "--> 3. Confirming order payment..."
curl -s -X POST "$BASE_URL/$ORDER_ID/confirm" \
  -H "Content-Type: application/json"
echo -e "\n---"

# 4. RESTAURANT ACCEPTS ORDER
# Transitions status from 'paid' to 'preparing' and fires 'order.status_changed'.
echo "--> 4. Restaurant accepting order..."
curl -s -X POST "$BASE_URL/$ORDER_ID/accept" \
  -H "Content-Type: application/json" \
  -d '{
    "actor": "restaurant"
  }'
echo -e "\n---"

# 5. COURIER ACCEPTS ORDER (READY FOR PICKUP)
# Transitions status from 'preparing' to 'ready_for_pickup'.
echo "--> 5. Courier accepting order (Marking ready for pickup)..."
curl -s -X POST "$BASE_URL/$ORDER_ID/accept" \
  -H "Content-Type: application/json" \
  -d '{
    "actor": "courier"
  }'
echo -e "\n---"

# 6. COURIER PICKS UP ORDER
# Transitions status from 'ready_for_pickup' to 'picked_up'.
echo "--> 6. Courier picking up order..."
curl -s -X POST "$BASE_URL/$ORDER_ID/accept" \
  -H "Content-Type: application/json" \
  -d '{
    "actor": "courier"
  }'
echo -e "\n---"

# 7. TRY TO CANCEL THE ORDER (EXPECT FAILURE)
# Business rules state we cannot cancel an order once it has been picked up.
echo "--> 7. Testing business validation (Attempting to cancel a picked-up order)..."
curl -s -X POST "$BASE_URL/$ORDER_ID/cancel" \
  -H "Content-Type: application/json"
echo -e "\n---"

# 8. COMPLETE DELIVERY
# Transitions status from 'picked_up' to 'delivered' and fires 'order.delivered'.
echo "--> 8. Completing delivery..."
curl -s -X POST "$BASE_URL/$ORDER_ID/complete" \
  -H "Content-Type: application/json"
echo -e "\n---"

# 9. FINAL LIFECYCLE CHECK
# Confirm the ultimate delivery status of the resource.
echo "--> 9. Fetching final order state..."
curl -s -X GET "$BASE_URL/$ORDER_ID"
echo -e "\n\n================================================================"
echo "TEST SUITE COMPLETE"
echo "================================================================"