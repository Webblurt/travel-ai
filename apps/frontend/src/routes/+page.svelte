<script lang="ts">
  import { getItinerary } from '../lib/api'
  import type { ItineraryResponse } from '../lib/types'
  import DayCard from '../lib/components/DayCard.svelte'

  let city = ''
  let days = ''
  let budget = ''
  let currency = ''
  let loading = false
  let error = ''
  let result: ItineraryResponse | null = null

  async function handleSubmit() {
    if (!city.trim()) {
      error = 'Please enter a city.'
      return
    }

    loading = true
    error = ''
    result = null

    try {
      const data = await getItinerary(city, days, budget, currency)

      if (!data || !data.city || !Array.isArray(data.days)) {
        throw new Error('Invalid JSON structure')
      }

      result = data
    } catch (err: any) {
      console.error('Error:', err)
      if (err.message.includes('Invalid JSON')) {
        error = 'The AI returned an invalid response. Please try again.'
      } else {
        error = 'Something went wrong while generating itinerary. Please try again.'
      }
    } finally {
      loading = false
    }
  }
</script>

<main class="p-6 max-w-3xl mx-auto">
  <h1 class="text-3xl font-bold mb-6 text-center text-blue-700">AI Travel Planner</h1>

  <form on:submit|preventDefault={handleSubmit} class="grid gap-4 bg-gray-50 p-6 rounded-xl shadow-sm">
    <input bind:value={city} placeholder="City" class="p-2 border rounded" />
    <input bind:value={days} placeholder="Number of days" class="p-2 border rounded" />
    <input bind:value={budget} placeholder="Budget (e.g. 1000)" class="p-2 border rounded" />
    <input bind:value={currency} placeholder="Currency (e.g. $)" class="p-2 border rounded" />

    <button
      type="submit"
      class="p-2 rounded bg-blue-600 text-white hover:bg-blue-700 transition disabled:opacity-50"
      disabled={loading}
    >
      {loading ? 'Generating...' : 'Generate itinerary'}
    </button>
  </form>

  {#if error}
    <p class="text-red-600 mt-4 text-center font-medium">{error}</p>
  {/if}

  {#if result}
    <section class="mt-8 space-y-6">
      <h2 class="text-2xl font-semibold text-center">Itinerary for {result.city}</h2>
      <p class="text-gray-700 mb-4 text-center">Estimated total cost: {result.itinerary_cost}</p>

      {#each result.days as day}
        <DayCard {day} />
      {/each}
    </section>
  {/if}
</main>
