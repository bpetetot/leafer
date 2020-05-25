import useSWR from 'swr'
import { fetchBase64, fetchJSON } from './utils'

const MEDIA = '/api/media'

export const useMedia = (mediaId) => {
  return useSWR(mediaId ? [MEDIA, mediaId] : null, (url, mediaId) => {
    return fetchJSON(`${url}/${mediaId}`)
  })
}

export const useMediasLibrary = (libraryId) => {
  return useSWR(libraryId ? [MEDIA, libraryId, 0] : null, (url, libraryId, serieId) => {
    return fetchJSON(`${url}?libraryId=${libraryId}&serieId=${serieId}`)
  })
}

export const useMediasSerie = (libraryId, serie) => {
  const shouldFetch = serie && serie.type === 'SERIE'
  return useSWR(
    shouldFetch ? [MEDIA, libraryId, serie.id] : null,
    (url, libraryId, serieId) => {
      return fetchJSON(`${url}?libraryId=${libraryId}&serieId=${serieId}`)
    }
  )
}

export const fetchMediaByIndex = async (libraryId, serieId, mediaIndex) => {
  const url = `${MEDIA}?libraryId=${libraryId}&serieId=${serieId}&mediaIndex=${mediaIndex}`
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
