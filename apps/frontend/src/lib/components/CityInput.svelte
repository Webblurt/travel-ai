<script>
  export let city = "";
  export let onSearch;
</script>

<div class="flex flex-col sm:flex-row gap-3 justify-center items-center my-10">
  <input
    bind:value={city}
    placeholder="Введите город..."
    class="border border-gray-400 rounded-xl px-4 py-3 w-72 text-gray-900 focus:outline-none focus:ring-2 focus:ring-blue-500"
  />
  <button
    on:click={() => onSearch?.(city)}
    class="bg-blue-600 hover:bg-blue-700 text-white font-medium px-6 py-3 rounded-xl transition"
  >
    Найти
  </button>
</div>
