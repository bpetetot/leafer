import React, { useCallback } from 'react'
import { useParams, useLocation } from 'react-router-dom'

import Header from './components/Header'
import { PageContainer } from './components/Container'
import Reader from './Reader'
import { fetchBase64, history, fetchJSON } from './utils'
import { useQueryParam } from './useQueryParam'
import useSWR from 'swr'

function Media() {
  const { mediaId } = useParams()
  const pageIndex = parseInt(useQueryParam('page', 0))
  const { pathname } = useLocation()

  const { data } = useSWR(`/api/media/${mediaId}`, fetchJSON)

  const loadImage = useCallback(
    (index) => fetchBase64(`/api/media/${mediaId}/content?page=${index}`),
    [mediaId]
  )

  const onHandlePageChanged = useCallback(
    (pageIndex) => history.push(`${pathname}?page=${pageIndex}`),
    [pathname]
  )

  const pageCount = data?.data?.pageCount || 0
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

export default Media
