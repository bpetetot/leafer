import React from 'react'
import { useParams, Link } from 'react-router-dom'

import { useLibrary } from 'services/libraries'
import { useMedia } from 'services/media'
import Header from 'components/Header'

const MediaHeader = ({ children }) => (
  <Header title={<Breadcrumb />}>{children}</Header>
)

const Breadcrumb = () => {
  const { libraryId, serieId, mediaId } = useParams()

  const { data: library } = useLibrary(libraryId)
  const { data: serie } = useMedia(serieId)
  const { data: media } = useMedia(mediaId)

  return (
    <ul style={{ display: 'flex' }}>
      {library && (
        <li>
          {serie ? (
            <Link to={`/library/${libraryId}`}>{library.name}</Link>
          ) : (
            library.name
          )}
        </li>
      )}
      {serie && (
        <li>
          &nbsp;&gt;&nbsp;
          {media ? (
            <Link to={`/library/${libraryId}/${serieId}`}>
              {serie.title}
            </Link>
          ) : (
            serie.title
          )}
        </li>
      )}
      {media && <li>&nbsp;&gt;&nbsp;#{String(media.volume || 0).padStart(3, '0')}</li>}
    </ul>
  )
}

export default MediaHeader
