<script lang="ts">
  import { fade, fly } from 'svelte/transition';
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
      if (!res.ok) throw new Error(`Ошибка: ${res.status}`);
      const data = await res.json();
      itinerary = data.itinerary || JSON.stringify(data, null, 2);
    } catch (e) {
      error = 'Ошибка при получении маршрута';
      console.error(e);
    } finally {
      loading = false;
    }
  }

  function splitDays(text: string) {
    const days = text.match(/(День\s*\d+[\s\S]*?)(?=День\s*\d+|$)/gi);
    return days ? days.map((d) => d.trim()) : [text.trim()];
  }
</script>

<main class="min-h-screen bg-gradient-to-br from-sky-100 to-blue-200 flex flex-col items-center justify-center p-6">
  <h1 class="text-4xl font-bold mb-8 text-sky-800 text-center" in:fly={{ y: -20, opacity: 0, duration: 500 }}>
    AI Travel Planner
  </h1>

  <div class="bg-white/80 backdrop-blur-md shadow-xl rounded-3xl p-8 w-full max-w-lg" in:fly={{ y: 20, opacity: 0, duration: 500, delay: 200 }}>
    <label for="city" class="block mb-2 text-gray-700 font-semibold">
      Введите город:
    </label>
    <input
      id="city"
      bind:value={city}
      placeholder="Например: Барселона"
      class="w-full border border-gray-300 rounded-xl p-3 mb-4 focus:ring-2 focus:ring-sky-500 focus:outline-none"
    />

    <button
      on:click={generateItinerary}
      disabled={loading}
      class="w-full bg-sky-600 text-white font-semibold py-3 rounded-xl hover:bg-sky-700 transition disabled:opacity-60"
    >
      {#if loading} Генерация маршрута... {:else} Сгенерировать маршрут {/if}
    </button>

    {#if error}
      <p class="text-red-500 mt-4 text-center" in:fade>{{error}}</p>
    {/if}

    {#if itinerary}
      <div class="mt-6 space-y-4" in:fade={{ delay: 200 }}>
        {#each splitDays(itinerary) as day, i}
          <div class="bg-white shadow-md rounded-2xl p-5 border border-gray-100 hover:shadow-lg transition" in:fly={{ y: 10, duration: 300 }}>
            <h2 class="font-bold text-xl text-sky-700 mb-3">
              День {i + 1}
            </h2>
            <div class="text-gray-700 whitespace-pre-wrap leading-relaxed">
              {@html day
                .replace(/День\s*\d+/i, '')
                .replace(/\*\*(.*?)\*\*/g, '<b>$1</b>')}
            </div>
          </div>
        {/each}
      </div>
    {/if}
  </div>

  <footer class="mt-10 text-gray-600 text-sm">
    © webblurt 2025
  </footer>
</main>
