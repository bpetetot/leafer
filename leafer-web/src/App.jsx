import React, { memo } from 'react'
import { Routes, Route } from 'react-router-dom'

import Layout from './layout'
import Book from './Book'
import Collection from './Collection'
import Library from './Library'
import Home from './Home'
import NotFound from './NotFound'
import Settings from './Settings'

function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/library/:libraryId" element={<Library />} />
        <Route path="/library/:libraryId/collection/:collectionId" element={<Collection />} />
        <Route path="/library/:libraryId/book/:bookId" element={<Book />} />
        <Route path="/settings" element={<Settings />} />
        <Route path="/*" element={<NotFound />} />
      </Routes>
    </Layout>
  )
}

export default memo(App)
