import { render, screen } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import userEvent from '@testing-library/user-event'
import FlowerInfoTab from '../../../src/components/flower/FlowerInfoTab'

test('renders FlowerInfoTab as grower', () => {
    const flower = { 
        _id: 'flowerId',
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        added_time: '2002-02-20T20:02:00.000Z',
        grower_email: 'grower1@example.com',
        quantity: 7,
        visible: false
     }
    const deleteImage = vi.fn()

    render(<FlowerInfoTab isGrower={true} flower={flower} deleteImage={deleteImage} />)

    const name = screen.getByText('Name')
    const latinName = screen.getByText('Latin name')
    const addedTime = screen.getByText('Added time')    
    const site = screen.getByText('Site')
    const qty = screen.getByText('Qty')
    const visible = screen.getByText('Visible to retailers')
})

test('renders FlowerInfoTab as retailer', () => {
    const flower = { 
        _id: 'flowerId',
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        added_time: '2002-02-20T20:02:00.000Z',
        grower_email: 'grower1@example.com',
        quantity: 7,
        visible: false
     }
    const deleteImage = vi.fn()

    render(<FlowerInfoTab isGrower={false} flower={flower} deleteImage={deleteImage} />)

    const name = screen.getByText('Name')
    const latinName = screen.getByText('Latin name')
    const addedTime = screen.getByText('Added time')    
    const site = screen.getByText('Grower')
    const qty = screen.getByText('Qty')
})

test('open ModifyFlowerForm when clicking button', async () => {
    const flower = { 
        _id: 'flowerId',
        name: 'Sunflower',
        latin_name: 'Helianthus annuus',
        added_time: '2002-02-20T20:02:00.000Z',
        grower_email: 'grower1@example.com',
        quantity: 7,
        visible: false
     }
    const deleteImage = vi.fn()
    const modifyFlower = vi.fn()
    const user = userEvent.setup()

    render(<FlowerInfoTab isGrower={true} flower={flower} deleteImage={deleteImage} modifyFlower={modifyFlower} />)

    const modifyButton = screen.getByText('Edit')
    await user.click(modifyButton)

    const save = screen.getByText('Save')
    const cancel = screen.getByText('Cancel')
    expect(save).toBeInTheDocument()
    expect(cancel).toBeInTheDocument()
})
