import { render, screen, fireEvent, waitFor, within } from '@testing-library/react'
import RetailerFlowerList from '../../src/components/retailer/RetailerFlowerList'
import { expect } from 'vitest'


test ('dummy test', async () => {
  expect(true).toEqual(true)
})
test ('Searches correctly when Scientific name search box is checked',async()=>{
  const flowers = [
    { _id: '456', name: 'Lily', latin_name: 'Lilium', added_time: '2024-01-01T09:11:11.000Z', grower_email: 'rosa@example.com', quantity: 5 },
    { _id: '789', name: 'Targetflower', latin_name: 'Rosa', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 },
    { _id: '790', name: 'Rosa', latin_name: 'Wolfsbane', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 }
]
  render(<RetailerFlowerList flowers={flowers} />)
  const scientificnamecheckbox = screen.getByTestId('scientificnamecheckbox')
  expect(scientificnamecheckbox).not.toBeChecked();
  fireEvent.click(scientificnamecheckbox)
  expect(scientificnamecheckbox).toBeChecked();
  const searchInput = screen.getByPlaceholderText('Search')
  fireEvent.change(searchInput, { target: { value: 'Rosa' } })

  await waitFor(() => {
    expect(screen.getByText('Targetflower')).toBeInTheDocument()
    expect(screen.queryByText('Wolfsbane')).toBeNull()
    expect(screen.queryByText('rosa@example.com')).toBeNull()

  })

})

  test ('Searches correctly when flower name search box is checked',async()=>{
    const flowers = [
      { _id: '456', name: 'Lily', latin_name: 'Lilium', added_time: '2024-01-01T09:11:11.000Z', grower_email: 'rosa@example.com', quantity: 5 },
      { _id: '789', name: 'Ruusu', latin_name: 'Rosa', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 },
      { _id: '790', name: 'Rosa', latin_name: 'Targetflower', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 }
  ]
    render(<RetailerFlowerList flowers={flowers} />)
    const flowernamecheckbox = screen.getByTestId('flowernamecheckbox')
    expect(flowernamecheckbox).not.toBeChecked();
    fireEvent.click(flowernamecheckbox)
    expect(flowernamecheckbox).toBeChecked();
    const searchInput = screen.getByPlaceholderText('Search')
    fireEvent.change(searchInput, { target: { value: 'Rosa' } })
  
    await waitFor(() => {
      expect(screen.getByText('Targetflower')).toBeInTheDocument()
      expect(screen.queryByText('Ruusu')).toBeNull()
      expect(screen.queryByText('rosa@example.com')).toBeNull()
    })

  
}
)

test ('Searches correctly when grower search box is checked',async()=>{
  const flowers = [
    { _id: '456', name: 'Targetflower', latin_name: 'Lilium', added_time: '2024-01-01T09:11:11.000Z', grower_email: 'rosa@example.com', quantity: 5 },
    { _id: '789', name: 'Ruusu', latin_name: 'Rosa', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 },
    { _id: '790', name: 'Rosa', latin_name: 'Wolfsbane', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 }
]
  render(<RetailerFlowerList flowers={flowers} />)
  const growercheckbox = screen.getByTestId('growercheckbox')
  expect(growercheckbox).not.toBeChecked();
  fireEvent.click(growercheckbox)
  expect(growercheckbox).toBeChecked();
  const searchInput = screen.getByPlaceholderText('Search')
  fireEvent.change(searchInput, { target: { value: 'Rosa' } })

  await waitFor(() => {
    expect(screen.getByText('Targetflower')).toBeInTheDocument()
    expect(screen.queryByText('Ruusu')).toBeNull()
    expect(screen.queryByText('Wolfsbane')).toBeNull()
  })


}
)
test('searches correctly when typing in the search box', async () => {
    const flowers = [
        { _id: '123', name: 'Sunflower', latin_name: 'Helianthus annuus', added_time: '1999-02-08T15:16:00.000Z', grower_email: 'grower1@example.com', quantity: 10 },
        { _id: '456', name: 'Lily', latin_name: 'Lilium', added_time: '2024-01-01T09:11:11.000Z', grower_email: 'grower2@example.com', quantity: 5 },
        { _id: '789', name: 'Rose', latin_name: 'Rosa', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 }
    ]

  render(<RetailerFlowerList flowers={flowers} />)

  const searchInput = screen.getByPlaceholderText('Search')
  fireEvent.change(searchInput, { target: { value: 'Lily' } })


  await waitFor(() => {
    expect(screen.getByText('Lily')).toBeInTheDocument()
    expect(screen.queryByText('Rose')).toBeNull()
  })

})

test('sorts flowers correctly when clicking the name', async () => {
    const flowers = [
      { _id: '123', name: 'Sunflower', latin_name: 'Helianthus annuus', added_time: '1999-02-08T15:16:00.000Z', grower_email: 'grower1@example.com', quantity: 10 },
      { _id: '456', name: 'Lily', latin_name: 'Lilium', added_time: '2024-01-01T09:11:11.000Z', grower_email: 'grower2@example.com', quantity: 5 },
      { _id: '789', name: 'Rose', latin_name: 'Rosa', added_time: '2010-06-14T13:45:00.000Z', grower_email: 'grower3@example.com', quantity: 8 }
    ]
  
    render(<RetailerFlowerList flowers={flowers} />)

    const nameHeaderList = screen.queryAllByText((content, element) => {
      const hasText = content.includes('Name')
      return hasText
    })
    const nameHeader = nameHeaderList[1]
  
    fireEvent.click(nameHeader)
  
    const rows = screen.getAllByRole('row')
    const sortedNames = rows.slice(1).map(row => {
      return within(row).getByText(/Sunflower|Lily|Rose/).textContent
    })
  
    expect(sortedNames).toEqual(['Lily', 'Rose', 'Sunflower']) 
})
