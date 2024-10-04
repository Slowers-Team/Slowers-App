import { render, screen } from '@testing-library/react'
import NewSiteForm from '../../src/components/NewSiteForm'
import { expect, test, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders NewSiteForm with name and notes inputs', () => {
    const createSite = vi.fn()

    render(<NewSiteForm createSite={createSite} />)

    const name = screen.getByLabelText('Name:')
    const note = screen.getByLabelText('Note:')
    const save = screen.getByText('Save')
})

test('updates input values when typing', async () => {
    const createSite = vi.fn()
    const user = userEvent.setup()

    render(<NewSiteForm createSite={createSite} />)

    const siteNameInput = screen.getByLabelText('Name:')
    const siteNoteInput = screen.getByLabelText('Note:')

    await user.type(siteNameInput, 'site test')
    await user.type(siteNoteInput, 'note test')

    expect(siteNameInput.value).toBe('site test')
    expect(siteNoteInput.value).toBe('note test')
})

test('calls createSite with correct values on submit', async () => {
    const createSite = vi.fn()
    const user = userEvent.setup()

    render(<NewSiteForm createSite={createSite} />)

    const siteNameInput = screen.getByLabelText('Name:')
    const siteNoteInput = screen.getByLabelText('Note:')
    const saveButton = screen.getByText('Save')

    await user.type(siteNameInput, 'site test')
    await user.type(siteNoteInput, 'note test')
    await user.click(saveButton)

    expect(createSite.mock.calls).toHaveLength(1)
    expect(createSite.mock.calls[0][0]).toEqual({ name: 'site test', note: 'note test' })
})

test('calls createSite with correct values on submit when note is empty', async () => {
    const createSite = vi.fn()
    const user = userEvent.setup()

    render(<NewSiteForm createSite={createSite} />)

    const siteNameInput = screen.getByLabelText('Name:')
    const siteNoteInput = screen.getByLabelText('Note:')
    const saveButton = screen.getByText('Save')

    await user.type(siteNameInput, 'site test')
    await user.click(saveButton)

    expect(createSite.mock.calls).toHaveLength(1)
    expect(createSite.mock.calls[0][0]).toEqual({ name: 'site test', note: '' })
})

test('clears input values after submit', async () => {
    const createSite = vi.fn()
    const user = userEvent.setup()

    render(<NewSiteForm createSite={createSite} />)

    const siteNameInput = screen.getByLabelText('Name:')
    const siteNoteInput = screen.getByLabelText('Note:')
    const saveButton = screen.getByText('Save')

    await user.type(siteNameInput, 'site test')
    await user.type(siteNoteInput, 'note test')
    await user.click(saveButton)

    expect(siteNameInput.value).toBe('')
    expect(siteNoteInput.value).toBe('')
})
