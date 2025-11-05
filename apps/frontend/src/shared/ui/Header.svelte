<script lang="ts">
  let menuOpen = false

  const links = [
    { href: '/', label: 'Generate Trip' },
    { href: '/itineraries', label: 'Saved Trips' },
    { href: '/about', label: 'About' },
  ]
</script>

<header class="sticky top-0 z-30 border-b border-gray-200/70 dark:border-gray-800/70 backdrop-blur-md bg-white/80 dark:bg-gray-900/80 transition-all duration-300">
  <div class="max-w-6xl mx-auto px-6 py-4 flex items-center justify-between">
    <a href="/" class="text-2xl font-semibold text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300 transition">
      Travel<span class="text-gray-900 dark:text-gray-100">AI</span>
    </a>

    <!-- Desktop Menu -->
    <nav class="hidden md:flex items-center gap-12 font-medium">
      {#each links as link}
        <a
            href={link.href}
            class="relative inline-block whitespace-nowrap px-1 pb-1
                    text-gray-700 dark:text-gray-300 no-underline transition
                    hover:(text-blue-600 dark:text-blue-400)
                    hover:after:(w-full)"
            >
            {link.label}
        </a>
      {/each}
    </nav>

    <button
      class="md:hidden text-gray-700 dark:text-gray-300 text-2xl"
      on:click={() => (menuOpen = !menuOpen)}
      aria-label="Toggle menu"
    >
      {#if menuOpen} ✖ {:else} ☰ {/if}
    </button>
  </div>

  {#if menuOpen}
    <nav class="md:hidden bg-white/90 dark:bg-gray-900/90 border-t border-gray-100 dark:border-gray-800 backdrop-blur-md">
      {#each links as link}
        <a href={link.href} class="block px-6 py-3 text-gray-700 dark:text-gray-300 border-b border-gray-100 dark:border-gray-800 hover:(bg-blue-50 dark:bg-gray-800 text-blue-600 dark:text-blue-400) transition">
          {link.label}
        </a>
      {/each}
    </nav>
  {/if}
</header>

<style>
  a { text-decoration: none; }
</style>
