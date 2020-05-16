/** @jsx jsx */
import { css, jsx } from '@emotion/core'

import {Container} from '../components/Container'

const Header = ({title, children}) => {
  return (
    <header css={styles.header}>
      <Container>
        <div css={styles.content}>
          <div css={styles.left}>
            {title}
          </div>
          <div css={styles.right}>
            {children}
          </div>
        </div>
      </Container>
    </header>
  )
}

const styles = {
  header: css`
    position: sticky;
    top: 0;
    z-index: 10;
    background-color: #ffffff;
    box-shadow: 0 1px 3px 0 rgba(0,0,0,.1), 0 1px 2px 0 rgba(0,0,0,.06);
  `,
  content: css`
    height: 4rem;    
    display: flex;
    justify-content: space-between;
    align-items: center;
  `,
  left: css`
    display: flex;
    align-items: center;
  `,
  right: css`
    display: flex;
    align-items: center;
    margin-left: 1rem;
  `,
}

export default Header
