<script lang="ts">
  export let day: {
    day_number: number
    morning: any
    afternoon: any
    evening: any
  }

  export let onSelect: (section: 'morning' | 'afternoon' | 'evening', dayNum: number) => void | Promise<void>
  let loadingSection: 'morning' | 'afternoon' | 'evening' | null = null

  async function withLoading(section: 'morning' | 'afternoon' | 'evening', dayNum: number) {
    if (loadingSection) return
    loadingSection = section
    try {
      const maybePromise = onSelect(section, dayNum)
      if (maybePromise instanceof Promise) {
        await maybePromise
      }
    } catch (err) {
      console.error(err)
    } finally {
      loadingSection = null
    }
  }
</script>

<div class="border border-gray/30 rounded-xl p-5 shadow-sm bg-white transition hover:shadow-md">
  <h3 class="text-xl font-semibold mb-3 text-blue-600">
    Day {day.day_number}
  </h3>

  <div class="grid gap-4 sm:grid-cols-3">
    <div
      class="relative p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100"
      role="button"
      tabindex="0"
      on:click={() => withLoading('morning', day.day_number)}
      on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && withLoading('morning', day.day_number)}
    >
      <h4 class="font-semibold text-gray-800 mb-1">Morning</h4>
      <p class="text-sm font-medium">{day.morning?.place}</p>
      <p class="text-xs text-gray-600">{day.morning?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
        {day.morning?.duration} | {day.morning?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.morning?.address}</p>

      {#if loadingSection === 'morning'}
        <div class="absolute inset-0 bg-white/70 flex items-center justify-center rounded-lg">
          <svg
            class="animate-spin h-6 w-6 text-blue-600"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              class="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              stroke-width="4"
            ></circle>
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"
            ></path>
          </svg>
        </div>
      {/if}
    </div>
    <div
      class="relative p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100"
      role="button"
      tabindex="0"
      on:click={() => withLoading('afternoon', day.day_number)}
      on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && withLoading('afternoon', day.day_number)}
    >
      <h4 class="font-semibold text-gray-800 mb-1">Afternoon</h4>
      <p class="text-sm font-medium">{day.afternoon?.place}</p>
      <p class="text-xs text-gray-600">{day.afternoon?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
        {day.afternoon?.duration} | {day.afternoon?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.afternoon?.address}</p>

      {#if loadingSection === 'afternoon'}
        <div class="absolute inset-0 bg-white/70 flex items-center justify-center rounded-lg">
          <svg
            class="animate-spin h-6 w-6 text-blue-600"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              class="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              stroke-width="4"
            ></circle>
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"
            ></path>
          </svg>
        </div>
      {/if}
    </div>
    <div
      class="relative p-3 bg-gray-50 rounded-lg cursor-pointer hover:bg-gray-100"
      role="button"
      tabindex="0"
      on:click={() => withLoading('evening', day.day_number)}
      on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && withLoading('evening', day.day_number)}
    >
      <h4 class="font-semibold text-gray-800 mb-1">Evening</h4>
      <p class="text-sm font-medium">{day.evening?.place}</p>
      <p class="text-xs text-gray-600">{day.evening?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
        {day.evening?.duration} | {day.evening?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.evening?.address}</p>

      {#if loadingSection === 'evening'}
        <div class="absolute inset-0 bg-white/70 flex items-center justify-center rounded-lg">
          <svg
            class="animate-spin h-6 w-6 text-blue-600"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              class="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              stroke-width="4"
            ></circle>
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"
            ></path>
          </svg>
        </div>
      {/if}
    </div>
  </div>
</div>
