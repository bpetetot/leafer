import React from 'react'
import { useForm } from 'react-hook-form'
import { useNavigate } from 'react-router-dom'

import { PageContainer } from 'components/Container'
import { Button } from 'components/Button'
import Header from 'components/Header'
import Input from 'components/Input'
import { addLibrary } from 'services/libraries'

const NewLibrary = () => {
  const { register, handleSubmit, errors, formState } = useForm()
  const navigate = useNavigate()

  const onSubmit = async (data) => {
    await addLibrary(data)
    navigate('/')
  }

  return (
    <>
      <Header title="Add new library" />
      <PageContainer>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Input
            ref={register({ required: 'Required' })}
            type="text"
            name="name"
            label="Name"
            error={errors.name}
          />
          <Input
            ref={register({ required: 'Required' })}
            type="text"
            name="path"
            label="Directory path"
            error={errors.path}
          />
          <Button>Save</Button>
          {formState.isSubmitting && 'submitting'}
        </form>
      </PageContainer>
    </>
  )
}

export default NewLibrary
