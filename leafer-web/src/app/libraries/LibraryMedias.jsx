import React from 'react'
import { useParams, Link } from 'react-router-dom'

import { PageContainer } from 'components/Container'
import { Grid, GridItem } from 'components/Grid'
import { useMediasLibrary } from 'services/media'
import MediaHeader from '../common/MediaHeader'
import Text from 'components/Text'

function LibraryMedias() {
  let { libraryId } = useParams()
  const { data: medias } = useMediasLibrary(libraryId)

  if (!medias) return <p>Loading...</p>
  return (
    <>
      <MediaHeader />
      <PageContainer>
        <Grid>
          {medias?.data?.map((media) => (
            <GridItem key={media.id}>
              <Link to={`/library/${libraryId}/${media.id}`}>
                <div
                  style={{
                    height: '224px',
                    background: `no-repeat url(${media.coverImageUrl || media.coverImageLocal})`,
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
                <Link to={`/library/${libraryId}/${media.id}`}>
                  {media.title}
                </Link>
                {media.type === 'COLLECTION' && (
                  <Text size="sm">
                    {media.mediaCount} media
                  </Text>
                )}
              </div>
            </GridItem>
          ))}
        </Grid>
      </PageContainer>
    </>
  )
}

export default LibraryMedias
