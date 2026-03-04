"""
Unlocoder API examples - Python
Get your free API key at: https://rapidapi.com/contactliamnoonan/api/unlocoder
Requires: pip install requests
"""

import requests

API_KEY = "YOUR_RAPIDAPI_KEY"
BASE_URL = "https://unlocoder.p.rapidapi.com"
HEADERS = {
    "Content-Type": "application/json",
    "x-rapidapi-key": API_KEY,
    "x-rapidapi-host": "unlocoder.p.rapidapi.com",
}


def convert_coordinate(input_text: str, precision: int = 6) -> dict:
    """Convert any coordinate format to all other formats."""
    response = requests.post(
        f"{BASE_URL}/api/convert",
        headers=HEADERS,
        json={"input": input_text, "precision": precision},
    )
    response.raise_for_status()
    return response.json()


def lookup_unlocode(code: str, reference_time: str | None = None) -> dict:
    """Get timezone and location data for a UN/LOCODE."""
    params = {}
    if reference_time:
        params["referenceTime"] = reference_time
    response = requests.get(
        f"{BASE_URL}/unlocodes/{code}",
        headers=HEADERS,
        params=params,
    )
    response.raise_for_status()
    return response.json()


def find_nearby(latitude: float, longitude: float) -> list[dict]:
    """Find the nearest UN/LOCODEs to a coordinate pair."""
    response = requests.get(
        f"{BASE_URL}/unlocodes/nearby",
        headers=HEADERS,
        params={"latitude": latitude, "longitude": longitude},
    )
    response.raise_for_status()
    return response.json()


if __name__ == "__main__":
    # 1. Convert decimal degrees
    result = convert_coordinate("51.5074, -0.1278")
    print(f"Detected: {result['detectedFormat']}")
    for fmt, value in result["outputs"].items():
        print(f"  {fmt}: {value}")

    # 2. Convert MGRS with high precision
    result = convert_coordinate("18TWL8563012345", precision=10)
    print(f"\nMGRS -> DD: {result['outputs']['DecimalDegrees']}")

    # 3. Resolve a UN/LOCODE
    result = convert_coordinate("GBLON")
    loc = result.get("location", {})
    print(f"\nGBLON -> {loc.get('name')}, TZ: {loc.get('timezoneId')}")

    # 4. Direct UN/LOCODE lookup
    locode = lookup_unlocode("USNYC")
    tz = locode.get("timezone", {})
    print(f"\nUS NYC: {locode['name']}, UTC{tz.get('utcOffset')}")

    # 5. Lookup with summer reference time (different UTC offset)
    locode = lookup_unlocode("USNYC", reference_time="2025-07-15T12:00:00Z")
    tz = locode.get("timezone", {})
    print(f"US NYC (summer): UTC{tz.get('utcOffset')}")

    # 6. Find nearby UN/LOCODEs
    nearby = find_nearby(40.7128, -74.0060)
    print("\nNearby UN/LOCODEs to New York:")
    for loc in nearby:
        print(f"  {loc['country']}{loc['location']} - {loc['distanceKm']:.1f} km")
