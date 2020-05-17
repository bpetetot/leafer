import React from 'react'
import ReactDOM from 'react-dom'
import { SWRConfig } from 'swr'
import { BrowserRouter, Routes, Route } from 'react-router-dom'

import { fetchJSON as fetcher } from './services/utils'
import Navbar from './app/Navbar'
import Home from './app/Home'
import MediaLibrary from './app/MediaLibrary'
import MediaDetail from './app/MediaDetail'
import MediaReader from './app/MediaReader'
import Settings from './app/Settings'
import NotFound from './app/NotFound'

import './styles'

ReactDOM.render(
  <React.StrictMode>
    <SWRConfig value={{ fetcher }}>
      <BrowserRouter>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/library/:libraryId" element={<MediaLibrary />} />
          <Route path="/library/:libraryId/:collectionId" element={<MediaDetail />} />
          <Route path="/library/:libraryId/:collectionId/:mediaId" element={<MediaReader />} />
          <Route path="/settings" element={<Settings />} />
          <Route path="/*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </SWRConfig>
  </React.StrictMode>,
  document.getElementById('root')
)
