import React from 'react'
import { useParams, Link } from 'react-router-dom'
import Header from './components/Header'
import { PageContainer } from './components/Container'
import { Grid, GridItem } from './components/Grid'
import { useMediaLibrary } from './services/media'

function MediaLibrary() {
  let { libraryId } = useParams()
  const { data: medias } = useMediaLibrary(libraryId)

  if (!medias) return <p>Loading...</p>
  return (
    <>
      <Header title="Library">
        <Link to="/">Back</Link>
      </Header>
      <PageContainer>
        <Grid>
          {medias?.data?.map((media) => (
            <GridItem key={media.id}>
              <Link to={`/library/${libraryId}/${media.id}`}>
                <div
                  style={{
                    height: '224px',
                    background: `no-repeat url(${media.coverImage})`,
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
                  <p
                    style={{
                      fontSize: '0.8rem',
                      color: '#6b7280',
                    }}
                  >
                    {media.mediaCount} media
                  </p>
                )}
              </div>
            </GridItem>
          ))}
        </Grid>
      </PageContainer>
    </>
  )
}

export default MediaLibrary
