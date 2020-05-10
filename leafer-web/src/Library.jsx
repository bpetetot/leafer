import React from 'react'
import { useParams, Link } from 'react-router-dom'
import useSWR from 'swr'
import Header from './components/Header'
import { PageContainer } from './components/Container'
import { Grid, GridItem } from './components/Grid'

const fetcher = (...args) =>
  fetch(...args)
    .then((res) => res.json())
    .catch(() => (window.location.href = '/lost-in-space'))

function Library() {
  let { libraryId } = useParams()
  const { data } = useSWR(`/api/libraries/${libraryId}`, fetcher)

  return (
    <>
      <Header title={data?.library?.info?.path}>
        <Link to="/">Back</Link>
      </Header>
      <PageContainer>
        <Grid>
          {data?.library?.collections?.map((collection) => (
            <GridItem key={collection.id}>
              <Link to={`/library/${libraryId}/collection/${collection.id}`}>
                <div
                  style={{
                    height: '224px',
                    background: `no-repeat url(${collection.coverImage.large})`,
                    backgroundSize: 'cover',
                  }}
                />
              </Link>
              <div
                style={{
                  paddingTop: '0.75rem',
                  overflow: 'hidden',
                  whiteSpace: 'nowrap',
                  textOverflow: 'ellipsis',
                }}
              >
                <Link to={`/library/${libraryId}/collection/${collection.id}`}>
                  {collection.title}
                </Link>
                <p
                  style={{
                    fontSize: '0.8rem',
                    color: '#6b7280',
                  }}
                >
                  {collection.books.length} books
                </p>
              </div>
            </GridItem>
          ))}
        </Grid>
      </PageContainer>
    </>
  )
}

export default Library
