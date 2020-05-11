import React, { useState, useEffect, useRef, useCallback } from 'react'
import cn from 'classnames'
import { useSwipeable } from 'react-swipeable'

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
  onPageChanged,
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
    // avoid loading if page in cache
    if (!cache.current[pageIndex]) {
      setPage(undefined)
    }
    // load current page
    handleLoadPage(pageIndex)
      .then(setPage)
      .then(() => {
        // load next page in cache
        handleLoadPage(pageIndex + 1)
      })
  }, [handleLoadPage, pageIndex])

  const nextPage = () => {
    if (pageIndex === null) return
    const nextIndex = pageIndex + 1
    if (nextIndex >= pageCount) return
    if (onPageChanged) onPageChanged(nextIndex)
  }

  const previousPage = () => {
    if (pageIndex === null) return
    const previousIndex = pageIndex - 1
    if (previousIndex < 0) return
    if (onPageChanged) onPageChanged(previousIndex)
  }

  const handlers = useSwipeable({
    onSwipedLeft: () => nextPage(),
    onSwipedRight: () => previousPage(),
    ...swipeConfig,
  })

  const readerClassNames = cn('reader', {
    'reader--fit-parent': displayMode === 'fit-parent',
  })

  return (
    <div className={readerClassNames} {...handlers}>
      <div className="action">
        <button type="button" onClick={previousPage}>
          Previous
        </button>
      </div>
      <div className="page">
        {page ? <img src={page} alt="Page X" /> : <p>Loading page...</p>}
      </div>
      <div className="action">
        <button type="button" onClick={nextPage}>
          Next
        </button>
      </div>
    </div>
  )
}

export default Reader
