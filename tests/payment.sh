#!/bin/bash

SERVER_URL="http://localhost:8080"

echo "=== 1. TEST POST /payments/charge (Successful Simulation) ==="
curl -X POST "$SERVER_URL/payments/charge" \
     -H "Content-Type: application/json" \
     -d '{
       "order_id": "ee76d33a-127b-4022-b586-3531fb5b565a",
       "customer_id": "usr_99217a8c",
       "amount": 250.50,
       "currency": "usd",
       "token": "tok_visa"
     }'
echo -e "\n\n"

echo "=== 2. TEST POST /payments/refund ==="
curl -X POST "$SERVER_URL/payments/refund" \
     -H "Content-Type: application/json" \
     -d '{
       "order_id": "ee76d33a-127b-4022-b586-3531fb5b565a"
     }'
echo -e "\n\n"

echo "=== 3. TEST POST /payments/webhook (Simulated Stripe Charge Webhook Payload) ==="
curl -X POST "$SERVER_URL/payments/webhook" \
     -H "Content-Type: application/json" \
     -H "Stripe-Signature: t=123456,v1=mock_signature" \
     -d '{
       "id": "evt_test_webhook_123",
       "object": "event",
       "type": "charge.succeeded",
       "data": {
         "object": {
           "id": "ch_mock_charge_id_777",
           "amount": 25050,
           "currency": "usd",
           "metadata": {
             "payment_id": "ee76d33a-127b-4022-b586-3531fb5b565a"
           }
         }
       }
     }'
echo -e "\n"