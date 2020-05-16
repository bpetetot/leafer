import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter as RouterProvider } from 'react-router-dom'
import { SWRConfig } from 'swr'

import App from './App'
import { fetchJSON as fetcher } from './services/utils'
import { ModalsProvider } from './components/Modal/ModalsProvider'

import './styles'

ReactDOM.render(
  <React.StrictMode>
    <SWRConfig value={{ fetcher }}>
      <RouterProvider>
        <ModalsProvider>
          <App />
        </ModalsProvider>
      </RouterProvider>
    </SWRConfig>
  </React.StrictMode>,
  document.getElementById('root')
)
