<input
  type="text"
  placeholder="Enter destination..."
  bind:value={destination}
  class="border p-2 rounded w-full mb-2"
/>

<script lang="ts">
  import { generateItinerary } from '$lib/api';
  let destination = 'Paris';
  let itinerary = null;
  let loading = false;
  let error = null;

  async function handleGenerate() {
    loading = true;
    itinerary = null;
    error = null;
    try {
      itinerary = await generateItinerary(destination);
    } catch (err: any) {
      error = err.message;
    } finally {
      loading = false;
    }
  }
</script>
