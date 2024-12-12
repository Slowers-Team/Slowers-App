import { render, screen } from '@testing-library/react'
import { test, vi } from 'vitest'
import { BrowserRouter } from 'react-router-dom'
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
})
