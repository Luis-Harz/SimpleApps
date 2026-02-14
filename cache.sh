#!/bin/bash

# Cloudflare Daten
ZONE_ID="15a3f624233b6d8edb131836882b5c98"
EMAIL="harztronicsgmbh@gmail.com"
API_KEY="fdea265f1b48dac5962b244f916f3236d6283"

# Cache komplett leeren
response=$(curl -s -X POST "https://api.cloudflare.com/client/v4/zones/$ZONE_ID/purge_cache" \
     -H "X-Auth-Email: $EMAIL" \
     -H "X-Auth-Key: $API_KEY" \
     -H "Content-Type: application/json" \
     --data '{"purge_everything":true}')

# Prüfen, ob es geklappt hat
success=$(echo $response | jq '.success')

if [ "$success" == 'true' ]; then
    echo "✅ Cache erfolgreich geleert!"
else
    echo "❌ Fehler beim Leeren des Caches!"
    echo "Antwort von Cloudflare: $response"
fi
