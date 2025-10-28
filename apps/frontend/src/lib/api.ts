import type { ItineraryResponse } from './types'

function basicValidate(it: any): it is ItineraryResponse {
  if (!it || typeof it !== 'object') return false
  if (typeof it.city !== 'string') return false
  if (typeof it.itinerary_cost !== 'string') return false
  if (!Array.isArray(it.days)) return false
  for (const d of it.days) {
    if (typeof d.day_number !== 'number' && typeof d.day_number !== 'string') return false
    if (!d.morning || !d.day || !d.evening) return false
  }
  return true
}

export async function getItinerary(
  city: string,
  days?: string,
  budget?: string,
  currency?: string
): Promise<ItineraryResponse> {
  const params = new URLSearchParams()

  params.append('city', city)
  if (days) params.append('days', days)
  if (budget) params.append('budget', budget)
  if (currency) params.append('currency', currency)

  const res = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/itineraries?${params.toString()}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${import.meta.env.VITE_API_TOKEN}`,
    },
  })

  const raw = await res.text()

  if (!res.ok) {
    const err = new Error(`API error: ${res.status} ${res.statusText}`)
    ;(err as any).status = res.status
    ;(err as any).raw = raw
    throw err
  }

  try {
    const parsed = JSON.parse(raw)
    if (!basicValidate(parsed)) {
      const err = new Error('Invalid JSON structure from AI')
      ;(err as any).raw = raw
      throw err
    }
    parsed.days = parsed.days.map((d: any) => ({
      ...d,
      day_number: Number(d.day_number)
    }))
    return parsed as ItineraryResponse
  } catch (e) {
    const err = new Error('Invalid JSON from AI')
    ;(err as any).raw = raw
    throw err
  }
}
