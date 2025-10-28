<script lang="ts">
  import { getItinerary } from '../lib/api'
  import type { ItineraryResponse } from '../lib/types'

  let city: string = ''
  let days: string = ''
  let budget: string = ''
  let currency: string = ''
  let loading: boolean = false
  let error: string = ''
  let result: ItineraryResponse | null = null

  async function handleSubmit() {
    if (!city.trim()) {
      error = 'Please, enter the city'
      return
    }

    loading = true
    error = ''
    result = null

    try {
      const data = await getItinerary(city, days, budget, currency)
      result = data
    } catch (err: any) {
      console.error(err)
      error = 'Error while itinerary creating'
    } finally {
      loading = false
    }
  }
</script>

<main class="p-6 max-w-3xl mx-auto">
  <h1 class="text-3xl font-bold mb-6 text-center">AI Travel Planner</h1>

  <form on:submit|preventDefault={handleSubmit} class="grid gap-4 bg-neutral-50 p-6 rounded-xl shadow-sm">
    <input bind:value={city} placeholder="City" class="p-2 border rounded" />
    <input bind:value={days} placeholder="Days count" class="p-2 border rounded" />
    <input bind:value={budget} placeholder="Budget (example 1000)" class="p-2 border rounded" />
    <input bind:value={currency} placeholder="Currency (example $)" class="p-2 border rounded" />

    <button
      type="submit"
      class="p-2 rounded bg-blue-600 text-white hover:bg-blue-700 transition disabled:opacity-50"
      disabled={loading}
    >
      {loading ? 'Creating...' : 'Create itinerary'}
    </button>
  </form>

  {#if error}
    <p class="text-red-600 mt-4">{error}</p>
  {/if}

  {#if result}
    <section class="mt-8 space-y-6">
      <h2 class="text-2xl font-semibold">City itinerary {result.city}</h2>
      <p class="text-gray-700 mb-4">Common Cost: {result.itinerary_cost}</p>

      {#each result.days as day}
        <div class="border rounded-lg p-4 bg-white shadow-sm">
          <h3 class="text-xl font-semibold mb-2">Day {day.day_number}</h3>

          {#each ['morning', 'day', 'evening'] as part}
            {@const activity = (day as any)[part]}
            <div class="mb-3">
              <p class="font-medium capitalize">{part} — {activity?.place}</p>
              <p class="text-gray-600 text-sm">{activity?.description}</p>
              <p class="text-sm text-gray-500">Type: {activity?.activity_type}, Cost: {activity?.cost}, Time: {activity?.duration}</p>
              <p class="text-xs text-gray-400">{activity?.address}</p>
            </div>
          {/each}
        </div>
      {/each}
    </section>
  {/if}
</main>
