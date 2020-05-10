import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter as RouterProvider } from 'react-router-dom'

import App from './App'
import { ModalsProvider } from './components/Modal/ModalsProvider'

import './styles'

ReactDOM.render(
  <React.StrictMode>
    <RouterProvider>
      <ModalsProvider>
        <App />
      </ModalsProvider>
    </RouterProvider>
  </React.StrictMode>,
  document.getElementById('root')
)
