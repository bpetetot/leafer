import React from 'react'
import { useParams, Link } from 'react-router-dom'

import { PageContainer } from 'components/Container'
import { List, ListItem } from 'components/List'
import { useMedia, useMediasCollection } from 'services/media'
import MediaHeader from '../common/MediaHeader'
import Text from 'components/Text'

function MediaDetail() {
  let { libraryId, collectionId } = useParams()
  const { data: media } = useMedia(collectionId)
  const { data: medias } = useMediasCollection(libraryId, media)

  if (!media) return <p>Loading...</p>

  return (
    <>
      <MediaHeader />
      <PageContainer>
        <div style={{ display: 'flex', marginBottom: '1.5rem' }}>
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
            <Text dangerouslySetInnerHTML={{ __html: media.description }} />
          </div>
        </div>
        <List>
          {(medias?.data || [media])?.map((item) => (
            <ListItem key={item.id}>
              <Link to={`/library/${libraryId}/${collectionId}/${item.id}`}>
                #{String(item.volume || 0).padStart(3, '0')} {media.title || media.titleNative}
              </Link>
            </ListItem>
          ))}
        </List>
      </PageContainer>
    </>
  )
}

export default MediaDetail
