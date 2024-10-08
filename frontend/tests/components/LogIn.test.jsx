import { render, screen } from '@testing-library/react'
import LogIn from '../../src/components/LogIn'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

test('renders LogIn form with email and password inputs', () => {
    const login = vi.fn()

    render(<LogIn login={login} />)

    const email = screen.getByLabelText('Email:')
    const password = screen.getByLabelText('Password:')
    const submitButton = screen.getByRole('button', { name: 'Log In'})
})

test('updates input values when typing', async () => {
    const login = vi.fn()
    const user = userEvent.setup()

    render(<LogIn login={login} />)

    const emailInput = screen.getByLabelText('Email:')
    const passwordInput = screen.getByLabelText('Password:')

    await user.type(emailInput, 'test@email.com') 
    await user.type(passwordInput, 'testpassword')

    expect(emailInput.value).toBe('test@email.com')
    expect(passwordInput.value).toBe('testpassword')
})

test('does not clear input values after submit if email does not match standard format', async() => {
    const login = vi.fn()
    const user = userEvent.setup()

    render(<LogIn login={login} />)

    const emailInput = screen.getByLabelText('Email:')
    const passwordInput = screen.getByLabelText('Password:')

    await user.type(emailInput, 'invalidtestemail')
    await user.type(passwordInput, 'testpassword')
    await user.click(screen.getByRole('button', { name: 'Log In'}))

    expect(emailInput.value).toBe('invalidtestemail')
    expect(passwordInput.value).toBe('testpassword')
})