# Travel AI Platform — Monorepo

An AI-powered travel planning platform that generates structured, multi-day itineraries based on user preferences.

Built with:

- **SvelteKit** (frontend)
- **Golang** modular backend
- **Turborepo** monorepo architecture
- Workflow driven by **AI-assisted development**

---

## Release Notes

### v1.4.0 Active Development
Branch: `v1.4.0`

**Updates**
- swagger documentation for backend added

---

### v1.3.0 Active Development
Branch: `v1.3.0`

**Updates**
- Frontend refactored
- Improved UI & request workflow
- Updated AI response handling
- Better JSON validation & error feedback

---

### v1.2.0 - Stable Baseline
Branch: `v1.2.0`

**Features**
- PostgreSQL database added to save itineraries
- `make compose` used for full system launch
- Initial backend routes & AI integration

---

## Backend API Overview

API Base: `/api/v1`

###  Itinerary Routes

| Method | Endpoint | Description |
|--------|---------|-------------|
| `GET` | `/itinerary` | Generates a new AI itinerary |
| `GET` | `/itineraries` | Returns saved itineraries (supports filters) |
| `GET` | `/itineraries/:id` | Get itinerary by ID |
| `GET` | `/itineraries/:id/:dayNum` | Get specific itinerary day |
| `GET` | `/itineraries/:id/:dayNum/:section` | Get activity details (morning/afternoon/evening) |

### Filters for `/itineraries`

| Parameter | Description |
|----------|-------------|
| `city` | Filter by city |
| `budget` | Filter by budget |
| `page` | Pagination page |
| `per_page` | Items per page |

### Sections for day details
Available values:
morning,
afternoon,
evening
