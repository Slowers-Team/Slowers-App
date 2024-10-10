import { render, screen } from '@testing-library/react'
import RegisterForm from '../../src/components/RegisterForm'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders RegisterForm with username, password, and email inputs', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const username = screen.getByLabelText('Username:')
    const password = screen.getByLabelText('Password:')
    const email = screen.getByLabelText('Email:')
})

test('updates input values when typing', async () => {
    const createNewUser = vi.fn()
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')

    await user.type(usernameInput, 'testuser')
    await user.type(passwordInput, 'testpassword')
    await user.type(emailInput, 'testemail')

    expect(usernameInput.value).toBe('testuser')
    expect(passwordInput.value).toBe('testpassword')
    expect(emailInput.value).toBe('testemail')
})

test('clears input values after successful submit', async () => {
    const createNewUser = vi.fn()
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
    const growerRadioButton = screen.getByLabelText('Grower')
    const termsCheckbox = screen.getByLabelText('I agree to the terms and conditions')
    const submitButton = screen.getByText('Register')

    await user.type(usernameInput, 'testuser')
    await user.type(passwordInput, 'testpassword')
    await user.type(emailInput, 'testemail@email.com')
    await user.click(growerRadioButton)
    await user.click(termsCheckbox)
    await user.click(submitButton)

    expect(usernameInput.value).toBe('')
    expect(passwordInput.value).toBe('')
    expect(emailInput.value).toBe('')
})

test('does not clear input values after submit if email does not match standard format', async () => {
    const createNewUser = vi.fn()
    const user = userEvent.setup()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
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

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
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

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
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
