import React from 'react'
import { useParams, Link } from 'react-router-dom'
import Header from './components/Header'
import { PageContainer } from './components/Container'
import { List, ListItem } from './components/List'
import { useMedia } from './services/media'

function MediaDetail() {
  let { libraryId, collectionId } = useParams()
  const { data: media } = useMedia(collectionId)

  if (!media) return <p>Loading...</p>

  const medias = media.type === 'COLLECTION' ? media.medias : [media]

  return (
    <>
      <Header title={media.title}>
        <Link to={`/library/${libraryId}`}>Back</Link>
      </Header>
      <PageContainer>
        <div style={{ display: 'flex' }}>
          <div style={{ maxWidth: '20%' }}>
            <img
              src={media.coverImage}
              alt={media.title || media.titleNative}
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
              {media.title}
            </h1>
            <p dangerouslySetInnerHTML={{ __html: media.description }} />
          </div>
        </div>
        {media.mediaCount && (
          <h2 style={{ marginTop: '2rem' }}>{media.mediaCount} media</h2>
        )}
        <List>
          {medias.map((media) => (
            <ListItem key={media.id}>
              <Link to={`/library/${libraryId}/${collectionId}/${media.id}`}>
                #{String(media.volume || 0).padStart(3, '0')} {media.fileName}
              </Link>
            </ListItem>
          ))}
        </List>
      </PageContainer>
    </>
  )
}

export default MediaDetail
