<script lang="ts">
  export let item: any
  export let onClose: () => void
  import { formatMapEmbedUrl, handleKeyClose } from '@/shared/lib/utils'
</script>

<!-- BACKDROP -->
<div
  class="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
  role="button"
  tabindex="0"
  aria-hidden="true"
  on:click={onClose}
  on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && onClose()}
>
  <!-- MODAL CARD -->
  <div
    class="bg-white p-6 rounded-xl shadow-xl max-w-3xl w-full overflow-y-auto max-h-[90vh] border border-gray/30 hover:shadow-2xl transition"
    role="dialog"
    aria-modal="true"
    aria-label={item?.place ?? 'Details'}
    tabindex="0"
    on:click|stopPropagation
    on:keydown={handleKeyClose}
  >
    <!-- HEADER -->
    <div class="flex items-start justify-between">
      <div>
        <h2 class="text-2xl font-semibold text-blue-600 mb-1">{item?.place}</h2>
        <p class="text-sm text-gray-600">
          {item?.activity_type} • {item?.duration} • {item?.cost}
        </p>
        {#if item?.address}
          <p class="text-xs text-gray-500 mt-1">{item.address}</p>
        {/if}
      </div>

      <button
        type="button"
        class="ml-4 text-gray-400 hover:text-gray-600 transition"
        on:click={onClose}
        aria-label="Close"
      >
        ✕
      </button>
    </div>

    <!-- DESCRIPTION -->
    {#if item?.description}
      <p class="mt-4 text-gray-700 leading-relaxed">{item.description}</p>
    {/if}

    <!-- IMAGES -->
    {#if Array.isArray(item?.details?.images) && item.details.images.length > 0}
      <div class="grid grid-cols-2 gap-3 mt-4">
        {#each item.details.images as img}
          <figure class="rounded-lg overflow-hidden border border-gray-200 bg-gray-50 hover:shadow-sm transition">
            <img
              src={img.url}
              alt={img.description ?? item.place}
              class="object-cover w-full h-40"
            />
            {#if img.description}
              <figcaption class="text-xs text-gray-500 mt-1 px-2 py-1">
                {img.description}
              </figcaption>
            {/if}
          </figure>
        {/each}
      </div>
    {/if}

    <!-- OVERVIEW -->
    {#if item?.details?.overview}
      <section class="mt-5 p-3 bg-gray-50 rounded-lg border border-gray-100">
        <h3 class="font-semibold text-gray-800 mb-1">Overview</h3>
        {#if item.details.overview.short_description}
          <p class="text-sm text-gray-600">
            {item.details.overview.short_description}
          </p>
        {/if}
        {#if item.details.overview.history_or_background}
          <p class="text-xs text-gray-500 mt-1">
            {item.details.overview.history_or_background}
          </p>
        {/if}
        {#if Array.isArray(item.details.overview.interesting_facts) && item.details.overview.interesting_facts.length}
          <ul class="list-disc list-inside text-xs text-gray-600 mt-2">
            {#each item.details.overview.interesting_facts as fact}
              <li>{fact}</li>
            {/each}
          </ul>
        {/if}
      </section>
    {/if}

    <!-- VISITING INFO -->
    {#if item?.details?.visiting_info}
      <section class="mt-5 p-3 bg-gray-50 rounded-lg border border-gray-100">
        <h3 class="font-semibold text-gray-800 mb-1">Visiting info</h3>
        <div class="text-sm text-gray-600 grid gap-1">
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
                class="text-blue-600 underline hover:text-blue-800 ml-1"
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
                class="text-blue-600 underline hover:text-blue-800 ml-1"
              >
                {item.details.visiting_info.contact_phone}
              </a>
            </div>
          {/if}
        </div>
      </section>
    {/if}

    <!-- LOCATION INFO -->
    {#if item?.details?.location_info}
      <section class="mt-5 p-3 bg-gray-50 rounded-lg border border-gray-100">
        <h3 class="font-semibold text-gray-800 mb-1">Location</h3>
        <div class="text-sm text-gray-600 grid gap-1">
          {#if item.details.location_info.neighborhood}
            <div>Neighborhood: {item.details.location_info.neighborhood}</div>
          {/if}
          {#if item.details.location_info.nearest_transport}
            <div>Nearest transport: {item.details.location_info.nearest_transport}</div>
          {/if}
          {#if item.details.location_info.distance_from_city_center}
            <div>
              Distance from center: {item.details.location_info.distance_from_city_center}
            </div>
          {/if}
        </div>

        {#if item.details.location_info.map_link}
          <div class="mt-3 space-y-2">
            <iframe
              src={formatMapEmbedUrl(item.details.location_info.map_link)}
              class="w-full h-64 rounded-lg border border-gray-200"
              loading="lazy"
              allowfullscreen
              title={"Map of " + (item.place ?? "location")}
            ></iframe>
            <a
              href={item.details.location_info.map_link}
              target="_blank"
              rel="noopener noreferrer"
              class="text-blue-600 text-sm underline hover:text-blue-800"
            >
              Open in Google Maps
            </a>
          </div>
        {/if}
      </section>
    {/if}

    <!-- TIPS -->
    {#if item?.details?.tips}
      <section class="mt-5 p-3 bg-gray-50 rounded-lg border border-gray-100">
        <h3 class="font-semibold text-gray-800 mb-1">Tips</h3>

        {#if Array.isArray(item.details.tips.what_to_bring) && item.details.tips.what_to_bring.length}
          <div class="mt-2">
            <div class="font-medium text-sm">What to bring</div>
            <ul class="list-disc list-inside text-xs text-gray-600 mt-1">
              {#each item.details.tips.what_to_bring as t}
                <li>{t}</li>
              {/each}
            </ul>
          </div>
        {/if}

        {#if Array.isArray(item.details.tips.local_tips) && item.details.tips.local_tips.length}
          <div class="mt-2">
            <div class="font-medium text-sm">Local tips</div>
            <ul class="list-disc list-inside text-xs text-gray-600 mt-1">
              {#each item.details.tips.local_tips as t}
                <li>{t}</li>
              {/each}
            </ul>
          </div>
        {/if}

        {#if Array.isArray(item.details.tips.safety_notes) && item.details.tips.safety_notes.length}
          <div class="mt-2">
            <div class="font-medium text-sm">Safety notes</div>
            <ul class="list-disc list-inside text-xs text-gray-600 mt-1">
              {#each item.details.tips.safety_notes as t}
                <li>{t}</li>
              {/each}
            </ul>
          </div>
        {/if}
      </section>
    {/if}

    {#if !item?.details}
      <p class="mt-4 text-sm text-gray-500">No extra details available.</p>
    {/if}

    <!-- FOOTER -->
    <div class="mt-6 flex justify-end">
      <button
        type="button"
        class="px-4 py-2 rounded-lg bg-gray-50 hover:bg-gray-100 border border-gray-200 transition text-sm font-medium text-gray-700"
        on:click={onClose}
      >
        Close
      </button>
    </div>
  </div>
</div>
