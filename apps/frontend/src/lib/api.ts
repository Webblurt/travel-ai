export async function generateItinerary(destination: string) {
  const res = await fetch('http://localhost:8080/api/itinerary', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ destination, days: 2 }),
  });
  if (!res.ok) throw new Error('Failed to fetch itinerary');
  return res.json();
}
