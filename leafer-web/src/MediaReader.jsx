import React, { useCallback } from 'react'
import { useParams, useLocation, useNavigate } from 'react-router-dom'

import Header from './components/Header'
import { PageContainer } from './components/Container'
import Reader from './Reader'
import { fetchJSON, fetchBase64 } from './utils'
import { useQueryParam } from './useQueryParam'
import useSWR from 'swr'

function MediaReader() {
  const { libraryId, collectionId, mediaId } = useParams()
  const navigate = useNavigate()
  const pageIndex = parseInt(useQueryParam('page', 0))
  const { pathname } = useLocation()

  const { data } = useSWR(`/api/media/${mediaId}`, fetchJSON)

  const loadImage = useCallback(
    (index) => fetchBase64(`/api/media/${mediaId}/content?page=${index}`),
    [mediaId]
  )

  const handlePageChanged = useCallback(
    (pageIndex) => navigate(`${pathname}?page=${pageIndex}`),
    [pathname, navigate]
  )

  const mediaIndex = data?.data?.mediaIndex || 0
  const handleChangeMedia = useCallback(
    async (step) => {
      const index = mediaIndex + step
      const url = `/api/media?libraryId=${libraryId}&parentMediaId=${collectionId}&mediaIndex=${index}`
      const response = await fetchJSON(url)
      const next = response?.data?.[0]
      if (next) {
        navigate(`/library/${libraryId}/${collectionId}/${next.id}`)
      }
    },
    [collectionId, libraryId, mediaIndex, navigate]
  )

  const pageCount = data?.data?.pageCount || 0
  return (
    <>
      <Header title={data?.data?.fileName}>
        <button onClick={() => handleChangeMedia(-1)}>Previous chapter</button>
        <button onClick={() => handleChangeMedia(1)}>Next chapter</button>
      </Header>
      <PageContainer>
        <Reader
          id={mediaId}
          pageIndex={pageIndex}
          pageCount={pageCount}
          loadPage={loadImage}
          onPageChanged={handlePageChanged}
        />
      </PageContainer>
    </>
  )
}

export default MediaReader
