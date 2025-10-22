<script lang="ts">
  import { onMount } from 'svelte';

  let city = '';
  let itinerary = '';
  let loading = false;
  let error = '';

  async function generateItinerary() {
    if (!city.trim()) {
      error = 'Введите город';
      return;
    }

    loading = true;
    error = '';
    itinerary = '';

    try {
      const res = await fetch('http://localhost:8080/api/itinerary', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ city }),
      });

      if (!res.ok) {
        throw new Error(`Ошибка: ${res.status}`);
      }

      const data = await res.json();
      itinerary = data.itinerary || JSON.stringify(data, null, 2);
    } catch (err) {
      error = 'Не удалось получить маршрут ';
      console.error(err);
    } finally {
      loading = false;
    }
  }
</script>

<main class="flex flex-col items-center justify-center min-h-screen bg-gray-50 text-gray-800 p-6">
  <h1 class="text-3xl font-bold mb-6 text-center"> AI Travel Planner</h1>

  <div class="bg-white shadow-md rounded-2xl p-6 w-full max-w-md">
    <label for="city" class="block mb-2 font-medium">Введите город</label>
    <input
      id="city"
      bind:value={city}
      placeholder="Например: Париж"
      class="w-full border border-gray-300 rounded-lg p-2 mb-4 focus:ring-2 focus:ring-blue-500"
    />

    <button
      on:click={generateItinerary}
      class="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition"
      disabled={loading}
    >
      {#if loading}
        ⏳ Генерация маршрута...
      {:else}
        ✈️ Сгенерировать маршрут
      {/if}
    </button>

    {#if error}
      <p class="text-red-500 mt-4">{error}</p>
    {/if}

    {#if itinerary}
      <div class="mt-6 bg-gray-100 p-4 rounded-lg whitespace-pre-wrap">
        {itinerary}
      </div>
    {/if}
  </div>
</main>
