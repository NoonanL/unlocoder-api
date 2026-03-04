/**
 * Unlocoder API examples - JavaScript (Node.js)
 * Get your free API key at: https://rapidapi.com/contactliamnoonan/api/unlocoder
 * No dependencies required (uses built-in fetch, Node 18+)
 */

const API_KEY = "YOUR_RAPIDAPI_KEY";
const BASE_URL = "https://unlocoder.p.rapidapi.com";
const HEADERS = {
  "Content-Type": "application/json",
  "x-rapidapi-key": API_KEY,
  "x-rapidapi-host": "unlocoder.p.rapidapi.com",
};

async function convertCoordinate(input, precision = 6) {
  const res = await fetch(`${BASE_URL}/api/convert`, {
    method: "POST",
    headers: HEADERS,
    body: JSON.stringify({ input, precision }),
  });
  if (!res.ok) throw new Error(`${res.status}: ${await res.text()}`);
  return res.json();
}

async function lookupUnlocode(code, referenceTime) {
  const params = referenceTime
    ? `?referenceTime=${encodeURIComponent(referenceTime)}`
    : "";
  const res = await fetch(`${BASE_URL}/unlocodes/${code}${params}`, {
    headers: HEADERS,
  });
  if (!res.ok) throw new Error(`${res.status}: ${await res.text()}`);
  return res.json();
}

async function findNearby(latitude, longitude) {
  const res = await fetch(
    `${BASE_URL}/unlocodes/nearby?latitude=${latitude}&longitude=${longitude}`,
    { headers: HEADERS }
  );
  if (!res.ok) throw new Error(`${res.status}: ${await res.text()}`);
  return res.json();
}

async function main() {
  // 1. Convert decimal degrees
  const dd = await convertCoordinate("51.5074, -0.1278");
  console.log(`Detected: ${dd.detectedFormat}`);
  for (const [fmt, value] of Object.entries(dd.outputs)) {
    console.log(`  ${fmt}: ${value}`);
  }

  // 2. Convert DMS
  const dms = await convertCoordinate(`40°42'46"N 74°0'22"W`);
  console.log(`\nDMS -> DD: ${dms.outputs.DecimalDegrees}`);

  // 3. Resolve a UN/LOCODE
  const locode = await convertCoordinate("GBLON");
  const loc = locode.location;
  console.log(`\nGBLON: ${loc?.name}, TZ: ${loc?.timezoneId}`);

  // 4. Direct UN/LOCODE lookup
  const nyc = await lookupUnlocode("USNYC");
  console.log(
    `\nUS NYC: ${nyc.name}, UTC${nyc.timezone?.utcOffset}`
  );

  // 5. Find nearby UN/LOCODEs
  const nearby = await findNearby(40.7128, -74.006);
  console.log("\nNearby UN/LOCODEs to New York:");
  for (const n of nearby) {
    console.log(`  ${n.country}${n.location} - ${n.distanceKm.toFixed(1)} km`);
  }
}

main().catch(console.error);
