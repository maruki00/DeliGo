#!/bin/bash
# -----------------------------------------------------------------------------
# Delivery & Dispatch Service - End-to-End API Integration Verification Test
# -----------------------------------------------------------------------------
HOST="http://localhost:8083/api/v1"
ORDER_ID="order-7777-uuid"
COURIER_ID="courier-near-456"

echo "=== 1. Initializing Couriers (Near and Far) ==="
curl -s -X POST "$HOST/couriers" -H "Content-Type: application/json" \
  -d "{\"id\":\"$COURIER_ID\",\"vehicle_type\":\"bicycle\",\"latitude\":40.7128,\"longitude\":-74.0060,\"is_active\":true}"
echo -e "\n---"
curl -s -X POST "$HOST/couriers" -H "Content-Type: application/json" \
  -d "{\"id\":\"courier-far-999\",\"vehicle_type\":\"car\",\"latitude\":40.8500,\"longitude\":-74.2000,\"is_active\":true}"
echo -e "\n\n"

echo "👉 NOTE: At this stage, simulate the RabbitMQ payload by publishing an event"
echo "   to the 'order.confirmed' queue with OrderID: '$ORDER_ID'. Once consumed,"
echo "   the service assigns the closest active courier ($COURIER_ID)."
echo -e "\n"

echo "=== 2. Courier Accepts Order Assignment ==="
curl -s -X POST "$HOST/orders/$ORDER_ID/accept" -H "Content-Type: application/json" \
  -d "{\"courier_id\":\"$COURIER_ID\"}"
echo -e "\n\n"

echo "=== 3. Courier Arrives at Restaurant ==="
curl -s -X POST "$HOST/orders/$ORDER_ID/arrive" -H "Content-Type: application/json" \
  -d "{\"courier_id\":\"$COURIER_ID\"}"
echo -e "\n\n"

echo "=== 4. Courier Picks Up Order (Starts Transit) ==="
curl -s -X POST "$HOST/orders/$ORDER_ID/pickup" -H "Content-Type: application/json" \
  -d "{\"courier_id\":\"$COURIER_ID\"}"
echo -e "\n\n"

echo "=== 5. In-Transit Real-Time GPS Telemetry Ping ==="
curl -s -X POST "$HOST/tracking/ping" -H "Content-Type: application/json" \
  -d "{\"courier_id\":\"$COURIER_ID\",\"order_id\":\"$ORDER_ID\",\"latitude\":40.7135,\"longitude\":-74.0048}"
echo -e "\n\n"

echo "=== 6. Courier Completes Delivery ==="
curl -s -X POST "$HOST/orders/$ORDER_ID/complete" -H "Content-Type: application/json" \
  -d "{\"courier_id\":\"$COURIER_ID\"}"
echo -e "\n\n"
echo "=== Test Suite Complete ==="