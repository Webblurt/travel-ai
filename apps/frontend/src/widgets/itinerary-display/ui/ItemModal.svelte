<script lang="ts">
  import { onMount } from 'svelte';
  import { formatMapEmbedUrl, handleKeyClose } from '@/shared/lib/utils';

  export let item: any;
  export let onClose: () => void;

  let fadeIn = false;
  onMount(() => (fadeIn = true));
</script>

<!-- BACKDROP -->
<div
  class="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
  on:click={onClose}
  on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && onClose()}
  tabindex="0"
  role="button"
  aria-hidden="true"
>
  <!-- MODAL -->
  <div
    class={`transition-all duration-700 ease-out transform ${
      fadeIn ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'
    } bg-white dark:bg-gray-900 shadow-xl dark:shadow-gray-900/50 
       border border-gray-100 dark:border-gray-700 rounded-2xl 
       p-6 max-w-3xl w-full max-h-[90vh] overflow-y-auto`}
    role="dialog"
    tabindex="0"
    aria-modal="true"
    aria-label={item?.place ?? 'Details'}
    on:click|stopPropagation
    on:keydown={handleKeyClose}
  >
    <!-- HEADER -->
    <div class="flex items-start justify-between mb-4">
      <div>
        <h2 class="text-2xl font-bold text-blue-600 dark:text-blue-400 mb-1">
          {item?.place}
        </h2>
        <p class="text-sm text-gray-600 dark:text-gray-400">
          {item?.activity_type} • {item?.duration} • {item?.cost}
        </p>
        {#if item?.address}
          <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">{item.address}</p>
        {/if}
      </div>

      <button
        type="button"
        class="ml-4 text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 transition text-xl"
        on:click={onClose}
        aria-label="Close"
      >
        ✕
      </button>
    </div>

    <!-- DESCRIPTION -->
    {#if item?.description}
      <p class="mt-2 text-gray-700 dark:text-gray-300 leading-relaxed">
        {item.description}
      </p>
    {/if}

    <!-- OVERVIEW -->
    {#if item?.details?.overview}
      <section class="mt-5 p-3 bg-gray-50 dark:bg-gray-800 rounded-lg border border-gray-100 dark:border-gray-700">
        <h3 class="font-semibold text-gray-800 dark:text-gray-100 mb-1">Overview</h3>
        {#if item.details.overview.short_description}
          <p class="text-sm text-gray-600 dark:text-gray-300">
            {item.details.overview.short_description}
          </p>
        {/if}
        {#if item.details.overview.history_or_background}
          <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
            {item.details.overview.history_or_background}
          </p>
        {/if}
        {#if Array.isArray(item.details.overview.interesting_facts) && item.details.overview.interesting_facts.length}
          <ul class="list-disc list-inside text-xs text-gray-600 dark:text-gray-400 mt-2">
            {#each item.details.overview.interesting_facts as fact}
              <li>{fact}</li>
            {/each}
          </ul>
        {/if}
      </section>
    {/if}

    <!-- IMAGES -->
    {#if Array.isArray(item?.details?.images) && item.details.images.length > 0}
      <div class="grid grid-cols-2 gap-3 mt-5">
        {#each item.details.images as img}
          <figure
            class="bg-white dark:bg-gray-800 rounded-xl overflow-hidden 
                   border border-gray-100 dark:border-gray-700 
                   shadow-sm hover:shadow-md dark:shadow-gray-900/40 transition"
          >
            <img src={img.url} alt={img.description ?? item.place} class="object-cover w-full h-40" />
            {#if img.description}
              <figcaption class="text-xs text-gray-600 dark:text-gray-400 px-2 py-1">
                {img.description}
              </figcaption>
            {/if}
          </figure>
        {/each}
      </div>
    {/if}

    <!-- VISITING INFO -->
    {#if item?.details?.visiting_info}
      <section class="mt-5 p-3 bg-gray-50 dark:bg-gray-800 rounded-lg border border-gray-100 dark:border-gray-700">
        <h3 class="font-semibold text-gray-800 dark:text-gray-100 mb-1">Visiting info</h3>
        <div class="text-sm text-gray-600 dark:text-gray-300 grid gap-1">
          {#if item.details.visiting_info.opening_hours}
            <div>Hours: {item.details.visiting_info.opening_hours}</div>
          {/if}
          {#if item.details.visiting_info.best_time_to_visit}
            <div>Best time: {item.details.visiting_info.best_time_to_visit}</div>
          {/if}
          {#if item.details.visiting_info.average_visit_duration}
            <div>Average duration: {item.details.visiting_info.average_visit_duration}</div>
          {/if}
          {#if item.details.visiting_info.entry_fee}
            <div>Entry fee: {item.details.visiting_info.entry_fee}</div>
          {/if}
          {#if item.details.visiting_info.website}
            <div>
              Website:
              <a
                href={item.details.visiting_info.website}
                class="text-blue-600 dark:text-blue-400 underline hover:text-blue-800 dark:hover:text-blue-300 ml-1"
                target="_blank"
                rel="noopener noreferrer"
              >
                {item.details.visiting_info.website}
              </a>
            </div>
          {/if}
          {#if item.details.visiting_info.contact_phone}
            <div>
              Phone:
              <a
                href={"tel:" + item.details.visiting_info.contact_phone}
                class="text-blue-600 dark:text-blue-400 underline hover:text-blue-800 dark:hover:text-blue-300 ml-1"
              >
                {item.details.visiting_info.contact_phone}
              </a>
            </div>
          {/if}
        </div>
      </section>
    {/if}

    <!-- TIPS -->
    {#if item?.details?.tips}
      <section class="mt-5 p-3 bg-gray-50 dark:bg-gray-800 rounded-lg border border-gray-100 dark:border-gray-700">
        <h3 class="font-semibold text-gray-800 dark:text-gray-100 mb-1">Tips</h3>

        {#if Array.isArray(item.details.tips.what_to_bring) && item.details.tips.what_to_bring.length}
          <div class="mt-2">
            <div class="font-medium text-sm text-gray-700 dark:text-gray-200">What to bring</div>
            <ul class="list-disc list-inside text-xs text-gray-600 dark:text-gray-400 mt-1">
              {#each item.details.tips.what_to_bring as t}
                <li>{t}</li>
              {/each}
            </ul>
          </div>
        {/if}

        {#if Array.isArray(item.details.tips.local_tips) && item.details.tips.local_tips.length}
          <div class="mt-2">
            <div class="font-medium text-sm text-gray-700 dark:text-gray-200">Local tips</div>
            <ul class="list-disc list-inside text-xs text-gray-600 dark:text-gray-400 mt-1">
              {#each item.details.tips.local_tips as t}
                <li>{t}</li>
              {/each}
            </ul>
          </div>
        {/if}

        {#if Array.isArray(item.details.tips.safety_notes) && item.details.tips.safety_notes.length}
          <div class="mt-2">
            <div class="font-medium text-sm text-gray-700 dark:text-gray-200">Safety notes</div>
            <ul class="list-disc list-inside text-xs text-gray-600 dark:text-gray-400 mt-1">
              {#each item.details.tips.safety_notes as t}
                <li>{t}</li>
              {/each}
            </ul>
          </div>
        {/if}
      </section>
    {/if}

    <!-- LOCATION -->
    {#if item?.details?.location_info}
      <section class="mt-5 p-3 bg-gray-50 dark:bg-gray-800 rounded-lg border border-gray-100 dark:border-gray-700">
        <h3 class="font-semibold text-gray-800 dark:text-gray-100 mb-1">Location</h3>
        <div class="text-sm text-gray-600 dark:text-gray-300 grid gap-1">
          {#if item.details.location_info.neighborhood}
            <div>Neighborhood: {item.details.location_info.neighborhood}</div>
          {/if}
          {#if item.details.location_info.nearest_transport}
            <div>Nearest transport: {item.details.location_info.nearest_transport}</div>
          {/if}
          {#if item.details.location_info.distance_from_city_center}
            <div>Distance from center: {item.details.location_info.distance_from_city_center}</div>
          {/if}
        </div>

        {#if item.details.location_info.map_link}
          <div class="mt-3 space-y-2">
            <iframe
              src={formatMapEmbedUrl(item.details.location_info.map_link)}
              class="w-full h-64 rounded-lg border border-gray-200 dark:border-gray-700"
              loading="lazy"
              allowfullscreen
              title={"Map of " + (item.place ?? "location")}
            ></iframe>
            <a
              href={item.details.location_info.map_link}
              target="_blank"
              rel="noopener noreferrer"
              class="text-blue-600 dark:text-blue-400 text-sm underline hover:text-blue-800 dark:hover:text-blue-300"
            >
              Open in Google Maps
            </a>
          </div>
        {/if}
      </section>
    {/if}

  </div>
</div>
