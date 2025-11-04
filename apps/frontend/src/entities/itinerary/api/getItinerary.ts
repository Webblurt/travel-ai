import { fetcher } from '@/shared/api/fetcher'
import type { ItineraryResponse } from '../model/types'

export async function getItinerary(city: string, days: string, budget: string, currency: string): Promise<ItineraryResponse> {
  return await fetcher(`${import.meta.env.VITE_API_URL}/itinerary?city=${city}&days=${days}&budget=${budget}&currency=${currency}`, `${import.meta.env.VITE_API_TOKEN}`)
}
