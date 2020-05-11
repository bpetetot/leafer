import React, { memo } from 'react'
import { Routes, Route } from 'react-router-dom'

import Layout from './layout'
import Media from './Media'
import MediaCollection from './MediaCollection'
import MediaLibrary from './MediaLibrary'
import Home from './Home'
import NotFound from './NotFound'
import Settings from './Settings'

function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/library/:libraryId" element={<MediaLibrary />} />
        <Route path="/library/:libraryId/:collectionId" element={<MediaCollection />} />
        <Route path="/library/:libraryId/:collectionId/:mediaId" element={<Media />} />
        <Route path="/settings" element={<Settings />} />
        <Route path="/*" element={<NotFound />} />
      </Routes>
    </Layout>
  )
}

export default memo(App)
