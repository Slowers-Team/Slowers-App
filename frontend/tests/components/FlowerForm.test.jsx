/* eslint-disable no-unused-vars */
import { fireEvent, render, screen } from '@testing-library/react'
import FlowerForm from '../../src/components/FlowerForm'
import { expect, vi } from 'vitest'

test('renders FlowerForm', () => {
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const name = screen.getByLabelText('Name:')
    const latinName = screen.getByLabelText('Latin name:')
    const save = screen.getByText('Save')
})

test('updates input values when typing', () => {
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')

    fireEvent.change(flowerNameInput, { target: { value: 'Sunflower' } })
    expect(flowerNameInput.value).toBe('Sunflower')

    fireEvent.change(flowerLatinNameInput, { target: { value: 'Helianthus annuus' } })
    expect(flowerLatinNameInput.value).toBe('Helianthus annuus')
})

test('clears input values after submit', () => {
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const saveButton = screen.getByText('Save')

    fireEvent.change(flowerNameInput, { target: { value: 'Lily' } })
    fireEvent.change(flowerLatinNameInput, { target: { value: 'Lilium' } })
    fireEvent.click(saveButton)

    expect(flowerNameInput.value).toBe('')
    expect(flowerLatinNameInput.value).toBe('')
})

test('calls createFlower with correct values on submit', () => {
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const saveButton = screen.getByText('Save')

    fireEvent.change(flowerNameInput, { target: { value: 'Rose' } })
    fireEvent.change(flowerLatinNameInput, { target: { value: 'Rosa' } })
    fireEvent.click(saveButton)

    expect(createFlower.mock.calls).toHaveLength(1)
    expect(createFlower.mock.calls[0][0]).toEqual({ name: 'Rose', latin_name: 'Rosa' })
})