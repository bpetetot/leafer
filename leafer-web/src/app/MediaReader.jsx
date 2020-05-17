import React, { useCallback } from 'react'
import { useParams, useLocation, useNavigate } from 'react-router-dom'

import { ReactComponent as NextIcon } from '../assets/icons/chevron-right.svg'
import { ReactComponent as NextMedia } from '../assets/icons/chevrons-right.svg'
import { ReactComponent as PreviousIcon } from '../assets/icons/chevron-left.svg'
import { ReactComponent as PreviousMedia } from '../assets/icons/chevrons-left.svg'
import { useQueryParam } from '../hooks/useQueryParam'
import { fetchMediaByIndex, fetchMediaPage, useMedia } from '../services/media'
import { PageContainer } from '../components/Container'

import MediaHeader from './MediaHeader'
import Reader from './Reader'
import { useKey } from '../hooks/useKey'

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
      navigate(`/library/${libraryId}/${collectionId}/${next.id}?page=0`)
    },
    [collectionId, libraryId, navigate]
  )

  const onNextPage = () => {
    if (pageIndex >= pageCount - 1) {
      handleChangeMedia(mediaIndex + 1)
    } else {
      handlePageChanged(pageIndex + 1)
    }
  }

  const onPreviousPage = () => {
    if (pageIndex <= 0) {
      handleChangeMedia(mediaIndex - 1)
    } else {
      handlePageChanged(pageIndex - 1)
    }
  }
  
  useKey('ArrowRight', onNextPage)
  useKey('ArrowLeft', onPreviousPage)

  return (
    <>
      <MediaHeader>
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <button onClick={() => handleChangeMedia(mediaIndex - 1)}>
            <PreviousMedia />
          </button>
          <button onClick={onPreviousPage}>
            <PreviousIcon />
          </button>
          <div>{`${pageIndex + 1} / ${pageCount}`}</div>
          <button onClick={onNextPage}>
            <NextIcon />
          </button>
          <button onClick={() => handleChangeMedia(mediaIndex + 1)}>
            <NextMedia />
          </button>
        </div>
      </MediaHeader>
      <PageContainer>
        <Reader
          id={mediaId}
          pageIndex={pageIndex}
          pageCount={pageCount}
          loadPage={loadImage}
          onNextPage={onNextPage}
          onPreviousPage={onPreviousPage}
        />
      </PageContainer>
    </>
  )
}

export default MediaReader
