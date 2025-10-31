<script lang="ts">
  export let day: {
    day_number: number
    morning: any
    day: any
    evening: any
  }
</script>

<div class="border border-gray/30 rounded-xl p-5 shadow-sm bg-white transition hover:shadow-md">
  <h3 class="text-xl font-semibold mb-3 text-blue-600">
    Day {day.day_number}
  </h3>

  <div class="grid gap-4 sm:grid-cols-3">
    <div class="p-3 bg-gray-50 rounded-lg">
      <h4 class="font-semibold text-gray-800 mb-1"> Morning</h4>
      <p class="text-sm font-medium">{day.morning?.place}</p>
      <p class="text-xs text-gray-600">{day.morning?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
         {day.morning?.duration} |  {day.morning?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.morning?.address}</p>
    </div>

    <div class="p-3 bg-gray-50 rounded-lg">
      <h4 class="font-semibold text-gray-800 mb-1"> Day</h4>
      <p class="text-sm font-medium">{day.day?.place}</p>
      <p class="text-xs text-gray-600">{day.day?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
         {day.day?.duration} |  {day.day?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.day?.address}</p>
    </div>

    <div class="p-3 bg-gray-50 rounded-lg">
      <h4 class="font-semibold text-gray-800 mb-1"> Evening</h4>
      <p class="text-sm font-medium">{day.evening?.place}</p>
      <p class="text-xs text-gray-600">{day.evening?.description}</p>
      <p class="text-xs text-gray-500 mt-1">
         {day.evening?.duration} |  {day.evening?.cost}
      </p>
      <p class="text-xs text-gray-400">{day.evening?.address}</p>
    </div>
  </div>
</div>
