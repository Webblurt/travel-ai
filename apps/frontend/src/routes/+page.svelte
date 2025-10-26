<script>
  import Header from "$lib/components/Header.svelte";
  import CityInput from "$lib/components/CityInput.svelte";
  import ResponseBlock from "$lib/components/ResponseBlock.svelte";

  let city = "";
  let response = "";
  let isLoading = false;

  const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

  async function handleSearch(cityName) {
    if (!cityName) return;
    isLoading = true;
    response = "";

    try {
      const res = await fetch(`${API_URL}/api/ai?city=${encodeURIComponent(cityName)}`);
      if (!res.ok) throw new Error("Ошибка ответа от API");
      const data = await res.json();
      response = data.message || "Пустой ответ";
    } catch (e) {
      response = "Ошибка при запросе к серверу";
      console.error(e);
    } finally {
      isLoading = false;
    }
  }
</script>

<main class="min-h-screen bg-gradient-to-br from-gray-900 via-blue-900 to-indigo-900 text-white flex flex-col">
  <Header />
  <div class="flex-grow flex flex-col items-center justify-center px-4">
    <CityInput bind:city onSearch={handleSearch} />
    <ResponseBlock text={response} {isLoading} />
  </div>
</main>
