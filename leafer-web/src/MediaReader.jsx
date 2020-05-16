import React, { useCallback } from 'react'
import { useParams, useLocation, useNavigate } from 'react-router-dom'

import Header from './layout/Header'
import { PageContainer } from './components/Container'
import Reader from './Reader'
import { useQueryParam } from './useQueryParam'
import { fetchMediaByIndex, fetchMediaPage, useMedia } from './services/media'

import { ReactComponent as NextIcon } from './assets/icons/chevron-right.svg'
import { ReactComponent as NextMedia } from './assets/icons/chevrons-right.svg'
import { ReactComponent as PreviousIcon } from './assets/icons/chevron-left.svg'
import { ReactComponent as PreviousMedia } from './assets/icons/chevrons-left.svg'

function MediaReader() {
  const { libraryId, collectionId, mediaId } = useParams()
  const { pathname } = useLocation()
  const navigate = useNavigate()

  const { data: media } = useMedia(mediaId)
  const mediaIndex = media?.mediaIndex || 0
  const pageCount = media?.pageCount || 0
  const pageIndex = parseInt(useQueryParam('page', 0))

  const loadImage = useCallback((index) => fetchMediaPage(mediaId, index), [
    mediaId,
  ])

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
      const next = await fetchMediaByIndex(libraryId, collectionId, index)
      if (!next) return
      navigate(`/library/${libraryId}/${collectionId}/${next.id}`)
    },
    [collectionId, libraryId, navigate]
  )

  return (
    <>
      <Header>
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
