import { useEffect, useCallback } from 'react'

export function useKey(targetKey, onKeyDown) {
  const handler = useCallback(
    ({ key }) => {
      if (key === targetKey) onKeyDown()
    },
    [targetKey, onKeyDown]
  )

  useEffect(() => {
    window.addEventListener('keydown', handler)
    return () => {
      window.removeEventListener('keydown', handler)
    }
  }, [handler])
}
