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

<div
  class="relative border border-gray-200 dark:border-gray-700 rounded-2xl p-6 shadow-sm bg-white dark:bg-gray-800
         transition-all duration-300 hover:shadow-lg dark:hover:shadow-gray-900/60 animate-fade-in"
>
  {#if isLoading}
    <div class="absolute inset-0 bg-white/70 dark:bg-gray-900/70 backdrop-blur-sm flex items-center justify-center z-10 rounded-2xl">
      <Loader size={36} />
    </div>
  {/if}

  <h3 class="text-2xl font-semibold mb-4 text-blue-600 dark:text-blue-400 text-center">
    Day {index + 1}
  </h3>

  <div class="grid gap-4 sm:grid-cols-3">
    {#each Object.entries(day).filter(([key]) => ['morning', 'afternoon', 'evening'].includes(key)) as [section, content]}
      <div
        role="button"
        tabindex="0"
        on:click={() => handleSelect(section as 'morning' | 'afternoon' | 'evening')}
        on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && handleSelect(section as any)}
        class="text-left p-4 rounded-xl border border-gray-100 dark:border-gray-700 bg-gray-50 dark:bg-gray-900/60
              hover:(bg-gray-100 dark:bg-gray-900/80) transition-all duration-300"
      >
        <h4 class="font-semibold text-gray-800 dark:text-gray-200 mb-1 capitalize">
          {section}
        </h4>
        <p class="text-sm font-medium text-gray-900 dark:text-gray-100">{content?.place || content?.title}</p>
        <p class="text-xs text-gray-600 dark:text-gray-400 mt-1">{content?.description}</p>
        <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
          {content?.duration} | {content?.cost}
        </p>
        <p class="text-xs text-gray-400 dark:text-gray-500">{content?.address}</p>
      </div>
    {/each}
  </div>
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
