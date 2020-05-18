import useSWR from 'swr'
import { fetchBase64, fetchJSON } from './utils'

const MEDIA = '/api/media'

export const useMedia = (mediaId) => {
  return useSWR(mediaId ? [MEDIA, mediaId] : null, (url, mediaId) => {
    return fetchJSON(`${url}/${mediaId}`)
  })
}

export const useMediasLibrary = (libraryId) => {
  return useSWR(libraryId ? [MEDIA, libraryId] : null, (url, libraryId) => {
    return fetchJSON(`${url}?libraryId=${libraryId}&parentMediaId=0`)
  })
}

export const useMediasCollection = (libraryId, mediaCollection) => {
  const shouldFetch = mediaCollection && mediaCollection.type === 'COLLECTION'
  return useSWR(
    shouldFetch ? [MEDIA, libraryId, mediaCollection.id] : null,
    (url, libraryId, parentMediaId) => {
      return fetchJSON(
        `${url}?libraryId=${libraryId}&parentMediaId=${parentMediaId}`
      )
    }
  )
}

export const fetchMediaByIndex = async (
  libraryId,
  parentMediaId,
  mediaIndex
) => {
  const url = `${MEDIA}?libraryId=${libraryId}&parentMediaId=${parentMediaId}&mediaIndex=${mediaIndex}`
  const media = await fetchJSON(url)
  return media?.data?.[0]
}

export const fetchMediaPage = async (mediaId, pageIndex) => {
  return fetchBase64(`${MEDIA}/${mediaId}/content?page=${pageIndex}`)
}

export const markAsRead = (mediaId) =>
  fetchJSON(`/api/media/${mediaId}/read`, { method: 'PATCH' })

export const markAsUnread = (mediaId) =>
  fetchJSON(`/api/media/${mediaId}/unread`, { method: 'PATCH' })
