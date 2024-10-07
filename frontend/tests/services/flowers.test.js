import flowers from '../../src/services/flowers'
import tokenService from '../../src/services/token'
import { expect, vi } from 'vitest'
import axios from 'axios'

vi.mock('axios')

tokenService.fetchToken = vi.fn().mockReturnValue('faketoken')

const config = {
    headers: { Authorization: tokenService.fetchToken() },
}

test('returns no flowers when database has no flowers', async() => {
    const mockFlowers = []

    axios.get.mockResolvedValue({ data: mockFlowers })
    
    const result = await flowers.getAll()

    expect(result).toEqual(mockFlowers)
    expect(result).length(0)
    expect(axios.get).toHaveBeenCalledWith('/api/flowers', config)
})

test('returns correct flowers when database has flowers', async() => {
    const mockFlowers = [
        {
            _id: '123',
            name: 'Sunflower',
            latin_name: 'Helianthus annuus',
            added_time: '1999-02-08T15:16:00.000Z'
        },
        {
            _id: '456',
            name: 'Lily',
            latin_name: 'Lilium',
            added_time: '2024-09-23T11:11:11.000Z'
        }
    ]

    axios.get.mockResolvedValue({ data: mockFlowers })

    const result = await flowers.getAll()

    expect(result).toEqual(mockFlowers)
    expect(result).length(2)
    expect(axios.get).toHaveBeenCalledWith('/api/flowers', config)
})

test('creates a flower correctly and uses the correct url', async() => {
    const newFlower = {
        name: 'Lily',
        latin_name: 'Lilium',
        added_time: '2024-09-23T11:11:11.000Z'
    }

    axios.post.mockResolvedValue({ data: newFlower })

    const result = await flowers.create(newFlower)

    expect(result).toEqual(newFlower)
    expect(axios.post).toHaveBeenCalledWith('/api/flowers', newFlower, config)
})

test('deletes a flower correctly and uses the correct url', async() => {
    const mockFlowers = [
        {
            _id: '123',
            name: 'Sunflower',
            latin_name: 'Helianthus annuus',
            added_time: '1999-02-08T15:16:00.000Z'
        },
        {
            _id: '456',
            name: 'Lily',
            latin_name: 'Lilium',
            added_time: '2024-09-23T11:11:11.000Z'
        }
    ]

    const sunflowerId = mockFlowers[0]._id
    
    axios.get.mockResolvedValue({ data: mockFlowers })
    axios.delete.mockResolvedValue({ data: mockFlowers[0] })

    await flowers.remove(sunflowerId)

    expect(axios.delete).toHaveBeenCalledWith('/api/flowers/' + sunflowerId, config)
})
