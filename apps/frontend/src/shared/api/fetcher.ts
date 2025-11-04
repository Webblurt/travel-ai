export async function fetcher(url: string, key: string, options: RequestInit = {}) {
  const res = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': key,
      ...(options.headers || {})
    }
  })

  const text = await res.text()
  if (!res.ok) {
    throw new Error(`HTTP ${res.status}: ${text}`)
  }

  try {
    return JSON.parse(text)
  } catch {
    return text
  }
}
