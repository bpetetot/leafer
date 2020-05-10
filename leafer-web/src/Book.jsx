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
  const { libraryId, bookId } = useParams()
  const [loading, setLoading] = useState(true)
  const [files, setFiles] = useState([])
  const { pathname } = useLocation()

  useEffect(() => {
    fetch(`/api/libraries/${libraryId}/book/${bookId}`)
      .then((response) => response.json())
      .then((files) => {
        setLoading(false)
        setFiles(files)
      })
  }, [libraryId, bookId])

  const loadImage = useCallback(
    (index) => {
      if (!files || files.length === 0) return Promise.resolve()
      const url = `/api/libraries/${libraryId}/book/${bookId}/file/${files[index]}`
      return fetch(url)
        .catch(() => (window.location.href = '/lost-in-space'))
        .then((response) => response.arrayBuffer())
        .then((buffer) => `data:image/*;base64, ${arrayBufferToBase64(buffer)}`)
    },
    [libraryId, bookId, files]
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
          pageCount={files.length}
          loadPage={loadImage}
          onPageChanged={onHandlePageChanged}
        />
      </PageContainer>
    </>
  )
}

export default Book
