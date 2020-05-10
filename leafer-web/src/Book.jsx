import React, { useEffect, useState, useCallback } from 'react'
import { useParams, useLocation } from 'react-router-dom'

import Header from './components/Header'
import {PageContainer} from './components/Container'
import Reader from './Reader'

function arrayBufferToBase64(buffer) {
  var binary = ''
  var bytes = [].slice.call(new Uint8Array(buffer))
  bytes.forEach((b) => (binary += String.fromCharCode(b)))
  return window.btoa(binary)
}

function useURLPageIndex() {
  const { search } = useLocation()
  const params = new URLSearchParams(search)
  return parseInt(params.get('page')) || 0
}

function Book() {
  const pageIndex = useURLPageIndex()
  const { bookId } = useParams()
  const [loading, setLoading] = useState(true)
  const [pageCount, setPageCount] = useState(0)
  const { pathname } = useLocation()

  useEffect(() => {
    fetch(`/api/media/${bookId}`)
      .then((response) => response.json())
      .then((json) => {
        setLoading(false)
        setPageCount(json.data.pageCount)
      })
  }, [bookId])

  const loadImage = useCallback(
    (index) => {
      const url = `/api/media/${bookId}/content?page=${index}`
      return fetch(url)
        .catch(() => (window.location.href = '/lost-in-space'))
        .then((response) => response.arrayBuffer())
        .then((buffer) => `data:image/*;base64, ${arrayBufferToBase64(buffer)}`)
    },
    [bookId]
  )

  const onHandlePageChanged = useCallback(
    (pageIndex) => {
      window.history.pushState(null, '', `${pathname}?page=${pageIndex}`)
    },
    [pathname]
  )

  if (loading) return <p>Loading book...</p>

  return (
    <>
      <Header title="Lecture" />
      <PageContainer>
        <Reader
          pageIndex={pageIndex}
          pageCount={pageCount}
          loadPage={loadImage}
          onPageChanged={onHandlePageChanged}
        />
      </PageContainer>
    </>
  )
}

export default Book
