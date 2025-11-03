<script lang="ts">
  import { getItinerary, getItemDetails } from '../lib/api'
  import DayCard from '../lib/components/DayCard.svelte'
  import ItemModal from '../lib/components/ItemModal.svelte'
  import type { ItineraryResponse } from '../lib/types'

  let city = ''
  let days = ''
  let budget = ''
  let currency = ''
  let loading = false
  let error = ''
  let result: ItineraryResponse | null = null

  let selectedItem: any = null

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
      console.error(err)
      if (String(err).includes('invalid JSON')) {
          error = 'Invalid AI response, please try again'
        } else {
          error = 'Error while creating itinerary'
        }
    } finally {
      loading = false
    }
  }

  async function handleSelect(section: 'morning' | 'afternoon' | 'evening', dayNum: number) {
    if (!result?.id) return
    try {
      const item = await getItemDetails(result.id, dayNum, section)
      selectedItem = item
    } catch (err) {
      console.error('Error fetching details', err)
    }
  }

  function closeModal() {
    selectedItem = null
  }
</script>

<main class="p-6 max-w-3xl mx-auto">
  <h1 class="text-3xl font-bold mb-6 text-center text-blue-700">AI Travel Planner</h1>
  <form on:submit|preventDefault={handleSubmit} class="grid gap-4 bg-gray-50 p-6 rounded-xl shadow-sm">
    <input bind:value={city} placeholder="City" class="p-2 border rounded" />
    <input bind:value={days} placeholder="Number of days" class="p-2 border rounded" />
    <input bind:value={budget} placeholder="Budget (e.g. 1000)" class="p-2 border rounded" />
    <input bind:value={currency} placeholder="Currency (e.g. $)" class="p-2 border rounded" />
    <button type="submit" class="p-2 rounded bg-blue-600 text-white hover:bg-blue-700 transition disabled:opacity-50" disabled={loading}>
      {loading ? 'Generating...' : 'Generate itinerary'}
    </button>
  </form>

  {#if result}
    <div class="mt-6 grid gap-4">
      {#each result.days as day}
        <DayCard {day} onSelect={handleSelect} />
      {/each}
    </div>
  {/if}

  {#if selectedItem}
    <ItemModal item={selectedItem} onClose={closeModal} />
  {/if}

  {#if error}
    <div class="mt-4 bg-red-50 border border-red-300 text-red-700 p-4 rounded">
      <p>{error}</p>

      {#if error === 'Invalid AI response, please try again'}
        <button
          on:click={handleSubmit}
          class="mt-2 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition"
        >
          Try again
        </button>
      {/if}
    </div>
  {/if}
</main>
