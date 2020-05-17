import React, { useContext, useState, useCallback, useMemo, useEffect } from 'react'
import { useDeviceDetect } from './useDeviceDetect'

const FullscreenContext = React.createContext()

export const useFullscreen = () => useContext(FullscreenContext)

export const FullscreenProvider = ({ children }) => {
  const [fullscreen, setFullscreen] = useState(false)

  const {isMobile, isTablet} = useDeviceDetect()

  const toggleFullscreen = useCallback(() => {
    setFullscreen(!fullscreen)
  }, [fullscreen])


  useEffect(() => {
    if (isMobile || isTablet) {
      setFullscreen(true)
    } else {
      setFullscreen(false)
    }
  }, [isMobile, isTablet])

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
