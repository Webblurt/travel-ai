import type { ItineraryResponse } from './types'

const BASE_URL = import.meta.env.VITE_API_URL

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
