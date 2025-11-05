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

  <div class="max-w-4xl w-full mx-auto flex flex-col items-center text-center transition-all duration-700 animate-fade-in">
    <h1 class="text-4xl font-bold text-blue-600 dark:text-blue-400 mb-8">
      Welcome to AI Travel Planner
    </h1>

    <!-- Form -->
    <form
      on:submit|preventDefault={handleSubmit}
      class="grid gap-4 bg-white dark:bg-gray-800 shadow-md p-8 rounded-2xl
         border border-gray-100 dark:border-gray-700 w-full max-w-2xl mx-auto"
    >
      <input
        bind:value={city}
        placeholder="City"
        class="p-3 border border-gray-300 dark:border-gray-600
               bg-white dark:bg-gray-700 rounded-xl
               text-gray-900 dark:text-gray-100
               focus:(outline-none ring-2 ring-blue-400 dark:ring-blue-500)"
      />
      <input
        bind:value={days}
        placeholder="Number of days"
        class="p-3 border border-gray-300 dark:border-gray-600
               bg-white dark:bg-gray-700 rounded-xl
               text-gray-900 dark:text-gray-100
               focus:(outline-none ring-2 ring-blue-400 dark:ring-blue-500)"
      />
      <input
        bind:value={budget}
        placeholder="Budget (e.g. 1000)"
        class="p-3 border border-gray-300 dark:border-gray-600
               bg-white dark:bg-gray-700 rounded-xl
               text-gray-900 dark:text-gray-100
               focus:(outline-none ring-2 ring-blue-400 dark:ring-blue-500)"
      />
      <input
        bind:value={currency}
        placeholder="Currency (e.g. $)"
        class="p-3 border border-gray-300 dark:border-gray-600
               bg-white dark:bg-gray-700 rounded-xl
               text-gray-900 dark:text-gray-100
               focus:(outline-none ring-2 ring-blue-400 dark:ring-blue-500)"
      />

      <button
        type="submit"
        class="py-3 rounded-xl text-white font-semibold
               bg-blue-600 hover:bg-blue-700 dark:(bg-blue-500 hover:bg-blue-600)
               transition-all duration-300 shadow-sm hover:shadow-md
               disabled:(opacity-50 cursor-not-allowed)"
        disabled={$loading}
      >
        {$loading ? 'Generating...' : 'Generate Itinerary'}
      </button>
    </form>

    <!-- Days card -->
    {#if $result}
      <div class="mt-8 grid gap-6 w-full max-w-3xl mx-auto animate-fade-in">
        {#each $result.days as day, i}
          <DayCard {day} index={i} onSelect={selectItem} />
        {/each}
      </div>
    {/if}

    <!-- Modal -->
    {#if $selectedItem}
      <ItemModal item={$selectedItem} onClose={closeModal} />
    {/if}

    <!-- Errors -->
    {#if $error}
      <div class="mt-6 bg-red-50 dark:bg-red-900/30 border border-red-300 dark:border-red-700
                  text-red-700 dark:text-red-400 p-4 rounded-xl">
        <p>{$error}</p>

        {#if $error === 'Invalid AI response, please try again'}
          <button
            on:click={handleSubmit}
            class="mt-3 px-5 py-2 bg-blue-600 text-white rounded-lg
                   hover:bg-blue-700 dark:(bg-blue-500 hover:bg-blue-600) transition"
          >
            Try again
          </button>
        {/if}
      </div>
    {/if}
  </div>


<style>
  @keyframes fade-in {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
  }
  .animate-fade-in {
    animation: fade-in 0.6s ease-out forwards;
  }
</style>
