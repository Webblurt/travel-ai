import type { Writable } from 'svelte/store'

export interface PlaceInfo {
  place: string
  description: string
  activity_type: string
  duration: string
  cost: string
  address?: string
}

export interface DayPlan {
  day_number: number
  morning: PlaceInfo
  afternoon: PlaceInfo
  evening: PlaceInfo
}

export interface ItineraryResponse {
  id: string
  city: string
  itinerary_cost: string
  days: DayPlan[]
}

