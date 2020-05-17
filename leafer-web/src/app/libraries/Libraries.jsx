import React from 'react'
import { Link } from 'react-router-dom'

import { ReactComponent as RemoveIcon } from 'assets/icons/trash.svg'
import Header from 'components/Header'
import { PageContainer } from 'components/Container'
import { IconButton } from 'components/Button'
import { List, ListItem } from 'components/List'
import { useLibraries } from 'services/libraries'

function Libraries() {
  const { data = [] } = useLibraries()

  return (
    <>
      <Header title="Libraries">
        <Link to="/library/new">Add library</Link>
      </Header>
      <PageContainer>
        <List>
          {data.map(({ id, path }) => (
            <ListItem key={id}>
              <Link to={`/library/${id}`}>{path}</Link>
              <IconButton>
                <RemoveIcon />
              </IconButton>
            </ListItem>
          ))}
        </List>
      </PageContainer>
    </>
  )
}

export default Libraries
