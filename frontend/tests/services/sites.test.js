import sites from '../../src/services/sites'
import tokenService from '../../src/services/token'
import { expect, test, vi } from 'vitest'
import axios from 'axios'

vi.mock('axios')

tokenService.fetchToken = vi.fn().mockReturnValue("faketoken")

const config = {
  headers: { Authorization: tokenService.fetchToken() },
}

test('returns no sites when database has no sites', async() => {
    const mockSites = []

    axios.get.mockResolvedValue({ data: mockSites })
    
    const result = await sites.get()

    expect(result).toEqual(mockSites)
    expect(result).length(0)
    expect(axios.get).toHaveBeenCalledWith('/api/sites', config)
})

test('returns correct sites when database has sites', async() => {
    const mockSites = [
        {
            _id: '123',
            name: 'site 1',
            note: 'note 1',
            added_time: '2011-11-11T11:11:11.000Z'
        },
        {
            _id: '456',
            name: 'site 2',
            note: 'note 2',
            added_time: '2024-10-05T11:11:11.000Z'
        }
    ]

    axios.get.mockResolvedValue({ data: mockSites })

    const result = await sites.get()

    expect(result).toEqual(mockSites)
    expect(result).length(2)
    expect(axios.get).toHaveBeenCalledWith('/api/sites', config)
})

test('creates a site correctly and uses the correct url', async() => {
    const newSite = [
      {
          _id: '123',
          name: 'site 1',
          note: 'note 1',
          added_time: '2011-11-11T11:11:11.000Z'
      },
    ]
    
    axios.post.mockResolvedValue({ data: newSite })

    const result = await sites.create(newSite)

    expect(result).toEqual(newSite)
    expect(axios.post).toHaveBeenCalledWith('/api/sites', newSite, config)
})

test('removes a site correctly and uses the correct url', async() => {
    const mockSites = [
        {
            _id: '123',
            name: 'site 1',
            note: 'note 1',
            added_time: '2011-11-11T11:11:11.000Z'
        },
        {
            _id: '456',
            name: 'site 2',
            note: 'note 2',
            added_time: '2024-10-05T11:11:11.000Z'
        }
    ]

    const siteid = mockSites[0]._id

    axios.get.mockResolvedValue({ data: mockSites })
    axios.delete.mockResolvedValue({ data: mockSites[0] })

    await sites.remove(siteid)

    expect(axios.delete).toHaveBeenCalledWith('/api/sites/' + siteid, config)
})
