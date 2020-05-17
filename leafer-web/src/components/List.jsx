import React from 'react'
import styles from './List.module.css'

export const List = (props) => {
  return <ul {...props} className={styles.list} />
}

export const ListItem = (props) => {
  return <li {...props} className={styles.item} />
}

