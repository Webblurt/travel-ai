<script>
  import { onMount } from 'svelte';

  let fadeIn = false;
  onMount(() => (fadeIn = true));
</script>

<div
  class={`transition-all duration-1000 ease-out transform ${
    fadeIn ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'
  } flex flex-col items-center text-center px-6 sm:px-8 lg:px-12`}
>
  <h1 class="text-4xl font-bold text-blue-600 dark:text-blue-400 mb-6">
    About Travel AI
  </h1>

  <p class="max-w-3xl text-gray-700 dark:text-gray-300 text-lg leading-relaxed mb-10">
    Welcome to <span class="font-semibold text-blue-500 dark:text-blue-400">Travel AI</span> - your intelligent trip planner.
    Our app creates personalized itineraries in seconds using the power of artificial intelligence.
  </p>

  <div class="grid sm:grid-cols-3 gap-6 w-full max-w-5xl mb-14">
    <div
      class="bg-white dark:bg-gray-800 shadow-md dark:shadow-gray-900/40
             p-8 rounded-2xl hover:shadow-lg dark:hover:shadow-gray-800/70
             transition-all duration-300 border border-gray-100 dark:border-gray-700"
    >
      <h2 class="text-xl font-semibold text-blue-500 dark:text-blue-400 mb-3">
        Plan Instantly
      </h2>
      <p class="text-gray-600 dark:text-gray-300 text-base">
        Enter your city, travel days, and budget - and get a full itinerary designed just for you.
      </p>
    </div>

    <div
      class="bg-white dark:bg-gray-800 shadow-md dark:shadow-gray-900/40
             p-8 rounded-2xl hover:shadow-lg dark:hover:shadow-gray-800/70
             transition-all duration-300 border border-gray-100 dark:border-gray-700"
    >
      <h2 class="text-xl font-semibold text-blue-500 dark:text-blue-400 mb-3">
        AI-Powered
      </h2>
      <p class="text-gray-600 dark:text-gray-300 text-base">
        Our smart system curates the best attractions, restaurants, and experiences for your style.
      </p>
    </div>

    <div
      class="bg-white dark:bg-gray-800 shadow-md dark:shadow-gray-900/40
             p-8 rounded-2xl hover:shadow-lg dark:hover:shadow-gray-800/70
             transition-all duration-300 border border-gray-100 dark:border-gray-700"
    >
      <h2 class="text-xl font-semibold text-blue-500 dark:text-blue-400 mb-3">
        Save Time
      </h2>
      <p class="text-gray-600 dark:text-gray-300 text-base">
        No more hours of research - your journey is ready in moments, tailored to your budget and taste.
      </p>
    </div>
  </div>
</div>
