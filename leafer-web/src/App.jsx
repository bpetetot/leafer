import React, { memo } from 'react'
import { Routes, Route } from 'react-router-dom'

import Layout from './layout'
import MediaReader from './MediaReader'
import MediaDetail from './MediaDetail'
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
        <Route path="/library/:libraryId/:collectionId" element={<MediaDetail />} />
        <Route path="/library/:libraryId/:collectionId/:mediaId" element={<MediaReader />} />
        <Route path="/settings" element={<Settings />} />
        <Route path="/*" element={<NotFound />} />
      </Routes>
    </Layout>
  )
}

export default memo(App)
