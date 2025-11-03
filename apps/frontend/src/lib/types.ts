export interface ItineraryResponse {
  id: string
  city: string
  itinerary_cost: string
  days: DayPlan[]
}

export interface DayPlan {
  day_number: number
  morning: Activity
  afternoon: Activity
  evening: Activity
}

export interface Activity {
  place: string
  description: string
  activity_type: string
  cost: string
  duration: string
  address: string
}
