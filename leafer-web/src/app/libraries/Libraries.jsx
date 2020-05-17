import React from 'react'
import { Link } from 'react-router-dom'

import { ReactComponent as RemoveIcon } from 'assets/icons/trash.svg'
import { ReactComponent as ScanIcon } from 'assets/icons/refresh-cw.svg'
import Header from 'components/Header'
import { PageContainer } from 'components/Container'
import { IconButton } from 'components/Button'
import { List, ListItem } from 'components/List'
import { useLibraries, removeLibrary, scanLibrary } from 'services/libraries'

function Libraries() {
  const { data = [] } = useLibraries()

  return (
    <>
      <Header title="Libraries">
        <Link to="/library/new">Add library</Link>
      </Header>
      <PageContainer>
        <List>
          {data.map(({ id, name, path }) => (
            <ListItem key={id}>
              <div>
                <Link to={`/library/${id}`}>{name}</Link>
                <p style={{ fontSize: '0.75rem', color: '#6b7280' }}>{path}</p>
              </div>
              <div style={{ display: 'flex' }}>
                <IconButton onClick={() => scanLibrary(id)}>
                  <ScanIcon />
                </IconButton>
                <IconButton onClick={() => removeLibrary(id)}>
                  <RemoveIcon />
                </IconButton>
              </div>
            </ListItem>
          ))}
        </List>
      </PageContainer>
    </>
  )
}

export default Libraries
