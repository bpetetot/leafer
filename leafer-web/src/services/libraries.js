import useSWR, { mutate } from 'swr'
import { fetchJSON } from './utils'

const LIBRARIES = '/api/libraries'

export const useLibraries = () => useSWR(LIBRARIES)

export const useLibrary = (libraryId) => {
  return useSWR(libraryId ? [LIBRARIES, libraryId] : null, (url, libraryId) => {
    return fetchJSON(`${url}/${libraryId}`)
  })
}
export const addLibrary = (library) => {
  return mutate(LIBRARIES, async (libraries) => {
    const newLibrary = await fetchJSON(LIBRARIES, {
      method: 'POST',
      body: JSON.stringify(library),
      headers: { 'Content-Type': 'application/json' },
    })
    return [...libraries, newLibrary]
  })
}

export const removeLibrary = async (id) => {
  await fetchJSON(`${LIBRARIES}/${id}`, { method: 'DELETE' })
  mutate(LIBRARIES)
}
