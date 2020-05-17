import styled from '@emotion/styled'

export const Button = styled.button`
  padding-top: .5rem;
  padding-bottom: .5rem;

  padding-left: 1rem;
  padding-right: 1rem;

  font-size: 0.875rem;
  font-weight: 600;

  border-color: #d2d6dc;
  border-width: 1px;
  border-radius: .25rem;

  background-color: #fff;

  box-shadow: 0 1px 3px 0 rgba(0,0,0,.1),0 1px 2px 0 rgba(0,0,0,.06);

  line-height: inherit;

  &:hover {
    background-color: #f7fafc;
  }
`

export const IconButton = styled(Button)`
  display: flex;
  justify-content: center;
  align-items: center;

  height: 48px;
  width: 48px;
  
  padding: 0;
  margin: 0 0.25rem;

  border: 0;
  border-radius: 50%;

  box-shadow: unset;
`
