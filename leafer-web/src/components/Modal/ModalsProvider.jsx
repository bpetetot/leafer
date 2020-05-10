import React, {
  useState,
  useContext,
  useCallback,
  useMemo,
  useEffect,
} from 'react'

const ModalsContext = React.createContext()

export const useModal = (name) => {
  const { registerModal, toggleModal, isOpenModal } = useContext(ModalsContext)

  useEffect(() => {
    registerModal(name)
  }, [registerModal, name])

  const isOpen = isOpenModal(name)
  const toggle = useCallback(() => toggleModal(name), [toggleModal, name])

  return useMemo(() => ({ isOpen, toggle }), [isOpen, toggle])
}

export const ModalsProvider = ({ children }) => {
  const [modals, setModal] = useState({})

  const registerModal = useCallback(
    (name) => {
      if (modals.hasOwnProperty(name)) return
      setModal({ ...modals, [name]: false })
    },
    [modals]
  )

  const toggleModal = useCallback(
    (name) => setModal({ ...modals, [name]: !modals[name] }),
    [modals]
  )

  const isOpenModal = useCallback((name) => modals[name], [modals])

  const value = useMemo(
    () => ({
      registerModal,
      toggleModal,
      isOpenModal,
    }),
    [registerModal, toggleModal, isOpenModal]
  )

  return (
    <ModalsContext.Provider value={value}>{children}</ModalsContext.Provider>
  )
}
