#!/bin/bash

auth_token="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjg5MTk3NzAsInVzZXJuYW1lIjoidGhvbWFzIn0.uAzlvOztv9gKuQ03h1Rcj9ZHCcP296thQymxLRtlb5w"

for i in {1..1000}; do
  customer_id=$(( $RANDOM % 10 + 1 ))
  order_date=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
  total=$(bc -l <<< "scale=3; $RANDOM / 1000.0")
  status=$(shuf -n 1 -e "pending" "shipped" "delivered")
  order_data=$(jq -n --argjson customer_id $customer_id --arg order_date "$order_date" --argjson total $total --arg status "$status" '{customer_id: $customer_id, order_date: $order_date, total: $total, status: $status}')

  curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: $auth_token" \
    -d "$order_data" \
    http://localhost:8000/api/orders
done

# curl -X POST \
#     -H "Content-Type: application/json" \
#     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjg5MTk3NzAsInVzZXJuYW1lIjoidGhvbWFzIn0.uAzlvOztv9gKuQ03h1Rcj9ZHCcP296thQymxLRtlb5w" \
#     -d '{"order_date": "2024-10-12T05:34:02Z", "total": "24.643", "status": "delivered" }' \
#     http://localhost:8000/api/orders