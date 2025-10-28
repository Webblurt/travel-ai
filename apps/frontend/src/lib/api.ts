import type { ItineraryResponse } from './types'

const BASE_URL = import.meta.env.VITE_API_URL

export async function getItinerary(
  city: string,
  days = '3',
  budget = 'any',
  currency = '$'
): Promise<ItineraryResponse> {
  const params = new URLSearchParams({
    city,
    days,
    budget,
    currency
  })

  const res = await fetch(`${BASE_URL}/api/v1/itineraries?${params.toString()}`)

  if (!res.ok) {
    throw new Error(`request error: ${res.status} ${res.statusText}`)
  }

  try {
    return await res.json()
  } catch (err) {
    console.error('Error while parsing json:', err)
    throw new Error('Unvalid server response')
  }
}
