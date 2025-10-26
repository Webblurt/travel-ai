<script>
  export let text = "";
  export let isLoading = false;
</script>

{#if isLoading}
  <div class="text-center text-gray-400 animate-pulse mt-6">
    Обработка запроса...
  </div>
{:else if text}
  <div class="max-w-2xl mx-auto bg-white/10 p-6 rounded-xl border border-white/20 shadow-lg text-center mt-6">
    <p class="text-lg leading-relaxed text-gray-100 whitespace-pre-wrap">
      {text}
    </p>
  </div>
{:else}
  <div class="text-gray-400 text-center italic mt-6">
    Введите город, чтобы получить рекомендации
  </div>
{/if}
