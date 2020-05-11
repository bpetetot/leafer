import React from 'react'
import { useParams, Link } from 'react-router-dom'
import useSWR from 'swr'
import Header from './components/Header'
import { PageContainer } from './components/Container'
import { List, ListItem } from './components/List'
import { fetchJSON } from './utils'

function MediaCollection() {
  let { libraryId, collectionId } = useParams()
  const { data } = useSWR(`/api/media/${collectionId}`, fetchJSON)

  if (!data) return <p>Loading...</p>
  const { data: collection } = data
  return (
    <>
      <Header title={collection.title || collection.titleNative}>
        <Link to={`/library/${libraryId}`}>Back</Link>
      </Header>
      <PageContainer>
        <div style={{ display: 'flex' }}>
          <div style={{ maxWidth: '20%' }}>
            <img
              src={collection.coverImage}
              alt={collection.title || collection.titleNative}
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
              {collection.title || collection.titleNative}
            </h1>
            <p dangerouslySetInnerHTML={{ __html: collection.description }} />
          </div>
        </div>
        <h2 style={{ marginTop: '2rem' }}>{collection.medias.length} book(s)</h2>
        <List>
          {collection.medias.map((media) => (
            <ListItem key={media.id}>
              <Link
                to={`/library/${libraryId}/${collectionId}/${media.id}`}
              >
                #{String(media.volume).padStart(3, '0')} {media.fileName}
              </Link>
            </ListItem>
          ))}
        </List>
      </PageContainer>
    </>
  )
}

export default MediaCollection
