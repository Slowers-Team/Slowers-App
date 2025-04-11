import userEvent from '@testing-library/user-event'
import { render, screen } from '@testing-library/react'
import CreateBusinessForm from '../../src/components/CreateBusinessForm'
import { expect, vi } from 'vitest'


test('sends form if fields are filled correctly', async () =>{
    const BusinessForm = vi.fn()
    const user = userEvent.setup()

    render(<CreateBusinessForm BusinessForm = {BusinessForm}/>)


    const Businessname = screen.getByPlaceholderText('Enter the name of the business')
    const BusinessID = screen.getByPlaceholderText('Enter business ID')
    const TypeOfTheBusiness = screen.getByText('Retailer')
    const BusinessEmail = screen.getByPlaceholderText('Enter email')
    const PhoneNumber = screen.getByPlaceholderText('Enter phone number')
    const Address = screen.getByPlaceholderText('Enter address')
    const PostalCode = screen.getByPlaceholderText('Enter postal code')
    const City = screen.getByPlaceholderText('Enter city/municipality')
    const AdditionalInformation = screen.getByPlaceholderText('Enter additional information')

    const CreateBusiness = screen.getByText('Create business')

    await user.type(Businessname, 'Flower Business') 
    await user.type(BusinessID, '1234567-8')
    await user.click(TypeOfTheBusiness)
    await user.type(PhoneNumber, '0501234567')
    await user.type(BusinessEmail, 'flower@business.com')
    await user.type(Address, 'Kukkatie 9A99')
    await user.type(PostalCode, '99999')
    await user.type(City, 'Korvatunturi')
    await user.type(AdditionalInformation, 'We sell flowers')

    await user.click(CreateBusiness)


    expect(BusinessForm.mock.calls).toHaveLength(1)


    
    
})

