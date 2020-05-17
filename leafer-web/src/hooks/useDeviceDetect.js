import { useEffect, useState, useMemo } from 'react'
import throttle from 'lodash/throttle'

const DELAY = 200

const MOBILE = 640
const TABLET = 768
const DESKTOP = 1024

const isClient = typeof window === 'object'

const getSize = () => {
  return isClient ? window.innerWidth : undefined
}

export const useDeviceDetect = () => {
  const [windowSize, setWindowSize] = useState(getSize)

  useEffect(() => {
    if (!isClient) return false

    const handleResize = () => setWindowSize(getSize())
    const throttleHandleResize = throttle(handleResize, DELAY)

    window.addEventListener('resize', throttleHandleResize)
    return () => {
      throttleHandleResize.cancel()
      window.removeEventListener('resize', throttleHandleResize)
    }
  }, [])

  return useMemo(() => ({
    isMobile: windowSize < MOBILE,
    isTablet: windowSize >= MOBILE && windowSize < TABLET,
    isDesktop: windowSize >= DESKTOP,
  }), [windowSize])
}