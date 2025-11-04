<script lang="ts">
  export let onSubmit: (data: {
    city: string
    days: string
    budget: string
    currency: string
  }) => void

  let city = ''
  let days = ''
  let budget = ''
  let currency = ''
</script>

<form on:submit|preventDefault={() => onSubmit({ city, days, budget, currency })} class="grid gap-4 bg-gray-50 p-6 rounded-xl shadow-sm">
  <input bind:value={city} placeholder="City" class="p-2 border rounded" />
  <input bind:value={days} placeholder="Number of days" class="p-2 border rounded" />
  <input bind:value={budget} placeholder="Budget (e.g. 1000)" class="p-2 border rounded" />
  <input bind:value={currency} placeholder="Currency (e.g. $)" class="p-2 border rounded" />
  <button type="submit" class="p-2 rounded bg-blue-600 text-white hover:bg-blue-700 transition">
    Generate itinerary
  </button>
</form>
