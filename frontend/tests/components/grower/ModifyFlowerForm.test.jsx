/* eslint-disable no-unused-vars */
import { render, screen } from '@testing-library/react'
import ModifyFlowerForm from '../../../src/components/grower/ModifyFlowerForm'
import { expect, test, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders ModifyFlowerFrom with name, latin name and qty inputs', () => {
    const modifyFlower = vi.fn()
    const flower = {
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        quantity: 7
    }

    render(<ModifyFlowerForm modifyFlower={modifyFlower} flower={flower} />)

    const name = screen.getByLabelText('Name')
    const latinName = screen.getByLabelText('Latin name')
    const addedTime = screen.getByText('Added time')
    const site = screen.getByText('Site')
    const qty = screen.getByLabelText('Quantity')
    const visible = screen.getByText('Visible to retailers')
    const del = screen.getByText('Delete')
    const save = screen.getByText('Save')
    const cancel = screen.getByText('Cancel')
})

test ('updates input values when typing', async() => {
    const modifyFlower = vi.fn()
    const flower = {
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        quantity: 7
    }
    const user = userEvent.setup()

    render(<ModifyFlowerForm modifyFlower={modifyFlower} flower={flower} />)
    
    const name = screen.getByLabelText('Name')
    const latinName = screen.getByLabelText('Latin name')
    const qty = screen.getByLabelText('Quantity')

    await user.clear(name)
    await user.type(name, 'Rose')
    await user.clear(latinName)
    await user.type(latinName, 'Rosa')
    await user.clear(qty)
    await user.type(qty, '10')

    expect(name.value).toBe('Rose')
    expect(latinName.value).toBe('Rosa')
    expect(qty.value).toBe('10')
})

test('calls modifyFlower with correct values on submit', async () => {
    const modifyFlower = vi.fn()
    const handleFlowerModify = vi.fn()
    const flower = {
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        quantity: 7
    }
    const user = userEvent.setup()

    render(<ModifyFlowerForm modifyFlower={modifyFlower} flower={flower} handleFlowerModify={handleFlowerModify} />)

    const name = screen.getByLabelText('Name')
    const latinName = screen.getByLabelText('Latin name')
    const qty = screen.getByLabelText('Quantity')
    const save = screen.getByText('Save')

    await user.clear(name)
    await user.clear(latinName)
    await user.clear(qty)
    await user.type(name, 'Rose')
    await user.type(latinName, 'Rosa')
    await user.type(qty, '10')
    await user.click(save)

    expect(modifyFlower.mock.calls).toHaveLength(1)
    expect(modifyFlower).toHaveBeenCalledWith({
        name: 'Rose',
        latin_name: 'Rosa',
        quantity: 10
    })
})

test('does not call modifyFlower if name is empty', async () => {
    const modifyFlower = vi.fn()
    const handleFlowerModify = vi.fn()
    const flower = {
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        quantity: 7
    }
    const user = userEvent.setup()

    render(<ModifyFlowerForm modifyFlower={modifyFlower} flower={flower} handleFlowerModify={handleFlowerModify} />)

    const name = screen.getByLabelText('Name')
    const save = screen.getByText('Save')

    await user.clear(name)
    await user.click(save)

    expect(modifyFlower.mock.calls).toHaveLength(0)
})

test('does not call modifyFlower if qty is invalid input', async () => {
    const modifyFlower = vi.fn()
    const handleFlowerModify = vi.fn()
    const flower = {
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        quantity: 7
    }
    const user = userEvent.setup()

    render(<ModifyFlowerForm modifyFlower={modifyFlower} flower={flower} handleFlowerModify={handleFlowerModify} />)

    const qty = screen.getByLabelText('Quantity')
    const save = screen.getByText('Save')

    await user.clear(qty)
    await user.type(qty, 'abc')
    await user.click(save)

    expect(modifyFlower.mock.calls).toHaveLength(0)

    await user.clear(qty)
    await user.type(qty, '-10')
    await user.click(save)

    expect(modifyFlower.mock.calls).toHaveLength(0)
})