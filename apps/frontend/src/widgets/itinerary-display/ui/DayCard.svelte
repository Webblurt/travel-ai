<script lang="ts">
  import Loader from '@/shared/ui/Loader.svelte'

  export let day: {
    morning: any
    afternoon: any
    evening: any
  }

  export let index: number
  export let onSelect: (section: 'morning' | 'afternoon' | 'evening', dayNum: number) => void

  let isLoading = false

  async function handleSelect(section: 'morning' | 'afternoon' | 'evening') {
    try {
      isLoading = true
      await onSelect(section, index + 1)
    } finally {
      isLoading = false
    }
  }
</script>

<div class="border border-gray/30 rounded-xl p-5 shadow-sm bg-white transition hover:shadow-md">
  {#if isLoading}
    <div class="absolute inset-0 bg-white/70 backdrop-blur-sm flex items-center justify-center z-10 rounded-xl">
      <Loader size={36} />
    </div>
  {/if}

  <h3 class="text-xl font-semibold mb-3 text-blue-600">
    Day {index + 1}
  </h3>

  <div class="grid gap-4 sm:grid-cols-3">
    <!-- Morning -->
    <div
      role="button"
      tabindex="0"
      on:click={() => handleSelect('morning')}
      on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && onSelect('morning', index + 1)}
      class="text-left p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition"
    >
      <h4 class="font-semibold text-gray-800 mb-1">Morning</h4>
      <p class="text-sm font-medium">{day.morning?.place || day.morning?.title}</p>
      <p class="text-xs text-gray-600">{day.morning?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
        {day.morning?.duration} | {day.morning?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.morning?.address}</p>
    </div>

    <!-- Afternoon -->
    <div
      role="button"
      tabindex="0"
      on:click={() => handleSelect('afternoon')}
      on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && onSelect('afternoon', index + 1)}
      class="text-left p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition"
    >
      <h4 class="font-semibold text-gray-800 mb-1">Afternoon</h4>
      <p class="text-sm font-medium">{day.afternoon?.place || day.afternoon?.title}</p>
      <p class="text-xs text-gray-600">{day.afternoon?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
        {day.afternoon?.duration} | {day.afternoon?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.afternoon?.address}</p>
    </div>

    <!-- Evening -->
    <div
      role="button"
      tabindex="0"
      on:click={() => handleSelect('evening')}
      on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && onSelect('evening', index + 1)}
      class="text-left p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition"
    >
      <h4 class="font-semibold text-gray-800 mb-1">Evening</h4>
      <p class="text-sm font-medium">{day.evening?.place || day.evening?.title}</p>
      <p class="text-xs text-gray-600">{day.evening?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
        {day.evening?.duration} | {day.evening?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.evening?.address}</p>
    </div>
  </div>
</div>
