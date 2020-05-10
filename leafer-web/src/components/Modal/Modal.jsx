import React, { useState, useEffect } from 'react'
import cn from 'classnames'
import { useModal } from './ModalsProvider'

const OPENED_TRANSITIONS = {
  background: 'ease-out duration-200 opacity-100',
  modal: 'ease-out duration-200 opacity-100 translate-y-0 sm:scale-100',
}

const CLOSED_TRANSITIONS = {
  wrapper: 'hidden pointer-events-none',
  background: 'ease-out duration-300 opacity-0',
  modal: 'ease-out duration-300 opacity-0 translate-y-4 sm:translate-y-0',
}

const Modal = ({ name, title, children, onSubmit, submitLabel }) => {
  const { isOpen, toggle } = useModal(name)

  const [transitions, setTransitions] = useState(CLOSED_TRANSITIONS)

  useEffect(() => {
    if (isOpen) {
      setTransitions(OPENED_TRANSITIONS)
    } else {
      setTransitions(CLOSED_TRANSITIONS)
    }
  }, [isOpen])

  const wrapperClassName = cn(
    'fixed bottom-0 inset-x-0 px-4 pb-4 sm:inset-0 sm:flex sm:items-center sm:justify-center',
    transitions.wrapper
  )

  const backgroundClassName = cn(
    'fixed inset-0 transition-opacity',
    transitions.background
  )

  const modalClassName = cn(
    'bg-white rounded-lg overflow-hidden shadow-xl transform transition-all sm:max-w-lg sm:w-full',
    transitions.modal
  )

  return (
    <div className={wrapperClassName}>
      <div className={backgroundClassName}>
        <div className="absolute inset-0 bg-gray-500 opacity-75"></div>
      </div>
      <div className={modalClassName}>
        <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div className="mt-3 text-center sm:mt-0 sm:text-left">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                {title}
              </h3>
              <div className="mt-3">{children}</div>
            </div>
        </div>
        <div className="px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
          {onSubmit && (
            <span className="flex w-full rounded-md shadow-sm sm:ml-3 sm:w-auto">
              <button onClick={onSubmit} color="red">
                {submitLabel}
              </button>
            </span>
          )}
          <span className="mt-3 flex w-full rounded-md shadow-sm sm:mt-0 sm:w-auto">
            <button onClick={toggle} color="white">
              Cancel
            </button>
          </span>
        </div>
      </div>
    </div>
  )
}

export default Modal
