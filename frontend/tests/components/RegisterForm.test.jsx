import { fireEvent, render, screen } from '@testing-library/react'
import RegisterForm from '../../src/components/RegisterForm'
import { expect, vi } from 'vitest'

test('renders RegisterForm', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const username = screen.getByLabelText('Username:')
    const password = screen.getByLabelText('Password:')
    const email = screen.getByLabelText('Email:')
})

test('updates input values when typing', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')

    fireEvent.change(usernameInput, { target: { value: 'testuser' } })
    expect(usernameInput.value).toBe('testuser')

    fireEvent.change(passwordInput, { target: { value: 'testpassword' } })
    expect(passwordInput.value).toBe('testpassword')

    fireEvent.change(emailInput, { target: { value: 'testemail' } })
    expect(emailInput.value).toBe('testemail')
})

test('clears input values after successful submit', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
    const submitButton = screen.getByText('Register')

    fireEvent.change(usernameInput, { target: { value: 'testuser' } })
    fireEvent.change(passwordInput, { target: { value: 'testpassword' } })
    fireEvent.change(emailInput, { target: { value: 'testemail@email.com' } })
    fireEvent.click(submitButton)

    expect(usernameInput.value).toBe('')
    expect(passwordInput.value).toBe('')
    expect(emailInput.value).toBe('')
})

test('does not clear input values after submit if email does not match standard format', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
    const submitButton = screen.getByText('Register')

    fireEvent.change(usernameInput, { target: { value: 'testuser' } })
    fireEvent.change(passwordInput, { target: { value: 'testpassword' } })
    fireEvent.change(emailInput, { target: { value: 'invalidtestemail' } })
    fireEvent.click(submitButton)

    expect(usernameInput.value).toBe('testuser')
    expect(passwordInput.value).toBe('testpassword')
    expect(emailInput.value).toBe('invalidtestemail')
})

test('calls createNewUser with correct values on submit', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
    const submitButton = screen.getByText('Register')

    fireEvent.change(usernameInput, { target: { value: 'testuser' } })
    fireEvent.change(passwordInput, { target: { value: 'testpassword' } })
    fireEvent.change(emailInput, { target: { value: 'testemail@email.com' } })
    fireEvent.click(submitButton)

    expect(createNewUser.mock.calls).toHaveLength(1)
    expect(createNewUser.mock.calls[0][0]).toEqual({
        username: 'testuser',
        password: 'testpassword',
        email: 'testemail@email.com'
    })
})

test('does not call createNewUser on submit if email does not match standard format', () => {
    const createNewUser = vi.fn()

    render(<RegisterForm createNewUser={createNewUser} />)

    const usernameInput = screen.getByLabelText('Username:')
    const passwordInput = screen.getByLabelText('Password:')
    const emailInput = screen.getByLabelText('Email:')
    const submitButton = screen.getByText('Register')

    fireEvent.change(usernameInput, { target: { value: 'testuser' } })
    fireEvent.change(passwordInput, { target: { value: 'testpassword' } })
    fireEvent.change(emailInput, { target: { value: 'invalidtestemail' } })
    fireEvent.click(submitButton)

    expect(createNewUser.mock.calls).toHaveLength(0)
})