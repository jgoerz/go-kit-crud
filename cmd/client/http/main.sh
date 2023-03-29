#!/usr/bin/env bash

HAS_JQ="$(which jq)"
if [ "${HAS_JQ}x" == "x" ]; then
	echo "
	This script requires the program jq to work properly.
	"
	exit 1
fi   

SERVICE_PORT=${ADDRESSBOOK_HTTP_PORT:-3333}
BASE_PATH="/addressbook/v1"

payload=$(cat <<'EOF'
{
  "tenant_id": 123,
  "first_name": "jane",
  "last_name": "doe",
  "active": true,
  "address": "123 Main Street",
  "some_secret": "secret"
}
EOF
)

curl \
  -XPOST \
  -H "Content-Type: application/json" \
  --data "${payload}" \
  http://localhost:${SERVICE_PORT}${BASE_PATH}/contacts | jq .


echo ""
echo "Intentional bad ID in API call"
echo ""
curl \
  http://localhost:${SERVICE_PORT}${BASE_PATH}/contacts/321 | jq .

echo ""
echo "Reading"
echo ""
curl \
  http://localhost:${SERVICE_PORT}${BASE_PATH}/contacts/1 | jq .


update=$(cat <<'EOF'
{
  "id": 1,
  "tenant_id": 123,
  "first_name": "jane-updated",
  "last_name": "doe-updated",
  "active": false,
  "address": "123 Main Street-updated",
  "some_secret": "secret-updated"
}
EOF
)

echo ""
echo "Updating"
echo ""
curl \
  -XPUT \
  -H "Content-Type: application/json" \
  --data "${update}" \
  http://localhost:${SERVICE_PORT}${BASE_PATH}/contacts/1 | jq .

echo ""
echo "Reading"
echo ""
curl \
  http://localhost:${SERVICE_PORT}${BASE_PATH}/contacts/1 | jq .
