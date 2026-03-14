#!/bin/bash

API_BASE="http://localhost:3001/api/v1"
API_KEY="RAHASIA"

echo "Seeding Tables..."
for i in {1..5}
do
  curl -s -X POST "$API_BASE/tables" \
    -H "X-Api-Key: $API_KEY" \
    -H "Content-Type: application/json" \
    -d "{\"number\": $i}" > /dev/null
done

echo "Seeding Menu Categories..."
CAT_COFFEE_ID=$(curl -s -X POST "$API_BASE/menu-categories" \
  -H "X-Api-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name": "Coffee"}' | grep -o '"id":[0-9]*' | cut -d: -f2)

CAT_TEA_ID=$(curl -s -X POST "$API_BASE/menu-categories" \
  -H "X-Api-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name": "Tea"}' | grep -o '"id":[0-9]*' | cut -d: -f2)

CAT_PASTRY_ID=$(curl -s -X POST "$API_BASE/menu-categories" \
  -H "X-Api-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name": "Pastry"}' | grep -o '"id":[0-9]*' | cut -d: -f2)

echo "Seeding Menu Items for Coffee..."
curl -s -X POST "$API_BASE/menu-items" \
  -H "X-Api-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": $CAT_COFFEE_ID, \"name\": \"Espresso\", \"price\": 3.50, \"description\": \"Strong and pure energy.\", \"is_available\": true, \"image_url\": \"https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?q=80&w=600&auto=format&fit=crop\"}" > /dev/null

curl -s -X POST "$API_BASE/menu-items" \
  -H "X-Api-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": $CAT_COFFEE_ID, \"name\": \"Latte\", \"price\": 4.50, \"description\": \"Smooth and creamy classic.\", \"is_available\": true, \"image_url\": \"https://images.unsplash.com/photo-1541167760496-1628856ab772?auto=format&fit=crop&w=600&q=80\"}" > /dev/null

echo "Seeding Menu Items for Tea..."
curl -s -X POST "$API_BASE/menu-items" \
  -H "X-Api-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": $CAT_TEA_ID, \"name\": \"Matcha Green Tea\", \"price\": 5.00, \"description\": \"Earthy and energetic matcha.\", \"is_available\": true, \"image_url\": \"https://images.unsplash.com/photo-1515823064-d6e0c04616a7?auto=format&fit=crop&w=600&q=80\"}" > /dev/null

echo "Seeding Menu Items for Pastry..."
curl -s -X POST "$API_BASE/menu-items" \
  -H "X-Api-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": $CAT_PASTRY_ID, \"name\": \"Croissant\", \"price\": 3.00, \"description\": \"Flaky and buttery indulgence.\", \"is_available\": true, \"image_url\": \"https://images.unsplash.com/photo-1555507036-ab1f4038808a?q=80&w=600&auto=format&fit=crop\"}" > /dev/null

echo "Seeding Complete!"
