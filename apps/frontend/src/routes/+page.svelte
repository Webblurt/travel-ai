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
      error = 'Введите город'
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
      error = 'Error while itinerary creation'
    } finally {
      loading = false
    }
  }
</script>

<main class="p-8 max-w-3xl mx-auto font-sans">
  <h1 class="text-4xl font-bold text-center mb-8">AI Travel Planner</h1>

  <form
    on:submit|preventDefault={handleSubmit}
    class="grid gap-4 bg-neutral-50 p-6 rounded-2xl shadow-lg border border-gray-200"
  >
    <input
      bind:value={city}
      placeholder="City"
      class="p-3 border rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
    />
    <input
      bind:value={days}
      placeholder="Days amount"
      class="p-3 border rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
    />
    <input
      bind:value={budget}
      placeholder="Budget (example 1000)"
      class="p-3 border rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
    />
    <input
      bind:value={currency}
      placeholder="Currency (example $)"
      class="p-3 border rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
    />

    <button
      type="submit"
      class="p-3 rounded-lg bg-blue-600 text-white hover:bg-blue-700 transition-all duration-300 shadow-md disabled:opacity-60"
      disabled={loading}
    >
      {loading ? 'Itinerary creating...' : 'Create Itinerary'}
    </button>
  </form>

  {#if error}
    <p class="text-red-600 mt-4 text-center">{error}</p>
  {/if}

  {#if loading}
    <div class="flex justify-center mt-8 animate-pulse text-gray-500">
      Itinerary uploading...
    </div>
  {/if}

  {#if result}
    <section class="mt-10 space-y-6 animate-fade-in">
      <h2 class="text-2xl font-semibold text-center text-blue-700">
        City itinerary {result.city}
      </h2>
      <p class="text-gray-600 text-center">
        Total cost: <strong>{result.itinerary_cost}</strong>
      </p>

      {#each result.days as day}
        <div
          class="border rounded-xl p-5 bg-white shadow-sm hover:shadow-md transition duration-300"
        >
          <h3 class="text-xl font-semibold mb-3 text-blue-600">
            День {day.day_number}
          </h3>

          {#each ['morning', 'day', 'evening'] as part}
            {@const activity = (day as any)[part]}
            <div class="mb-4 border-l-4 border-blue-400 pl-3">
              <p class="font-medium capitalize text-gray-800">
                {part === 'day' ? 'In the day' : part === 'morning' ? 'In the morning' : 'In the evening'} — {activity?.place}
              </p>
              <p class="text-gray-600 text-sm">{activity?.description}</p>
              <p class="text-sm text-gray-500">
                Type: {activity?.activity_type} • Cost: {activity?.cost} • Time: {activity?.duration}
              </p>
              <p class="text-xs text-gray-400">{activity?.address}</p>
            </div>
          {/each}
        </div>
      {/each}
    </section>
  {/if}
</main>

<style>
  @keyframes fade-in {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
  }
  .animate-fade-in {
    animation: fade-in 0.6s ease-in-out;
  }
</style>
