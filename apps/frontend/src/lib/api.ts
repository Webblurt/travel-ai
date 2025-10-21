export async function generateItinerary() {
  const res = await fetch('http://localhost:8080/api/itinerary', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
  });

  if (!res.ok) {
    throw new Error(`Failed to fetch itinerary: ${res.statusText}`);
  }

  return await res.json();
}
