#!/usr/bin/env bash
# Unlocoder API examples using curl
# Get your free API key at: https://rapidapi.com/contactliamnoonan/api/unlocoder

API_KEY="YOUR_RAPIDAPI_KEY"
HOST="unlocoder.p.rapidapi.com"
BASE="https://$HOST"

# ──────────────────────────────────────────────
# 1. Convert coordinates (Decimal Degrees)
# ──────────────────────────────────────────────
echo "=== Convert DD ==="
curl -s -X POST "$BASE/api/convert" \
  -H "Content-Type: application/json" \
  -H "x-rapidapi-key: $API_KEY" \
  -H "x-rapidapi-host: $HOST" \
  -d '{"input": "51.5074, -0.1278"}' | jq .

# ──────────────────────────────────────────────
# 2. Convert coordinates (DMS format)
# ──────────────────────────────────────────────
echo "=== Convert DMS ==="
curl -s -X POST "$BASE/api/convert" \
  -H "Content-Type: application/json" \
  -H "x-rapidapi-key: $API_KEY" \
  -H "x-rapidapi-host: $HOST" \
  -d '{"input": "40°42'\''46\"N 74°0'\''22\"W"}' | jq .

# ──────────────────────────────────────────────
# 3. Convert with custom precision (10 dp)
# ──────────────────────────────────────────────
echo "=== Convert with precision ==="
curl -s -X POST "$BASE/api/convert" \
  -H "Content-Type: application/json" \
  -H "x-rapidapi-key: $API_KEY" \
  -H "x-rapidapi-host: $HOST" \
  -d '{"input": "51.5074, -0.1278", "precision": 10}' | jq .

# ──────────────────────────────────────────────
# 4. Resolve a UN/LOCODE via convert endpoint
# ──────────────────────────────────────────────
echo "=== Resolve UN/LOCODE ==="
curl -s -X POST "$BASE/api/convert" \
  -H "Content-Type: application/json" \
  -H "x-rapidapi-key: $API_KEY" \
  -H "x-rapidapi-host: $HOST" \
  -d '{"input": "GBLON"}' | jq .

# ──────────────────────────────────────────────
# 5. Lookup UN/LOCODE directly
# ──────────────────────────────────────────────
echo "=== Lookup GBLON ==="
curl -s "$BASE/unlocodes/GBLON" \
  -H "x-rapidapi-key: $API_KEY" \
  -H "x-rapidapi-host: $HOST" | jq .

# ──────────────────────────────────────────────
# 6. Lookup with reference time (historical offset)
# ──────────────────────────────────────────────
echo "=== Lookup with reference time ==="
curl -s "$BASE/unlocodes/USNYC?referenceTime=2025-07-15T12:00:00Z" \
  -H "x-rapidapi-key: $API_KEY" \
  -H "x-rapidapi-host: $HOST" | jq .

# ──────────────────────────────────────────────
# 7. Find nearby UN/LOCODEs
# ──────────────────────────────────────────────
echo "=== Nearby UN/LOCODEs ==="
curl -s "$BASE/unlocodes/nearby?latitude=40.7128&longitude=-74.0060" \
  -H "x-rapidapi-key: $API_KEY" \
  -H "x-rapidapi-host: $HOST" | jq .
