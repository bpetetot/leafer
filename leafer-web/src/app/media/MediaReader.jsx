import React, { useCallback, useState } from 'react'
import { useParams, useLocation, useNavigate } from 'react-router-dom'

import { ReactComponent as NextIcon } from 'assets/icons/chevron-right.svg'
import { ReactComponent as NextMedia } from 'assets/icons/chevrons-right.svg'
import { ReactComponent as PreviousIcon } from 'assets/icons/chevron-left.svg'
import { ReactComponent as PreviousMedia } from 'assets/icons/chevrons-left.svg'
import { ReactComponent as BasicReaderIcon } from 'assets/icons/book-open.svg'
import { ReactComponent as ScrollReaderIcon } from 'assets/icons/chevrons-down.svg'

import { useKey } from 'hooks/useKey'
import { useFullscreen } from 'hooks/useFullscreen'
import { useQueryParam } from 'hooks/useQueryParam'
import {
  fetchMediaByIndex,
  fetchMediaPage,
  useMedia,
  markAsRead,
} from 'services/media'
import { PageContainer } from 'components/Container'
import { IconButton } from 'components/Button'

import MediaHeader from '../common/MediaHeader'
import Reader from './Reader'
import ScrollReader from './ScrollReader'

function MediaReader() {
  const { libraryId, serieId, mediaId } = useParams()
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
      if (pageIndex === pageCount - 1) markAsRead(media.id)
      navigate(`${pathname}?page=${pageIndex}`)
    },
    [pageCount, media, navigate, pathname]
  )

  const handleChangeMedia = useCallback(
    async (index) => {
      const next = await fetchMediaByIndex(libraryId, serieId, index)
      if (!next) return
      navigate(`/library/${libraryId}/${serieId}/${next.id}?page=0`)
    },
    [serieId, libraryId, navigate]
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

  const { fullscreen, toggleFullscreen } = useFullscreen()
  useKey('ArrowRight', onNextPage)
  useKey('ArrowLeft', onPreviousPage)
  useKey('f', toggleFullscreen)

  const [readerType, setReaderType] = useState('basic')

  return (
    <>
      {!fullscreen && (
        <MediaHeader>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <IconButton onClick={() => handleChangeMedia(mediaIndex - 1)}>
              <PreviousMedia />
            </IconButton>
            {readerType === 'basic' && (
              <>
                <IconButton onClick={onPreviousPage}>
                  <PreviousIcon />
                </IconButton>
                <div style={{ margin: '0 1rem' }}>
                  {`${pageIndex + 1} / ${pageCount}`}
                </div>
                <IconButton onClick={onNextPage}>
                  <NextIcon />
                </IconButton>
              </>
            )}
            <IconButton onClick={() => handleChangeMedia(mediaIndex + 1)}>
              <NextMedia />
            </IconButton>
            {readerType === 'basic' ? (
              <IconButton onClick={() => setReaderType('scroll')}>
                <ScrollReaderIcon />
              </IconButton>
            ) : (
              <IconButton onClick={() => setReaderType('basic')}>
                <BasicReaderIcon />
              </IconButton>
            )}
          </div>
        </MediaHeader>
      )}
      <PageContainer>
        {readerType === 'basic' ? (
          <Reader
            id={mediaId}
            pageIndex={pageIndex}
            pageCount={pageCount}
            loadPage={loadImage}
            onNextPage={onNextPage}
            onPreviousPage={onPreviousPage}
          />
        ) : (
          <ScrollReader
            id={mediaId}
            pageCount={pageCount}
            loadPage={loadImage}
          />
        )}
      </PageContainer>
    </>
  )
}

export default MediaReader
