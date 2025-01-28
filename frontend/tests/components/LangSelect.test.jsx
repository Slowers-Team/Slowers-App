/* eslint-disable no-unused-vars */
import { render, screen } from '@testing-library/react'
import LangSelect from '../../src/components/LangSelect'
import { expect } from 'vitest'
import userEvent from '@testing-library/user-event'
import i18n from '../../src/i18n'

test('renders LangSelect with "en" and "fi" options', () => {
  render(<LangSelect />)

  const en = screen.getByRole('button', { name: /🇬🇧 English/i })
  const fi = screen.getByRole('button', { name: /🇫🇮 Suomi/i })
  const sv = screen.getByRole('button', { name: /🇸🇪 Svenska/i })
})

test('UI language is English by default', () => {
  expect(i18n.resolvedLanguage).toBe('en')
})

test('changes UI language to Finnish when "fi" is clicked', async() => {
  const user = userEvent.setup()

  render(<LangSelect />)

  const fi = screen.getByRole('button', { name: /🇫🇮 Suomi/i })
  await user.click(fi)

  expect(i18n.resolvedLanguage).toBe('fi')
})

test('changes UI language to Swedish when "sv" is clicked', async() => {
  const user = userEvent.setup()

  render(<LangSelect />)

  const sv = screen.getByRole('button', { name: /🇸🇪 Svenska/i })
  await user.click(sv)

  expect(i18n.resolvedLanguage).toBe('sv')
})
