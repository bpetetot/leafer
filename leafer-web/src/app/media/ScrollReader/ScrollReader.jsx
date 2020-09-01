import React, { useState, useEffect, useCallback } from 'react'
import useInfiniteLoader from 'react-use-infinite-loader';

import './ScrollReader.css'

function ScrollReader({ id, pageCount, loadPage }) {
  const [pages, setPages] = useState([])
  const [canLoadMore, setCanLoadMore] = useState(true)

  const loadMore = useCallback(
    (curIndex) => {
      loadPage(curIndex).then((page) => {
        setCanLoadMore(curIndex + 1 < pageCount)
        setPages((pages) => [...pages, page])
      })
    },
    [loadPage, pageCount]
  )
  
  const { loaderRef, resetPage } = useInfiniteLoader({ loadMore, canLoadMore })

  useEffect(() => {
    window.scrollTo(0, 0)
    resetPage()
    setPages([])
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id])

  return (
    <div className="reader">
      {pages.length === 0 && <p>loading...</p>}
      {pages.map((p, i) => (
        <div key={i} className="page">
          <img src={p} alt={`Page ${i}`} />
        </div>
      ))}
      <div ref={loaderRef} className="loaderRef" />
    </div>
  )
}

export default ScrollReader
