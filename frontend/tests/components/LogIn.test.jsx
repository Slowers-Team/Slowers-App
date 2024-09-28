import { fireEvent, render, screen } from "@testing-library/react";
import LogIn from "../../src/components/LogIn";
import { beforeEach, describe, expect, vi } from "vitest";

describe('Login form', () => {

    beforeEach(() => {
        const login = vi.fn()
        render(<LogIn login={login} />)
    })

    test('renders LogIn form with email and password inputs', () => {
        const email = screen.getByLabelText('Email:')
        const password = screen.getByLabelText('Password:')
        const submitButton = screen.getByRole('button', { name: 'Log In'})
    })

    test('updates input values when typing', () => {
        const emailInput = screen.getByLabelText('Email:')
        const passwordInput = screen.getByLabelText('Password:')

        fireEvent.change(emailInput, { target: { value: 'testemail' } })
        fireEvent.change(passwordInput, { target: { value: 'testpassword' } })

        expect(emailInput.value).toBe('testemail')
        expect(passwordInput.value).toBe('testpassword')
    })

    test('does not clear input values after submit if email does not match standard format', () => {
        const emailInput = screen.getByLabelText('Email:')
        const passwordInput = screen.getByLabelText('Password:')

        fireEvent.change(emailInput, { target: { value: 'invalidtestemail' } })
        fireEvent.change(passwordInput, { target: { value: 'testpassword' } })
        fireEvent.click(screen.getByRole('button', { name: 'Log In'}))
    
        expect(emailInput.value).toBe('invalidtestemail')
        expect(passwordInput.value).toBe('testpassword')
    })
})