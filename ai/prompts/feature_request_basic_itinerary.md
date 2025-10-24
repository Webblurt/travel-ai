# Feature Request - Generate Travel Itinerary

Goal:
User can click "Generate itinerary" in frontend and receive a day-by-day travel plan from the backend using AI.

Context:
- Frontend is SvelteKit, backend is Go.
- Communication via REST API: POST /api/itinerary
- AI should generate itinerary data in a JSON format like:
  {
    "days": [
      { "day": 1, "plan": [ { "activity": "...", "time": "..." } ] }
    ]
  }

Requirements when asking AI to generate code:
- Backend must have handler in `internal/http/itinerary.go`
- Handler must validate request and return mock itinerary for now.
- Frontend should have a SvelteKit page `/itinerary` with a button to trigger generation.

Reminder for AI:
- Always write idiomatic Go (modular, no logic in main.go).
- Always write SvelteKit with `<script lang="ts">`.
- Use Tailwind for layout.
