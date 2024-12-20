import { render, screen } from '@testing-library/react'
import AddFlower from '../../../src/components/grower/AddFlower'
import userEvent from '@testing-library/user-event'

test('renders add flower button', () => {
    const createFlower = vi.fn()

    render(<AddFlower createFlower={createFlower}/>)

    const buttonText = screen.getByText('+ Add a new flower')
})

test('open flower form when clicking button', async () => {
    const createFlower = vi.fn()
    const user = userEvent.setup()

    render(<AddFlower createFlower={createFlower}/>)

    const flowerButton = screen.getByText('+ Add a new flower')
    await user.click(flowerButton)

    const flowerNameInput = screen.getByLabelText('Name:')
    const flowerLatinNameInput = screen.getByLabelText('Latin name:')
    const saveButton = screen.getByText('Save')
})
