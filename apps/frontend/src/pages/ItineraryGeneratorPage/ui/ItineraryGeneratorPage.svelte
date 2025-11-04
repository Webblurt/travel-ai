<script lang="ts">
  import { generateItinerary, selectItem, closeModal, result, loading, error, selectedItem } 
    from '@/features/generate-itinerary/model/useGenerateItinerary'
  import DayCard from '@/widgets/itinerary-display/ui/DayCard.svelte'
  import ItemModal from '@/widgets/itinerary-display/ui/ItemModal.svelte'

  let city = ''
  let days = ''
  let budget = ''
  let currency = ''

  function handleSubmit() {
    generateItinerary(city, days, budget, currency)
  }
</script>

<main class="p-6 max-w-3xl mx-auto">
  <h1 class="text-3xl font-bold mb-6 text-center text-blue-700">AI Travel Planner</h1>

  <!-- Form -->
  <form on:submit|preventDefault={handleSubmit} class="grid gap-4 bg-gray-50 p-6 rounded-xl shadow-sm">
    <input bind:value={city} placeholder="City" class="p-2 border rounded" />
    <input bind:value={days} placeholder="Number of days" class="p-2 border rounded" />
    <input bind:value={budget} placeholder="Budget (e.g. 1000)" class="p-2 border rounded" />
    <input bind:value={currency} placeholder="Currency (e.g. $)" class="p-2 border rounded" />
    <button
      type="submit"
      class="p-2 rounded bg-blue-600 text-white hover:bg-blue-700 transition disabled:opacity-50"
      disabled={$loading}
    >
      {$loading ? 'Generating...' : 'Generate itinerary'}
    </button>
  </form>

  <!-- Days card -->
  {#if $result}
    <div class="mt-6 grid gap-4">
      {#each $result.days as day, i}
        <DayCard {day} index={i} onSelect={selectItem} />
      {/each}
    </div>
  {/if}

  <!-- Modal  -->
  {#if $selectedItem}
    <ItemModal item={$selectedItem} onClose={closeModal} />
  {/if}

  <!-- Errors -->
  {#if $error}
    <div class="mt-4 bg-red-50 border border-red-300 text-red-700 p-4 rounded">
      <p>{$error}</p>

      {#if $error === 'Invalid AI response, please try again'}
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
