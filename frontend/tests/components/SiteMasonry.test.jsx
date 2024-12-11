import { render, screen, waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { BrowserRouter } from 'react-router-dom'
import SiteMasonry from '../../src/components/SiteMasonry'
import ImageService from '../../src/services/images'

test('renders SiteMasonry', async () => {
  const sites = [{ 
    _id: 'siteId', 
    name: 'siteName', 
    note: 'Note',
    favorite_image: 'imageId'}]

  ImageService.getByID = vi.fn(() => Promise.resolve('image'))

  render(
    <BrowserRouter>
      <SiteMasonry sites={sites} />
    </BrowserRouter>
  )

  const siteName = screen.getByText('siteName')
  const siteNote = screen.getByText('Note')

  await waitFor(() => {
    expect(screen.getByRole('img')).toHaveAttribute('src', 'image')
    expect(screen.getAllByRole('img')).toHaveLength(1)
  })
})

test('renders SiteMasonry with multiple sites', async () => {
  const sites = [{ 
    _id: 'siteId1', 
    name: 'siteName1', 
    note: 'Note1',
    favorite_image: 'imageId1'},
    { 
    _id: 'siteId2', 
    name: 'siteName2', 
    note: 'Note2',
    favorite_image: 'imageId2'}]

  ImageService.getByID = vi.fn(() => Promise.resolve('image'))

  render(
    <BrowserRouter>
      <SiteMasonry sites={sites} />
    </BrowserRouter>
  )

  const siteName1 = screen.getByText('siteName1')
  const siteNote1 = screen.getByText('Note1')
  const siteName2 = screen.getByText('siteName2')
  const siteNote2 = screen.getByText('Note2')

  await waitFor(() => {
    expect(screen.getAllByRole('img')).toHaveLength(2)
    expect(screen.getAllByRole('img')[0]).toHaveAttribute('src', 'image')
    expect(screen.getAllByRole('img')[1]).toHaveAttribute('src', 'image')
  })
})

test('renders SiteMasonry with no images', async () => {
  const sites = [{ 
    _id: 'siteId', 
    name: 'siteName', 
    note: 'Note',
    favorite_image: null}]

  ImageService.getByID = vi.fn(() => Promise.resolve('image'))

  render(
    <BrowserRouter>
      <SiteMasonry sites={sites} />
    </BrowserRouter>
  )

  const siteName = screen.getByText('siteName')
  const siteNote = screen.getByText('Note')

  await waitFor(() => {
    expect(screen.queryByRole('img')).toBeNull()
  })
})

test('renders SiteMasonry with image and empty', async () => {
  const sites = [{
    _id: 'siteId', 
    name: 'siteName', 
    note: 'Note',
    favorite_image: 'imageId'},
    {
    _id: 'siteId2', 
    name: 'siteName2', 
    note: 'Note2',
    favorite_image: null}]

  ImageService.getByID = vi.fn(() => Promise.resolve('image'))

  render(
    <BrowserRouter>
      <SiteMasonry sites={sites} />
    </BrowserRouter>
  )

  const siteName = screen.getByText('siteName')
  const siteNote = screen.getByText('Note')
  const siteName2 = screen.getByText('siteName2')
  const siteNote2 = screen.getByText('Note2')

  await waitFor(() => {
    expect(screen.getByRole('img')).toHaveAttribute('src', 'image')
    expect(screen.getAllByRole('img')).toHaveLength(1)
  })
})
