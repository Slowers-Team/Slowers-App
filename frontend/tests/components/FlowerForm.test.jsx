/* eslint-disable no-unused-vars */
import { render, screen } from '@testing-library/react'
import FlowerForm from '../../src/components/FlowerForm'
import { vi } from 'vitest'

test('renders FlowerForm', () => {
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const name = screen.getByText('Name:')
    const latinName = screen.getByText('Latin name:')
    const save = screen.getByText('Save')
})
