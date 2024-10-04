import { render, screen } from '@testing-library/react'
import SiteFlexbox from '../../src/components/SiteFlexbox'
import { expect, test, vi } from 'vitest'
import userEvent from '@testing-library/user-event'
import { MemoryRouter } from 'react-router-dom'

test('renders main SiteFlexbox without sites', () => {
    const createSite = vi.fn()

    render(
      <MemoryRouter>
          <SiteFlexbox createSite={createSite} sites={[]} />
      </MemoryRouter>
    )

    expect(screen.getByText('Flowers')).toBeInTheDocument()
    expect(screen.getByText('Add new site')).toBeInTheDocument()
})

test('renders main SiteFlexbox with site', () => {
    const createSite = vi.fn()  
    
    const mockSites = [
      {
          _id: '123',
          name: 'site 1',
          note: 'note 1',
          added_time: '2011-11-11T11:11:11.000Z'
      },
      {
          _id: '456',
          name: 'site 2',
          note: 'note 2',
          added_time: '2024-10-05T11:11:11.000Z'
      }
    ]

    render(
      <MemoryRouter>
          <SiteFlexbox createSite={createSite} sites={mockSites} />
      </MemoryRouter>
    )

    expect(screen.getByText('site 1')).toBeInTheDocument()
    expect(screen.getByText('site 2')).toBeInTheDocument()

    expect(screen.getByText('Note: note 1')).toBeInTheDocument()
    expect(screen.getByText('Note: note 2')).toBeInTheDocument()
})
