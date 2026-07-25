#!/usr/bin/env bash

# System verification testing automation configuration variables
BASE_URL="http://localhost:8080/api/v1"
PRODUCT_UUID="d18d85f8-80f2-4bc9-9fb5-cb52a65492d5"
ORDER_UUID="a3c9b19e-4ff6-42d4-8d4e-b5f7e02d8471"
CUSTOMER_UUID="7e6e5d4c-3b2a-1a0f-9e8d-7c6b5a4b3c2d"

echo "=== 1. TESTING INPUT BINDING VALIDATION (Should Fail due to out-of-bounds metrics) ==="
curl -i -X POST "$BASE_URL/feedbacks" \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": "'"$ORDER_UUID"'",
    "customer_id": "'"$CUSTOMER_UUID"'",
    "product_id": "'"$PRODUCT_UUID"'",
    "product_rating": 99, 
    "product_review": "Absolute garbage validation configuration boundary leak verification test.",
    "courier_rating": 5,
    "courier_review": "Driver was fine."
  }'

echo -e "\n\n=== 2. TESTING HAPPY PATH CREATION (Submission #1) ==="
curl -i -X POST "$BASE_URL/feedbacks" \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": "'"$ORDER_UUID"'",
    "customer_id": "'"$CUSTOMER_UUID"'",
    "product_id": "'"$PRODUCT_UUID"'",
    "product_rating": 5,
    "product_review": "Incredible durability! Best purchase this calendar cycle.",
    "courier_rating": 4,
    "courier_review": "Arrived securely packaged, just slightly off schedule track window."
  }'

echo -e "\n\n=== 3. TESTING HAPPY PATH CREATION WITH ISSUE REPORT (Submission #2) ==="
curl -i -X POST "$BASE_URL/feedbacks" \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": "c8aef103-2411-4b13-bb11-825526e03102",
    "customer_id": "'"$CUSTOMER_UUID"'",
    "product_id": "'"$PRODUCT_UUID"'",
    "product_rating": 2,
    "product_review": "Functional errors observed after deployment runtime uses.",
    "courier_rating": 5,
    "courier_review": "Courier was brilliant.",
    "reported_issue": "Missing connection elements packaging alignment parts."
  }'

echo -e "\n\n=== 4. AWAITING ASYNCHRONOUS BACKGROUND METRICS COMPUTATION (Pausing 2s) ==="
sleep 2

echo -e "\n=== 5. RETRIEVING INTERACTIVE UPSERTED PROJECTION METRICS ==="
curl -i -X GET "$BASE_URL/analytics/$PRODUCT_UUID"