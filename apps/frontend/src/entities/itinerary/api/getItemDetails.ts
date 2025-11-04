import { fetcher } from '@/shared/api/fetcher'

export async function getItemDetails(id: string, dayNum: number, section: string) {
  return await fetcher(`${import.meta.env.VITE_API_URL}/itineraries/${id}/${dayNum}/${section}`, `${import.meta.env.VITE_API_TOKEN}`)
}
