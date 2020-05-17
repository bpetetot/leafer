import React, { useContext, useState, useCallback, useMemo } from 'react'

const FullscreenContext = React.createContext()

export const useFullscreen = () => useContext(FullscreenContext)

export const FullscreenProvider = ({ children }) => {
  const [fullscreen, setFullscreen] = useState(false)

  const toggleFullscreen = useCallback(() => {
    setFullscreen(!fullscreen)
  }, [fullscreen])

  const value = useMemo(
    () => ({
      fullscreen,
      toggleFullscreen,
    }),
    [fullscreen, toggleFullscreen]
  )

  return (
    <FullscreenContext.Provider value={value}>
      {children}
    </FullscreenContext.Provider>
  )
}
