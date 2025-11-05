<script>
  import "uno.css"
  import "@/app.css"
  import Footer from "@/shared/ui/Footer.svelte"
  import Header from "@/shared/ui/Header.svelte"

  let dark = false

  if (typeof window !== 'undefined') {
    const saved = localStorage.getItem('theme')
    if (saved) {
      dark = saved === 'dark'
    } else {
      dark = window.matchMedia('(prefers-color-scheme: dark)').matches
    }
  }

  $: {
    document.documentElement.classList.toggle('dark', dark)
    if (typeof window !== 'undefined') {
      localStorage.setItem('theme', dark ? 'dark' : 'light')
    }
  }
</script>

<div
  class="min-h-screen flex flex-col font-sans transition-colors duration-500
         bg-gradient-to-b from-blue-50 to-white text-gray-800
         dark:bg-gradient-to-b dark:from-gray-900 dark:to-gray-950 dark:text-gray-100"
>
  <Header />

  <div class="fixed top-4 right-4 z-30">
    <button
      class="p-2 rounded-full bg-white dark:bg-gray-800 shadow hover:shadow-md transition
             border border-gray-200 dark:border-gray-700"
      on:click={() => (dark = !dark)}
      title="Toggle dark mode"
    >
      {#if dark}
        🌙
      {:else}
        ☀️
      {/if}
    </button>
  </div>

  <main class="flex-1 max-w-5xl w-full mx-auto px-4 sm:px-6 lg:px-8 py-10 animate-fade-in">
    <slot />
  </main>

  <Footer />
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
