/** @jsx jsx */
import { css, jsx } from '@emotion/core'
import React from 'react'

const Input = React.forwardRef(({ name, label, error }, ref) => {
  return (
    <div>
      <label htmlFor={name} css={styles.label}>
        {label}
      </label>
      <div css={styles.inputWrapper}>
        <input
          ref={ref}
          id={name}
          name={name}
          css={[styles.input, error && styles.inputError]}
        />
      </div>
      <div css={styles.error}>{error && error.message}</div>
    </div>
  )
})

const styles = {
  label: css`
    display: block;
    color: #374151;
    line-height: 1.25rem;
    font-size: 0.875rem;
    font-weight: 500;
  `,
  inputWrapper: css`
    position: relative;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    margin-top: 0.25rem;
    border-radius: 0.375rem;
  `,
  input: css`
    display: block;
    width: 100%;
    background-color: #fff;
    border: 1px solid #d2d6dc;
    border-radius: 0.375rem;
    padding: 0.5rem 0.75rem;
    font-size: 1rem;
    line-height: 1.5;
    @media (min-width: 640px) {
      line-height: 1.25rem;
      font-size: 0.875rem;
    }
  `,
  inputError: css`
    border-color: #f56565;
  `,
  error: css`
    height: 1.5rem;
    line-height: 1.5;
    font-size: 0.75rem;
    color: #f56565;
  `,
}

export default Input
