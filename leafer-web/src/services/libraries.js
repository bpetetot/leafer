import useSWR, { mutate } from 'swr'
import { fetchJSON } from './utils'

const LIBRARIES = '/api/libraries'

export const useLibraries = () => useSWR(LIBRARIES)

export const useLibrary = (libraryId) => {
  return useSWR(libraryId ? [LIBRARIES, libraryId] : null, (url, libraryId) => {
    return fetchJSON(`${url}/${libraryId}`)
  })
}
export const addLibrary = (data) => {
  return mutate(LIBRARIES, async (libraries = []) => {
    const library = await fetchJSON(LIBRARIES, {
      method: 'POST',
      body: JSON.stringify(data),
      headers: { 'Content-Type': 'application/json' },
    })
    return [...libraries, library]
  })
}

export const removeLibrary = async (id) => {
  return mutate(LIBRARIES, async (libraries = []) => {
    await fetchJSON(`${LIBRARIES}/${id}`, { method: 'DELETE' })
    return libraries.filter(library => library.id !== id)
  })
}

export const scanLibrary = async (id) => {
  await fetchJSON(`${LIBRARIES}/${id}/scan`)
}
