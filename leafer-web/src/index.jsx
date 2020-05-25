import React from 'react'
import ReactDOM from 'react-dom'
import { SWRConfig } from 'swr'
import { BrowserRouter, Routes, Route } from 'react-router-dom'

import { fetchJSON as fetcher } from './services/utils'
import {FullscreenProvider} from './hooks/useFullscreen'

import { Libraries, LibraryMedias, NewLibrary } from './app/libraries'
import { MediaDetail, MediaReader } from './app/media'
import Navbar from './app/common/Navbar'
import Settings from './app/Settings'
import NotFound from './app/NotFound'


import './styles'

ReactDOM.render(
  <React.StrictMode>
    <SWRConfig value={{ fetcher }}>
      <FullscreenProvider>
        <BrowserRouter>
          <Navbar />
          <Routes>
            <Route path="/" element={<Libraries />} />
            <Route path="/library/new" element={<NewLibrary />} />
            <Route path="/library/:libraryId" element={<LibraryMedias />} />
            <Route path="/library/:libraryId/:serieId" element={<MediaDetail />} />
            <Route path="/library/:libraryId/:serieId/:mediaId" element={<MediaReader />} />
            <Route path="/settings" element={<Settings />} />
            <Route path="/*" element={<NotFound />} />
          </Routes>
        </BrowserRouter>
      </FullscreenProvider>
    </SWRConfig>
  </React.StrictMode>,
  document.getElementById('root')
)
