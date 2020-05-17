import React from 'react'
import { Link } from 'react-router-dom'

import Header from 'components/Header'
import { PageContainer } from 'components/Container'
import { List, ListItem } from 'components/List'
import { useLibraries, removeLibrary } from 'services/libraries'

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
              <button onClick={() => removeLibrary(id)}>Remove</button>
            </ListItem>
          ))}
        </List>
      </PageContainer>
    </>
  )
}

export default Libraries
