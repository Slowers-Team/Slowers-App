import { render, screen } from '@testing-library/react'
import AddSite from '../../../src/components/grower/AddSite'
import userEvent from '@testing-library/user-event'
import { test, vi } from 'vitest'

test('renders add site button', () => {
  const createSite = vi.fn()

  render(<AddSite createSite={createSite}/>)

  const buttonText = screen.getByText('+ Add a new site')
})

test('open siteform when clicking button', async () => {
  const createSite = vi.fn()
  const user = userEvent.setup()

  render(<AddSite createSite={createSite}/>)

  const siteButton = screen.getByText('+ Add a new site')
  await user.click(siteButton)

  const name = screen.getByText('Name:')
  const note = screen.getByText('Note:')
  const saveButton = screen.getByText('Save')

  expect(name).toBeInTheDocument()
  expect(note).toBeInTheDocument()
  expect(saveButton).toBeInTheDocument()
})