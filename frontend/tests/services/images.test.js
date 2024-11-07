import images from '../../src/services/images'
import tokenService from '../../src/services/token'
import testImage from '../testdata/testImage.png'
import { expect, vi } from 'vitest'
import axios from 'axios'

vi.mock('axios')

tokenService.fetchToken = vi.fn().mockReturnValue('faketoken')

test('creates a new image correctly', async() => {
    const newImage = {
        entity: '66f5027d6430d371f8636c3c',
        note: 'this is a note',
        image: testImage
    }

    const config = {
        headers: { 
            Authorization: tokenService.fetchToken(),
            'Content-Type': 'multipart/form-data'
        },
    }


    const expected = {
        entity: newImage.entity,
        file_format: "png",
        note: 'this is a note',
    }

    axios.post.mockResolvedValue({ data: expected })

    const result = await images.create(newImage)

    expect(result).toMatchObject(expected)
    expect(axios.post).toHaveBeenCalledWith('/api/images', newImage, config)
})

test('deletes an image correctly', async () => {
    const imageId = '123'

    const config = {
        headers: { Authorization: tokenService.fetchToken() },
    }

    const expected = { message: 'Image deleted successfully' }

    axios.delete.mockResolvedValue({ data: expected })

    const result = await images.deleteImage(imageId)

    expect(result).toMatchObject(expected)
    expect(axios.delete).toHaveBeenCalledWith(`/api/images/${imageId}`, config)
})
