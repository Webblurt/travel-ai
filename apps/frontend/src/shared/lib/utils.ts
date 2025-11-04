export function capitalize(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1)
}

export function delay(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

export function formatMapEmbedUrl(link: string): string {
    if (!link) return ''
    try {
      const url = new URL(link)
      const q = url.searchParams.get('q') || ''
      if (q) {
        return `https://www.google.com/maps?q=${encodeURIComponent(q)}&output=embed`
      }
      return `https://www.google.com/maps?q=${encodeURIComponent(link)}&output=embed`
    } catch {
      return `https://www.google.com/maps?q=${encodeURIComponent(link)}&output=embed`
    }
}

export let onClose: () => void;
export function handleKeyClose(e: KeyboardEvent) {
    if (e.key === 'Escape') {
      onClose();
    }
}
