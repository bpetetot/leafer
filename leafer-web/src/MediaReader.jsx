import React, { useCallback } from 'react'
import { useParams, useLocation, useNavigate } from 'react-router-dom'

import Header from './components/Header'
import { PageContainer } from './components/Container'
import Reader from './Reader'
import { fetchJSON, fetchBase64 } from './utils'
import { useQueryParam } from './useQueryParam'
import useSWR from 'swr'

import { ReactComponent as NextIcon } from './assets/icons/chevron-right.svg'
import { ReactComponent as NextMedia } from './assets/icons/chevrons-right.svg'
import { ReactComponent as PreviousIcon } from './assets/icons/chevron-left.svg'
import { ReactComponent as PreviousMedia } from './assets/icons/chevrons-left.svg'

function MediaReader() {
  const { libraryId, collectionId, mediaId } = useParams()
  const { pathname } = useLocation()
  const navigate = useNavigate()

  const { data } = useSWR(`/api/media/${mediaId}`, fetchJSON)
  const mediaIndex = data?.data?.mediaIndex || 0
  const pageCount = data?.data?.pageCount || 0
  const pageIndex = parseInt(useQueryParam('page', 0))

  const loadImage = useCallback(
    (index) => fetchBase64(`/api/media/${mediaId}/content?page=${index}`),
    [mediaId]
  )

  const handlePageChanged = useCallback(
    (pageIndex) => {
      if (pageIndex === null) return
      if (pageIndex < 0 || pageIndex >= pageCount) return
      navigate(`${pathname}?page=${pageIndex}`)
    },
    [pathname, pageCount, navigate]
  )

  const handleChangeMedia = useCallback(
    async (index) => {
      const url = `/api/media?libraryId=${libraryId}&parentMediaId=${collectionId}&mediaIndex=${index}`
      const response = await fetchJSON(url)
      const next = response?.data?.[0]
      if (next) {
        navigate(`/library/${libraryId}/${collectionId}/${next.id}`)
      }
    },
    [collectionId, libraryId, navigate]
  )

  return (
    <>
      <Header title={data?.data?.fileName}>
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <button onClick={() => handleChangeMedia(mediaIndex - 1)}>
            <PreviousMedia />
          </button>
          <button onClick={() => handlePageChanged(pageIndex - 1)}>
            <PreviousIcon />
          </button>
          <div>{`${pageIndex + 1} / ${pageCount}`}</div>
          <button onClick={() => handlePageChanged(pageIndex + 1)}>
            <NextIcon />
          </button>
          <button onClick={() => handleChangeMedia(mediaIndex + 1)}>
            <NextMedia />
          </button>
        </div>
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
