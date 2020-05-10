import React, { useCallback } from 'react'
import { Link } from 'react-router-dom'
import useSWR from 'swr'
import Header from './components/Header'
import { PageContainer } from './components/Container'
import { List, ListItem } from './components/List'

const fetcher = (...args) =>
  fetch(...args)
    .then((res) => res.json())
    .catch(() => (window.location.href = '/lost-in-space'))

function Home() {
  const { data, mutate } = useSWR('/api/libraries', fetcher)

  const removeLibrary = useCallback(
    async (id) => {
      await fetch(`/api/libraries/${id}`, {
        method: 'DELETE',
      })
      mutate({
        ...data,
        libraries: data.libraries.filter((lib) => lib.id !== id),
      })
    },
    [data, mutate]
  )

  return (
    <>
      <Header title="Libraries" />
      <PageContainer>
        <List>
          {data?.libraries.map(({ id, path }) => (
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

// const addLibrary = useCallback(async () => {
//   if (!path) return
//   const response = await fetch('/api/libraries', {
//     method: 'POST',
//     body: JSON.stringify({ path }),
//     headers: {
//       'Content-Type': 'application/json',
//     },
//   })
//   const { library } = await response.json()
//   mutate({ ...data, libraries: [...data.libraries, library] })
//   toggle()
// }, [path, data, mutate, toggle])

export default Home
