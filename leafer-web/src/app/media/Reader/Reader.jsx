import React, { useState, useEffect, useRef, useCallback } from 'react'
import { useSwipeable } from 'react-swipeable'
import cn from 'classnames'

import './Reader.css'

const swipeConfig = {
  delta: 10,
  preventDefaultTouchmoveEvent: false,
  trackTouch: true,
  trackMouse: false,
  rotationAngle: 0,
}

function Reader({
  id,
  pageIndex = 0,
  pageCount,
  loadPage,
  onNextPage,
  onPreviousPage,
  displayMode,
}) {
  const [page, setPage] = useState(null)
  const cache = useRef([])

  const handleLoadPage = useCallback(
    async (index) => {
      if (index < 0 || index >= pageCount) return
      let curPage = cache.current[index]
      if (!curPage) {
        curPage = await loadPage(index)
        cache.current[index] = curPage
      }
      return curPage
    },
    [loadPage, pageCount]
  )

  useEffect(() => {
    cache.current = []
  }, [id])

  useEffect(() => {
    window.scrollTo(0, 0)
  }, [page])

  useEffect(() => {
    // load current page
    handleLoadPage(pageIndex)
      .then(setPage)
      .then(() => {
        // load next page in cache
        handleLoadPage(pageIndex + 1)
      })
  }, [handleLoadPage, pageIndex])

  const handlers = useSwipeable({
    onSwipedLeft: onNextPage,
    onSwipedRight: onPreviousPage,
    ...swipeConfig,
  })

  const readerClassNames = cn('reader', {
    'reader--fit-parent': displayMode === 'fit-parent',
  })

  return (
    <div className={readerClassNames} {...handlers}>
      <div className="page">
        {page ? <img src={page} alt="Page X" /> : <p>Loading page...</p>}
      </div>
    </div>
  )
}

export default Reader
