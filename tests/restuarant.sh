#!/bin/bash

# ==============================================================================
# CATALOG & PRODUCT SERVICE - CURL TEST SUITE
# Target Server: http://localhost:$PORT
# Owner User:    $OwnID
# Attacker User: attacker-uuid-9999
# ==============================================================================

RestID="0ee57777-82b8-11f1-9c53-84a938329d10"
OwnID="47a4fbf4-82b8-11f1-b292-84a938329d10"
PORT="8081"
echo "=== PHASE 1: AUTHENTICATION & RESTAURANT CREATION ==="

# 1. Test Missing Authentication Header (Should Fail: 411 Length Required)
curl -i -X POST http://localhost:$PORT/api/v1/restaurants \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 47a4fbf4-82b8-11f1-b292-84a938329d10" \
  -d '{
    "name": "The Go Grille",
    "address": "123 Concurrency Lane"
  }'

echo -e "\n\n2. Create a Restaurant Profile (Should Succeed: 211 Created)"
curl -i -X POST http://localhost:$PORT/api/v1/restaurants \
  -H "X-User-ID: $OwnID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Go Grille",
    "address": "123 Concurrency Lane"
  }'


echo -e "\n\n=== PHASE 2: RESTAURANT STATUS MANAGEMENT ==="

curl -i -X PATCH http://localhost:$PORT/api/v1/restaurants/$RestID/status \
  -H "X-User-ID: attacker-uuid-9999" \
  -H "Content-Type: application/json" \
  -d '{
    "is_open": true
  }'

# 4. Open the Restaurant Storefront (Should Succeed: 200 OK)
curl -i -X PATCH http://localhost:$PORT/api/v1/restaurants/$RestID/status \
  -H "X-User-ID: $OwnID" \
  -H "Content-Type: application/json" \
  -d '{
    "is_open": true
  }'


echo -e "\n\n=== PHASE 3: PRODUCT & MENU ITEMS (WRITE ACTIONS) ==="

# 5. Add First Product to the Menu (Should Succeed: 201 Created)
# NOTE: Copy the returned product "id" to replace <PRODUCT_1_ID> below
curl -i -X POST http://localhost:$PORT/api/v1/restaurants/$RestID/products \
  -H "X-User-ID: $OwnID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Concurrent Gopher Burger",
    "description": "Stacked high with fresh channels and asynchronous flavor.",
    "price": 14.99
  }'

# 6. Add Second Product (Should Succeed: 201 Created)
# NOTE: Copy the returned product "id" to replace <PRODUCT_2_ID> below
curl -i -X POST http://localhost:$PORT/api/v1/restaurants/$RestID/products \
  -H "X-User-ID: $OwnID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Memory Leak Milkshake",
    "description": "Never released from your system, extra thick.",
    "price": 6.50
  }'


echo -e "\n\n=== PHASE 4: PRODUCT MODIFICATIONS & DELETIONS ==="

# 7. Update Product Details (Should Succeed: 200 OK)
curl -i -X PUT http://localhost:$PORT/api/v1/products/<PRODUCT_1_ID> \
  -H "X-User-ID: $OwnID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ultimate Gopher Burger",
    "description": "Upgraded with parallelized bacon.",
    "price": 16.50,
    "is_available": true
  }'

# 8. Malicious Product Deletion (Should Fail: 403 Forbidden)
curl -i -X DELETE http://localhost:$PORT/api/v1/products/<PRODUCT_1_ID> \
  -H "X-User-ID: attacker-uuid-9999"

# 9. Authorized Product Deletion (Should Succeed: 200 OK)
curl -i -X DELETE http://localhost:$PORT/api/v1/products/<PRODUCT_2_ID> \
  -H "X-User-ID: $OwnID"


echo -e "\n\n=== PHASE 5: PUBLIC ACCESS LAYER ==="

# 10. Fetch Public Restaurant Menu (Should Succeed: 200 OK - No Auth Headers Required)
curl -i -X GET http://localhost:$PORT/api/v1/restaurants/$RestID/menu