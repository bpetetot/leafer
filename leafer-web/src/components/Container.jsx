import styled from '@emotion/styled'

export const Container = styled.div`
  margin-left: auto;
  margin-right: auto;
  max-width: 80rem;

  @media (min-width: 1024px) {
    padding-left: 2rem;
    padding-right: 2rem;
  }

  @media (min-width: 640px) {
    padding-left: 1.5rem;
    padding-right: 1.5rem;
  }
`

export const PageContainer = styled(Container)`
  padding-top: 2.5rem;
  padding-bottom: 2.5rem;
`