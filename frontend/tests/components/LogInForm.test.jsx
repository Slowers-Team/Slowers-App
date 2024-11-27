import { render, screen } from '@testing-library/react'
import LogInForm from '../../src/components/LogInForm'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'
import { RouterProvider, createMemoryRouter } from 'react-router-dom'

test('renders LogIn form with email and password inputs', () => {
    const router = createMemoryRouter([
        { path: "/login", element: <LogInForm />, action() {return {invalidLogin: true}} }
    ], {initialEntries: ["/login"]} )
    render(<RouterProvider router={router}/>)

    const email = screen.getByLabelText('Email')
    const password = screen.getByLabelText('Password')
    const submitButton = screen.getByRole('button', { name: 'Log in'})
})

test('updates input values when typing', async () => {
    const user = userEvent.setup()

    const router = createMemoryRouter([
        { path: "/login", element: <LogInForm />, action() {return {invalidLogin: true}} }
    ], {initialEntries: ["/login"]} )
    render(<RouterProvider router={router}/>)

    const emailInput = screen.getByLabelText('Email')
    const passwordInput = screen.getByLabelText('Password')

    await user.type(emailInput, 'test@email.com') 
    await user.type(passwordInput, 'testpassword')

    expect(emailInput.value).toBe('test@email.com')
    expect(passwordInput.value).toBe('testpassword')
})

test('does not clear input values after submit if email does not match standard format', async() => {
    const user = userEvent.setup()

    const router = createMemoryRouter([
        { path: "/login", element: <LogInForm />, action() {return {invalidLogin: true}} }
    ], {initialEntries: ["/login"]} )
    render(<RouterProvider router={router}/>)


    const emailInput = screen.getByLabelText('Email')
    const passwordInput = screen.getByLabelText('Password')

    await user.type(emailInput, 'invalidtestemail')
    await user.type(passwordInput, 'testpassword')
    await user.click(screen.getByRole('button', { name: 'Log in'}))

    expect(emailInput.value).toBe('invalidtestemail')
    expect(passwordInput.value).toBe('testpassword')
})
