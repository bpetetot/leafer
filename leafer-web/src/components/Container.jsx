import React from 'react'
import cn from 'classnames'
import styles from './Container.module.css'

export const Container = ({ className, ...props}) => {
  return <div {...props} className={cn(styles.container, className)} />
}

export const PageContainer = (props) => {
  return <Container {...props} className={styles.page} />
}
