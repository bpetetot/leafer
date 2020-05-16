import React from 'react'
import { useParams, Link } from 'react-router-dom'

import { useLibrary } from '../services/libraries'
import { useMedia } from '../services/media'

const Breadcrumb = () => {
  const { libraryId, collectionId, mediaId } = useParams()

  const { data: library } = useLibrary(libraryId)
  const { data: collection } = useMedia(collectionId)
  const { data: media } = useMedia(mediaId)

  return (
    <ul style={{ display: 'flex' }}>
      <li>
        <Link to="/">Libraries</Link>
      </li>
      {library && (
        <li>
          &nbsp;&gt;&nbsp;
          {collection ? (
            <Link to={`/library/${libraryId}`}>{library.name}</Link>
          ) : (
            library.name
          )}
        </li>
      )}
      {collection && (
        <li>
          &nbsp;&gt;&nbsp;
          {media ? (
            <Link to={`/library/${libraryId}/${collectionId}`}>
              {collection.estimatedName}
            </Link>
          ) : (
            collection.estimatedName
          )}
        </li>
      )}
      {media && <li>&nbsp;&gt;&nbsp;{media.fileName}</li>}
    </ul>
  )
}

export default Breadcrumb
