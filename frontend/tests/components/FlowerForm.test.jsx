/* eslint-disable no-unused-vars */
import { render, screen } from '@testing-library/react'
import FlowerForm from '../../src/components/FlowerForm'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders FlowerForm with name and latin name inputs', () => {
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const name = screen.getByLabelText('Name:')
    const latinName = screen.getByLabelText('Latin name:')
    const save = screen.getByText('Save')
})

test('updates input values when typing', async() => {
    const user = userEvent.setup()
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')

    await user.type(flowerNameInput, 'Sunflower')
    await user.type(flowerLatinNameInput, 'Helianthus annuus')

    expect(flowerNameInput.value).toBe('Sunflower')
    expect(flowerLatinNameInput.value).toBe('Helianthus annuus')
})

test('clears input values after submit', async () => {
    const createFlower = vi.fn()
    const user = userEvent.setup()

    render(<FlowerForm createFlower={createFlower} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const saveButton = screen.getByText('Save')

    await user.type(flowerNameInput, 'Lily')
    await user.type(flowerLatinNameInput, 'Lilium')
    await user.click(saveButton)

    expect(flowerNameInput.value).toBe('')
    expect(flowerLatinNameInput.value).toBe('')
})

test('calls createFlower with correct values on submit', async () => {
    const createFlower = vi.fn()
    const user = userEvent.setup()

    render(<FlowerForm createFlower={createFlower} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const saveButton = screen.getByText('Save')

    await user.type(flowerNameInput, 'Rose')
    await user.type(flowerLatinNameInput, 'Rosa')
    await user.click(saveButton)

    expect(createFlower.mock.calls).toHaveLength(1)
    expect(createFlower.mock.calls[0][0]).toEqual({ name: 'Rose', latin_name: 'Rosa' })
})