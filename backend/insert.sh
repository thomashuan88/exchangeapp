#!/bin/bash

auth_token=""

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

