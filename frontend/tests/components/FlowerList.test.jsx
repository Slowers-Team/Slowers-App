import { render, screen } from '@testing-library/react'
import FlowerList from '../../src/components/FlowerList'
import { vi } from 'vitest'

test('renders FlowerList without flowers', () => {
    const deleteFlower = vi.fn()
    render(<FlowerList flowers={[]} deleteFlower={deleteFlower} />)

    const name = screen.getByText('Name')
    const latinName = screen.getByText('Latin name')
    const addedTime = screen.getByText('Added time')
})