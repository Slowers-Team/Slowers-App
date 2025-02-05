import { render, screen, waitFor } from '@testing-library/react'
import RegisterForm from '../../src/components/RegisterForm'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders RegisterForm with username, password, and email inputs', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const username = screen.getByPlaceholderText('Enter username')
    const password = screen.getByPlaceholderText('Enter password')
    const email = screen.getByPlaceholderText('Enter email')
})

test('updates input values when typing', async () => {
    const createNewUser = vi.fn()
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByPlaceholderText('Enter username')
    const passwordInput = screen.getByPlaceholderText('Enter password')
    const emailInput = screen.getByPlaceholderText('Enter email')

    await user.type(usernameInput, 'testuser')
    await user.type(passwordInput, 'testpassword')
    await user.type(emailInput, 'testemail')

    expect(usernameInput.value).toBe('testuser')
    expect(passwordInput.value).toBe('testpassword')
    expect(emailInput.value).toBe('testemail')
})

test('clears input values after successful submit', async () => {
    const createNewUser = vi.fn().mockResolvedValue({})
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByPlaceholderText('Enter username')
    const passwordInput = screen.getByPlaceholderText('Enter password')
    const emailInput = screen.getByPlaceholderText('Enter email')
    const growerRadioButton = screen.getByLabelText('Grower')
    const termsCheckbox = screen.getByLabelText('I agree to the terms and conditions')
    const submitButton = screen.getByText('Register')

    await user.type(usernameInput, 'testuser')
    await user.type(passwordInput, 'testpassword')
    await user.type(emailInput, 'testemail@email.com')
    await user.click(growerRadioButton)
    await user.click(termsCheckbox)
    await user.click(submitButton)

    await waitFor(() => {
        expect(usernameInput.value).toBe('')
        expect(passwordInput.value).toBe('')
        expect(emailInput.value).toBe('')
    })
})

test('does not clear input values after submit if email does not match standard format', async () => {
    const createNewUser = vi.fn()
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByPlaceholderText('Enter username')
    const passwordInput = screen.getByPlaceholderText('Enter password')
    const emailInput = screen.getByPlaceholderText('Enter email')
    const growerRadioButton = screen.getByLabelText('Grower')
    const termsCheckbox = screen.getByLabelText('I agree to the terms and conditions')
    const submitButton = screen.getByText('Register')

    await user.type(usernameInput, 'testuser')
    await user.type(passwordInput, 'testpassword')
    await user.type(emailInput, 'invalidtestemail')
    await user.click(growerRadioButton)
    await user.click(termsCheckbox)
    await user.click(submitButton)

    expect(usernameInput.value).toBe('testuser')
    expect(passwordInput.value).toBe('testpassword')
    expect(emailInput.value).toBe('invalidtestemail')
})

test('calls createNewUser with correct values on submit', async () => {
    const createNewUser = vi.fn()
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByPlaceholderText('Enter username')
    const passwordInput = screen.getByPlaceholderText('Enter password')
    const emailInput = screen.getByPlaceholderText('Enter email')
    const growerRadioButton = screen.getByLabelText('Grower')
    const termsCheckbox = screen.getByLabelText('I agree to the terms and conditions')
    const submitButton = screen.getByText('Register')

    await user.type(usernameInput, 'testuser')
    await user.type(passwordInput, 'testpassword')
    await user.type(emailInput, 'testemail@email.com')
    await user.click(growerRadioButton)
    await user.click(termsCheckbox)
    await user.click(submitButton)

    expect(createNewUser.mock.calls).toHaveLength(1)
    expect(createNewUser.mock.calls[0][0]).toEqual({
        username: 'testuser',
        password: 'testpassword',
        email: 'testemail@email.com',
        role: 'grower',
    })
})

test('does not call createNewUser on submit if email does not match standard format', async () => {
    const createNewUser = vi.fn()
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByPlaceholderText('Enter username')
    const passwordInput = screen.getByPlaceholderText('Enter password')
    const emailInput = screen.getByPlaceholderText('Enter email')
    const growerRadioButton = screen.getByLabelText('Grower')
    const termsCheckbox = screen.getByLabelText('I agree to the terms and conditions')
    const submitButton = screen.getByText('Register')

    await user.type(usernameInput, 'testuser')
    await user.type(passwordInput, 'testpassword')
    await user.type(emailInput, 'invalidtestemail')
    await user.click(growerRadioButton)
    await user.click(termsCheckbox)
    await user.click(submitButton)

    expect(createNewUser.mock.calls).toHaveLength(0)
})

test('shows error message after submit if email is already in use', async () => {
    const createNewUser = vi.fn(async (userData) => {
        // nyt käytössä kovakoodattu sähköposti
        // myöhemmin tarkoitus luoda testikäyttäjä testitietokantaan
        if (userData.email === 'testemail@email.com') {
            throw { response: { data: 'Email already exists' } }
        }
    })

    const user = userEvent.setup()
    render(<RegisterForm createNewUser={createNewUser} />)

    await user.type(screen.getByPlaceholderText('Enter username'), 'testuser')
    await user.type(screen.getByPlaceholderText('Enter password'), 'testpassword')
    await user.type(screen.getByPlaceholderText('Enter email'), 'testemail@email.com')
    await user.click(screen.getByLabelText('Grower'))
    await user.click(screen.getByLabelText('I agree to the terms and conditions'))
    await user.click(screen.getByText('Register'))

    const errorMessage = await screen.findByText('An error occurred. Please try again.')
    expect(errorMessage).toBeInTheDocument()
})