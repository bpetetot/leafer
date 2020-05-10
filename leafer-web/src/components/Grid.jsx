import styled from '@emotion/styled'

export const Grid = styled.ul`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(128px, 1fr));
  grid-gap: 2.5rem;
  margin: 0 auto;
`

export const GridItem = styled.ul`
  width: 100%;
  height: 100%;
`
