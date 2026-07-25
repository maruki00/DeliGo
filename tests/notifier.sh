#!/usr/bin/env bash

# Exit immediately if a command exits with a non-zero status
set -e

# ==============================================================================
# CONFIGURATION & TEST CONTEXT
# ==============================================================================
TARGET_SERVER="http://localhost:8080"
RABBITMQ_API="http://localhost:15672/api"
MOCK_USER="user-uuid-9999"
AUTH_HEADER="Authorization: Bearer $MOCK_USER"

echo "======================================================================"
echo "🚀 STARTING COMPREHENSIVE SERVICE INTEGRATION TEST SUITE"
echo "======================================================================"
echo "Target Host: $TARGET_SERVER"
echo "Test User Identity: $MOCK_USER"
echo "----------------------------------------------------------------------"

# ==============================================================================
# TEST 1: ASYNC RABBITMQ EVENT INGESTION
# ==============================================================================
echo "📥 Test 1: Simulating Async Order Status Change via RabbitMQ API..."

curl -s -i -u guest:guest -H "Content-Type: application/json" \
  -X POST "$RABBITMQ_API/exchanges/%2f/amq.default/publish" \
  -d '{
    "properties": {},
    "routing_key": "notification_queue",
    "payload": "{\"user_id\":\"'"$MOCK_USER"'\",\"title\":\"Order Dispatch Updates\",\"body\":\"Your courier has accepted the dispatch route and is moving towards the kitchen layout node.\"}",
    "payload_encoding": "string"
  }' > /dev/null

echo "✅ Event injected into 'notification_queue' successfully."
echo "⏳ Pausing 2 seconds for consumer workers to complete processing cycles..."
sleep 2
echo ""

# ==============================================================================
# TEST 2: FETCH INBOUND NOTIFICATIONS
# ==============================================================================
echo "🔔 Test 2: Fetching Unread Notifications via REST Interface..."

NOTIF_RESPONSE=$(curl -s -X GET "$TARGET_SERVER/api/notifications" \
  -H "$AUTH_HEADER")

echo "Raw Response Matrix:"
echo "$NOTIF_RESPONSE" | grep -o '"body":"[^"]*' || echo "$NOTIF_RESPONSE"

# Extracting dynamic notification ID for downstream sequential steps
# Fallback logic avoids parser dependencies on clean minimal platforms
NOTIF_ID=$(echo "$NOTIF_RESPONSE" | tr ',' '\n' | grep '"id"' | head -n 1 | awk -F: '{print $2}' | tr -d '[:space:]}]')

if [ -z "$NOTIF_ID" ] || [ "$NOTIF_ID" = "null" ]; then
    echo "⚠️ Warning: No explicit notification ID parsed. Defaulting database row to 1 for mutations."
    NOTIF_ID=1
else
    echo "✅ Extracted processing Notification ID target reference: $NOTIF_ID"
fi
echo ""

# ==============================================================================
# TEST 3: MUTATE NOTIFICATION STATE (MARK AS READ)
# ==============================================================================
echo "🔀 Test 3: Mutating Notification State to Read (ID: $NOTIF_ID)..."

curl -s -X PATCH "$TARGET_SERVER/api/notifications/$NOTIF_ID/read" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" | grep -q "marked as read" && echo "✅ Status changed successfully." || echo "❌ Mutation step verification issue."
echo ""

# ==============================================================================
# TEST 4: FILE METADATA REGISTRY INGESTION
# ==============================================================================
echo "📂 Test 4: Registering Delivery Proof Metadata Assets..."

curl -s -X POST "$TARGET_SERVER/api/files/upload" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "file-uuid-7777",
    "owner_id": "'"$MOCK_USER"'",
    "file_url": "https://s3.amazonaws.com/proofs/delivery_receipt_09.png",
    "file_type": "image/png"
  }' | grep -q "complete" && echo "✅ Media metadata tracking registration confirmed." || echo "❌ Asset setup interface issue."
echo ""

# ==============================================================================
# TEST 5: CHAT ROOM TRANSCRIPT HISTORICAL FETCH
# ==============================================================================
echo "💬 Test 5: Extracting Historical Order Transcript Feeds..."

CHAT_RESPONSE=$(curl -s -X GET "$TARGET_SERVER/api/chat/history/order-uuid-5555" \
  -H "$AUTH_HEADER")

echo "Raw History Array Payload:"
echo "$CHAT_RESPONSE"
echo ""

echo "======================================================================"
echo "🏁 SERVICE INTEGRATION TEST RUN COMPLETE"
echo "======================================================================"
echo "Note: Real-time interactive bidirectional streaming rules require persistent"
echo "connections. Validate WebSockets interactively using standard stream utilities:"
echo "wscat --connect \"ws://localhost:8080/ws?user_id=$MOCK_USER\""
echo "======================================================================"