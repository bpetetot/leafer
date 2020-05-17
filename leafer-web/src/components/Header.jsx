import React from 'react'

import { Container } from './Container'
import styles from './Header.module.css'

const Header = ({ title, children }) => {
  return (
    <header className={styles.header}>
      <Container>
        <div className={styles.content}>
          <div className={styles.left}>{title}</div>
          <div className={styles.right}>{children}</div>
        </div>
      </Container>
    </header>
  )
}

export default Header
