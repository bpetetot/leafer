import React from 'react'
import cn from 'classnames'
import styles from './Button.module.css'

export const Button = ({ className, ...props }) => {
  return <button {...props} className={cn(styles.button, className)} />
}

export const IconButton = (props) => {
  return <Button {...props} className={styles.iconButton} />
}
