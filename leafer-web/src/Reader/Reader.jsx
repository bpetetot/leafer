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
  pageIndex = 0,
  pageCount,
  loadPage,
  onPageChanged,
  displayMode,
}) {
  const [page, setPage] = useState(null)
  const [currentPageIndex, setCurrentPageIndex] = useState(pageIndex)
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

  useEffect(() => setCurrentPageIndex(pageIndex), [pageIndex])

  useEffect(() => {
    // avoid loading if page in cache
    if (!cache.current[currentPageIndex]) {
      setPage(undefined)
    }
    // load current page
    handleLoadPage(currentPageIndex)
      .then(setPage)
      .then(() => {
        // load next page in cache
        handleLoadPage(currentPageIndex + 1)
      })
  }, [handleLoadPage, currentPageIndex])

  const nextPage = () => {
    if (currentPageIndex === null) return
    const nextIndex = currentPageIndex + 1
    if (nextIndex >= pageCount) return
    setCurrentPageIndex(nextIndex)
    if (onPageChanged) onPageChanged(nextIndex)
  }

  const previousPage = () => {
    if (currentPageIndex === null) return
    const previousIndex = currentPageIndex - 1
    if (previousIndex < 0) return
    setCurrentPageIndex(previousIndex)
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
      {/* <div className="message">
        {currentPageIndex + 1} / {pageCount}
      </div> */}
    </div>
  )
}

export default Reader
