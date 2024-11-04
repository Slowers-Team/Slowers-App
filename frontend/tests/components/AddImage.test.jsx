import { render, screen } from '@testing-library/react'
import AddImage from '../../src/components/image/AddImage'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders add image button', () => {
  const entity = '66f5027d6430d371f8636c3c'

  render(<AddImage entity={entity}/>)

  const buttonText = screen.getByText('Add a new image')
})

test('open imageform when clicking button', async () => {
  const entity = '66f5027d6430d371f8636c3c'
  const user = userEvent.setup()

  render(<AddImage entity={entity}/>)

  const imageButton = screen.getByText('Add a new image')
  await user.click(imageButton)

  const form = screen.getByText('Select image:')
  
  
})
