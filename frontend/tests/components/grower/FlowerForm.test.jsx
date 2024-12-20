/* eslint-disable no-unused-vars */
import { render, screen } from '@testing-library/react'
import FlowerForm from '../../../src/components/grower/FlowerForm'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders FlowerForm with name and latin name inputs', () => {
    const createFlower = vi.fn()

    render(<FlowerForm createFlower={createFlower} />)

    const name = screen.getByLabelText('Name:')
    const latinName = screen.getByLabelText('Latin name:')
    const qty = screen.getByLabelText('Qty:')
    const save = screen.getByText('Save')
})

test('updates input values when typing', async() => {
    const createFlower = vi.fn()
    const user = userEvent.setup()

    render(<FlowerForm createFlower={createFlower} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const flowerQtyInput = screen.getByLabelText('Qty:')

    await user.type(flowerNameInput, 'Sunflower')
    await user.type(flowerLatinNameInput, 'Helianthus annuus')
    await user.type(flowerQtyInput, '7')

    expect(flowerNameInput.value).toBe('Sunflower')
    expect(flowerLatinNameInput.value).toBe('Helianthus annuus')
    expect(flowerQtyInput.value).toBe('7')
})

test('resets input values after submit', async () => {
    const createFlower = vi.fn()
    const handleClose = vi.fn()
    const user = userEvent.setup()

    render(<FlowerForm createFlower={createFlower} handleClose={handleClose} />)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const flowerQtyInput = screen.getByLabelText('Qty:')
    const saveButton = screen.getByText('Save')

    await user.type(flowerNameInput, 'Lily')
    await user.type(flowerLatinNameInput, 'Lilium')
    await user.type(flowerQtyInput, '7')
    await user.click(saveButton)

    expect(flowerNameInput.value).toBe('')
    expect(flowerLatinNameInput.value).toBe('')
    expect(flowerQtyInput.value).toBe('0')
})

test('calls createFlower with correct values on submit', async () => {
    const createFlower = vi.fn()
    const handleClose = vi.fn()
    const user = userEvent.setup()

    render(<FlowerForm createFlower={createFlower} handleClose={handleClose}/>)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const flowerQtyInput = screen.getByLabelText('Qty:')
    const saveButton = screen.getByText('Save')

    await user.type(flowerNameInput, 'Rose')
    await user.type(flowerLatinNameInput, 'Rosa')
    await user.type(flowerQtyInput, '7')
    await user.click(saveButton)

    expect(createFlower.mock.calls).toHaveLength(1)
    expect(createFlower.mock.calls[0][0]).toEqual({ name: 'Rose', latin_name: 'Rosa', quantity: 7 })
})
