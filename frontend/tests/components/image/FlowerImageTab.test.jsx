import { render, screen } from '@testing-library/react'
import { test, vi } from 'vitest'
import userEvent from '@testing-library/user-event'
import FlowerImageTab from '../../../src/components/image/FlowerImageTab'

test('renders FlowerImageTab as grower', () => {
    const flower = { _id: 'flowerId' }
    const deleteImage = vi.fn()

    render(<FlowerImageTab isGrower={true} flower={flower} deleteImage={deleteImage} />)

    const addImageButton = screen.getByText('Add a new image')
    const noImages = screen.getByText("This flower doesn't have any images yet")
})

test('renders FlowerImageTab as retailer', () => {
    const flower = { _id: 'flowerId' }
    const deleteImage = vi.fn()

    render(<FlowerImageTab isGrower={false} flower={flower} deleteImage={deleteImage} />)

    const noImages = screen.getByText("This flower doesn't have any images yet")
})

test('open ImageForm when clicking button', async () => {
    const flower = { _id: 'flowerId' }
    const deleteImage = vi.fn()
    const user = userEvent.setup()

    render(<FlowerImageTab isGrower={true} flower={flower} deleteImage={deleteImage} />)

    const imageButton = screen.getByText('Add a new image')
    await user.click(imageButton)

    const selectImage = screen.getByText('Select image:')
    const note = screen.getByText('Note:')
})
