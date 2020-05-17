import React from 'react'
import { Link } from 'react-router-dom'

import { ReactComponent as Logo } from 'assets/logo.svg'
import { ReactComponent as IconSettings } from 'assets/icons/settings.svg'
import { useFullscreen } from 'hooks/useFullscreen'
import { Container } from 'components/Container'

import styles from './Navbar.module.css'

const Navbar = () => {
  const { fullscreen } = useFullscreen()

  if (fullscreen) return null

  return (
    <nav className={styles.navbar}>
      <Container>
        <div className={styles.content}>
          <div className={styles.left}>
            <Link to="/" className={styles.logo}>
              <Logo />
            </Link>
          </div>
          <div className={styles.right}>
            <Link to="/settings" className={styles.button}>
              <IconSettings />
            </Link>
          </div>
        </div>
      </Container>
    </nav>
  )
}

export default Navbar
