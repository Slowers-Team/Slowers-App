import { render, screen } from '@testing-library/react'
import FlowerList from '../../src/components/FlowerList'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders FlowerList without flowers', () => {
    const deleteFlower = vi.fn()
    render(<FlowerList flowers={[]} deleteFlower={deleteFlower} />)

    const name = screen.getByText('Name')
    const latinName = screen.getByText('Latin name')
    const addedTime = screen.getByText('Added time')
})

test('renders FlowerList with correct flowers and has the correct number of delete buttons', () => {
    const deleteFlower = vi.fn()

    const flowers = [
        {
            _id: '123',
            name: 'Sunflower',
            latin_name: 'Helianthus annuus',
            added_time: '1999-02-08T15:16:00.000Z'
        },
        {
            _id: '456',
            name: 'Lily',
            latin_name: 'Lilium',
            added_time: '2024-01-01T09:11:11.000Z'
        }
    ]

    render(<FlowerList flowers={flowers} deleteFlower={deleteFlower} />)

    expect(screen.getByText('Sunflower')).toBeInTheDocument()
    expect(screen.getByText('Helianthus annuus')).toBeInTheDocument()

    expect(screen.getByText('Lily')).toBeInTheDocument()
    expect(screen.getByText('Lilium')).toBeInTheDocument()

    expect(screen.getByText('8.2.1999 17:16')).toBeInTheDocument()
    expect(screen.getByText('1.1.2024 11:11')).toBeInTheDocument()

    const deleteButtons = screen.getAllByText('Delete')

    expect(deleteButtons.length).toBe(2)
})

test('calls deleteFlower when delete button is clicked', async () => {
    const deleteFlower = vi.fn()
    const user = userEvent.setup()

    const flowers = [
        {
            _id: '123',
            name: 'Sunflower',
            latin_name: 'Helianthus annuus',
            added_time: '1999-02-08T15:16:00.000Z'
        },
        {
            _id: '456',
            name: 'Lily',
            latin_name: 'Lilium',
            added_time: '2024-09-23T11:11:11.000Z'
        }
    ]

    render(<FlowerList flowers={flowers} deleteFlower={deleteFlower} />)

    const deleteButton = screen.getAllByText('Delete')[0]
    
    await user.click(deleteButton)

    expect(deleteFlower).toHaveBeenCalledWith(flowers[0])
})