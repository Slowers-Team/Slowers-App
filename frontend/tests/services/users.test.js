import users from '../../src/services/users'
import { expect, vi } from 'vitest'
import axios from 'axios'

vi.mock('axios')

test('creates a new user correctly', async() => {
    const newUser = {
        username: 'testuser',
        password: 'testpassword',
        email: 'testemail@email.com'
    }

    axios.post.mockResolvedValue({ data: newUser })

    const result = await users.create(newUser)

    expect(result).toEqual(newUser)
    expect(axios.post).toHaveBeenCalledWith('/api/register', newUser)
})