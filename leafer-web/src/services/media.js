import useSWR from 'swr'
import { fetchBase64, fetchJSON } from './utils'

const MEDIA = '/api/media'

export const useMedia = (mediaId) => {
  return useSWR([MEDIA, mediaId], (url, mediaId) => {
    return fetchJSON(`${url}/${mediaId}`)
  })
}

export const useMediaLibrary = (libraryId) => {
  return useSWR([MEDIA, libraryId], (url, libraryId) => {
    return fetchJSON(`${url}?libraryId=${libraryId}&parentMediaId=0`)
  })
}

export const fetchMediaByIndex = async (libraryId, collectionId, mediaIndex) => {
  const url = `${MEDIA}?libraryId=${libraryId}&parentMediaId=${collectionId}&mediaIndex=${mediaIndex}`
  const media = await fetchJSON(url)
  return media?.data?.[0]
}

export const fetchMediaPage = async (mediaId, pageIndex) => {
  return fetchBase64(`${MEDIA}/${mediaId}/content?page=${pageIndex}`)
}
