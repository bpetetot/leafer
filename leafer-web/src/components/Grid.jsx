import React from 'react'
import styles from './Grid.module.css'

export const Grid = (props) => {
  return <ul {...props} className={styles.grid} />
}

export const GridItem = (props) => {
  return <li {...props} className={styles.item} />
}
