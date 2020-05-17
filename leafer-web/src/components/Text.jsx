import React, { memo } from 'react'
import cn from 'classnames'
import styles from './Text.module.css'

const Text = ({ size, ...props }) => {
  const classes = cn({
    [styles.small]: size === 'sm',
  })

  return <p className={classes} {...props} />
}

export default memo(Text)
