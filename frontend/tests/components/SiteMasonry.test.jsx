import { render, screen } from '@testing-library/react'
import { test } from 'vitest'
import { BrowserRouter } from 'react-router-dom'
import SiteMasonry from '../../src/components/SiteMasonry'

test('renders SiteMasonry', () => {
  const sites = [{ 
    _id: 'siteId', 
    name: 'siteName', 
    note: 'Note',
    favorite_image: 'imageId'}]

  render(
    <BrowserRouter>
      <SiteMasonry sites={sites} />
    </BrowserRouter>
  )

  const siteName = screen.getByText('siteName')
  const siteNote = screen.getByText('Note')
})

