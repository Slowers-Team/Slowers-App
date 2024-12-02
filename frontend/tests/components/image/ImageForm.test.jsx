import { render, screen } from '@testing-library/react'
import ImageForm from '../../../src/components/image/ImageForm'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'


test('renders form', () => {
  const createImage = vi.fn()

  render(<ImageForm createImage={createImage}/>)

  const selector = screen.getByLabelText('Select image:')
  const note = screen.getByLabelText('Note:')
  const submit = screen.getByText('Save')
})

test('image can be selected', async() => {
  const createImage = vi.fn()
  const user = userEvent.setup()

  render(<ImageForm createImage={createImage}/>)

  const imageSelector = screen.getByLabelText('Select image:')
  const file = new File(['hello'], 'hello.png', {type: 'image/png'})

  await user.upload(imageSelector, file)

  expect(imageSelector.files[0]).toBe(file)
  expect(imageSelector.files.item(0)).toBe(file)
  expect(imageSelector.files).toHaveLength(1)
})

test('image note can be written', async() => {
  const createImage = vi.fn()
  const user = userEvent.setup()

  render(<ImageForm createImage={createImage}/>)

  const noteInput = screen.getByLabelText('Note:')

  await user.type(noteInput, "this is a note")

  expect(noteInput.value).toBe("this is a note")
})

test('note is required', async() => {
  const createImage = vi.fn()
  const user = userEvent.setup()

  render(<ImageForm createImage={createImage}/>)

  const submit = screen.getByText('Save')
  const imageSelector = screen.getByLabelText('Select image:')
  const file = new File(['hello'], 'hello.png', {type: 'image/png'})

  await user.upload(imageSelector, file)
  await user.click(submit)

  expect(imageSelector.files[0]).toBe(file)
  expect(imageSelector.files.length).toBe(1)
})

test('image is required', async() => {
  const createImage = vi.fn()
  const user = userEvent.setup()

  render(<ImageForm createImage={createImage}/>)

  const submit = screen.getByText('Save')
  const noteInput = screen.getByLabelText('Note:')
  const imageSelector = screen.getByLabelText('Select image:')

  await user.type(noteInput, "this is a note")
  await user.click(submit)

  expect(noteInput.value).toBe("this is a note")
  expect(imageSelector.files.length).toBe(0)
})
