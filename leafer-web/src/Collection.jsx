import React from 'react'
import { useParams, Link } from 'react-router-dom'
import useSWR from 'swr'
import Header from './components/Header'
import { PageContainer } from './components/Container'
import { List, ListItem } from './components/List'

const fetcher = (...args) =>
  fetch(...args)
    .then((res) => res.json())
    .catch(() => (window.location.href = '/lost-in-space'))

function Collection() {
  let { libraryId, collectionId } = useParams()
  const { data } = useSWR(
    `/api/libraries/${libraryId}/collections/${collectionId}`,
    fetcher
  )

  if (!data) return <p>Loading...</p>
  const { collection } = data
  return (
    <>
      <Header title={collection.title}>
        <Link to={`/library/${libraryId}`}>Back</Link>
      </Header>
      <PageContainer>
        <div style={{ display: 'flex' }}>
          <div style={{ maxWidth: '20%' }}>
            <img
              src={collection.coverImage.large}
              alt={collection.title}
              style={{
                maxWidth: '100%',
                maxHeight: '100%',
              }}
            />
          </div>
          <div
            style={{
              marginLeft: '3rem',
              paddingBottom: '1rem',
              minWidth: '10%',
              width: '100%',
            }}
          >
            <h1
              style={{
                color: '#161e2e',
                fontSize: '1.875rem',
                lineHeight: '1.25',
                fontWeight: '700',
                marginBottom: '2rem',
              }}
            >
              {collection.title}
            </h1>
            <p dangerouslySetInnerHTML={{ __html: collection.description }} />
          </div>
        </div>
        <h2 style={{ marginTop: '2rem' }}>
          {collection.books.length} book(s)
        </h2>
        <List>
          {collection.books.map((book) => (
            <ListItem key={book.id}>
              <Link to={`/library/${libraryId}/book/${book.id}`}>
                {book.tomeName} #{String(book.tomeNumber).padStart(3, '0')}
              </Link>
            </ListItem>
          ))}
        </List>
      </PageContainer>
    </>
  )
}

export default Collection
