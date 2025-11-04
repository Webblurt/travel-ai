import { writable, get } from 'svelte/store'
import type { ItineraryResponse } from '@/entities/itinerary'
import { getItinerary, getItemDetails } from '@/entities/itinerary'

export const result = writable<ItineraryResponse | null>(null)
export const loading = writable(false)
export const error = writable<string | null>(null)
export const selectedItem = writable<any | null>(null)

export async function generateItinerary(city: string, days: string, budget: string, currency: string) {
  if (!city.trim()) {
    error.set('Please enter a city.')
    return
  }

  loading.set(true)
  error.set(null)
  result.set(null)

  try {
    const data = await getItinerary(city, days, budget, currency)
    if (!data || !data.city || !Array.isArray(data.days)) {
      throw new Error('Invalid JSON structure')
    }
    result.set(data)
  } catch (err: any) {
    console.error(err)
    if (String(err).includes('invalid JSON')) {
      error.set('Invalid AI response, please try again')
    } else {
      error.set('Error while creating itinerary')
    }
  } finally {
    loading.set(false)
  }
}

export async function selectItem(section: 'morning' | 'afternoon' | 'evening', dayNum: number) {
  const current = get(result)
  if (!current?.id) return

  try {
    const item = await getItemDetails(current.id, dayNum, section)
    selectedItem.set(item)
  } catch (err) {
    console.error('Error fetching details', err)
  }
}

export function closeModal() {
  selectedItem.set(null)
}