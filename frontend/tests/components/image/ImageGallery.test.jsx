import { render, screen, within } from '@testing-library/react'
import { test, vi } from 'vitest'
import userEvent from '@testing-library/user-event'
import ImageGallery from '../../../src/components/image/ImageGallery'

test('renders ImageGallery without site images', () => {
    const images = []
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()

    render(<ImageGallery images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const noImages = screen.getByText("This site doesn't have any images yet")
})

test('renders ImageGallery without flower images', () => {
    const images = []
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()

    render(<ImageGallery images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} type="flower" />)

    const noImages = screen.getByText("This flower doesn't have any images yet")
})

test('renders ImageGallery with one image', () => {
    const images = [{ _id: '1', url: 'flower.png' }]
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()

    render(<ImageGallery images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const image = screen.getByRole('img')

    expect(image).toHaveAttribute('src', 'flower.png')
})

test('renders ImageGallery with multiple images', () => {
    const images = [{ _id: '1', url: 'flower1.png' }, { _id: '2', url: 'flower2.png' }]
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()

    render(<ImageGallery images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const renderedImages = screen.getAllByRole('img')

    expect(renderedImages).toHaveLength(2)
    expect(renderedImages[0]).toHaveAttribute('src', 'flower1.png')
    expect(renderedImages[1]).toHaveAttribute('src', 'flower2.png')
})

test('Buttons are not rendered for retailer', () => {
    const images = [{ _id: '1', url: 'flower.png' }]
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()

    render(<ImageGallery isGrower={false} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const deleteButton = screen.queryByRole('button', {name: 'Delete'})
    const favoriteButton = screen.queryByRole('button', {name: 'Favorite'})

    expect(deleteButton).toBeNull()
    expect(favoriteButton).toBeNull()
})

test('Buttons are rendered for grower', () => {
    const images = [{ _id: '1', url: 'flower.png' }]
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()

    render(<ImageGallery isGrower={true} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const deleteButton = screen.getByRole('button', {name: 'Delete'})
    const favoriteButton = screen.getByRole('button', {name: 'Favorite'})
})

test('calls deleteImage when delete button is clicked', async () => {
    const images = [{ _id: '1', url: 'flower.png' }]
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()
    const user = userEvent.setup()

    render(<ImageGallery isGrower={true} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const deleteButton = screen.getByRole('button', {name: 'Delete'})
    await user.click(deleteButton)

    expect(deleteImage).toHaveBeenCalledWith(images[0])
})

test('calls favoriteImage when favorite button is clicked', async () => {
    const images = [{ _id: '1', url: 'flower1.png' }, { _id: '2', url: 'flower2.png' }]
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()
    const user = userEvent.setup()

    render(<ImageGallery isGrower={true} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const image2 = screen.getAllByRole('img')[1]
    const imageBox = image2.closest('.image-box')

    const favoriteButton = within(imageBox).getByRole('button', {name: 'Favorite'})
    await user.click(favoriteButton)

    expect(favoriteImage).toHaveBeenCalledWith(images[1])
})

test('favorite button is disabled if image is already favorited', async () => {
    const images = [{ _id: '1', url: 'flower1.png' }, { _id: '2', url: 'flower2.png' }]
    const deleteImage = vi.fn()
    const favoriteImage = vi.fn()
    const user = userEvent.setup()

    render(<ImageGallery isGrower={true} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} />)

    const image1 = screen.getAllByRole('img')[0]
    const imageBox = image1.closest('.image-box')

    const favoriteButton = within(imageBox).getByRole('button', {name: 'Favorite'})
    await user.click(favoriteButton)

    expect(favoriteImage).not.toHaveBeenCalled()
})
