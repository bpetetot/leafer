import React from 'react'
import cn from 'classnames'
import styles from './Input.module.css'

const Input = React.forwardRef(({ name, label, error }, ref) => {
  return (
    <div>
      <label htmlFor={name} className={styles.label}>
        {label}
      </label>
      <div className={styles.inputWrapper}>
        <input
          ref={ref}
          id={name}
          name={name}
          className={cn(styles.input, error && styles.inputError)}
        />
      </div>
      <div className={styles.error}>{error && error.message}</div>
    </div>
  )
})

export default Input
