<script lang="ts">
  import { generateItinerary } from '$lib/api';
  let itinerary: any = null;
  let loading = false;
  let error: string | null = null;

  async function handleGenerate() {
    loading = true;
    error = null;
    itinerary = null;
    try {
      itinerary = await generateItinerary();
    } catch (err: any) {
      error = err.message;
    } finally {
      loading = false;
    }
  }
</script>

<div class="p-6 max-w-lg mx-auto">
  <h1 class="text-2xl font-bold mb-4">AI Itinerary Generator</h1>
  <button
    on:click={handleGenerate}
    class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
    disabled={loading}
  >
    {#if loading}
      Generating...
    {:else}
      Generate Itinerary
    {/if}
  </button>

  {#if error}
    <p class="text-red-600 mt-4">{error}</p>
  {/if}

  {#if itinerary}
    <div class="mt-6 bg-gray-100 p-4 rounded">
      <h2 class="text-lg font-semibold mb-2">Your Itinerary:</h2>
      {#each itinerary.days as day}
        <div class="mb-3">
          <p class="font-semibold">Day {day.day}</p>
          <ul class="list-disc ml-6">
            {#each day.plan as item}
              <li>{item}</li>
            {/each}
          </ul>
        </div>
      {/each}
    </div>
  {/if}
</div>
