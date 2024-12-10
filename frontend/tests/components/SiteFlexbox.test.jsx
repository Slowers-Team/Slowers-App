import { render, screen } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { BrowserRouter } from 'react-router-dom'
import userEvent from '@testing-library/user-event'
import SiteFlexbox from '../../src/components/SiteFlexbox'

test('renders SiteFlexbox', () => {
  const sites = [{ 
    _id: 'siteId', 
    name: 'siteName', 
    favorite_image: 'imageId'}]
  const createSite = vi.fn()

  render(
    <BrowserRouter>
      <SiteFlexbox createSite={createSite} sites={sites} />
    </BrowserRouter>
  )

  const siteName = screen.getByText('siteName')
  const addSite = screen.getByText('Add a new site')
})

test('open SiteForm when clicking button', async () => {
  const sites = [{ 
    _id: 'siteId', 
    name: 'siteName', 
    favorite_image: 'imageId'}]
  const createSite = vi.fn()
  const user = userEvent.setup()

  render(
    <BrowserRouter>
      <SiteFlexbox createSite={createSite} sites={sites} />
    </BrowserRouter>
  )

  const addSite = screen.getByText('Add a new site')
  await user.click(addSite)

  const name = screen.getByText('Name:')
  const note = screen.getByText('Note:')
  const saveButton = screen.getByText('Save')
})
