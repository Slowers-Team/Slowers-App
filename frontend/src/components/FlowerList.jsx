import { useState } from "react"
import { useTranslation } from 'react-i18next'

const FlowerList = ({ flowers, deleteFlower }) => {
  const { t, i18n } = useTranslation()
  const [searchTerm, setSearchTerm] = useState("")
  const [sortConfig, setSortConfig] = useState({ key: '', direction: '' })

  const handleSort = (key) => {
    let direction = 'asc'
    if (sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc'
    }
    setSortConfig({ key, direction })
  }

  const renderSortIcon = (key) => {
    if (sortConfig.key === key) {
      return sortConfig.direction === 'asc' ? ' \u25B2' : ' \u25BC'
    }
    return ' \u25BE'
  }

  const sortedFlowers = [...flowers].sort((a, b) => {
    if (!sortConfig.key) return 0

    const aValue = a[sortConfig.key]
    const bValue = b[sortConfig.key]

    if (aValue < bValue) return sortConfig.direction === 'asc' ? -1 : 1
    if (aValue > bValue) return sortConfig.direction === 'asc' ? 1 : -1
    return 0
  })

  const filteredFlowers = sortedFlowers.filter(flower =>
    flower.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    flower.latin_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    new Date(flower.added_time).toLocaleDateString('fi').toLowerCase().includes(searchTerm.toLowerCase())
  )

  return (
    <div>
      <input
        type="text"
        placeholder={t('Search')}
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
      <table id="flowerList">
        <thead>
          <tr>
            <th onClick={() => handleSort('name')} style={{ cursor: 'pointer' }}>
              {t('flower.data.name')}
              {renderSortIcon('name')}
            </th>
            <th onClick={() => handleSort('latin_name')} style={{ cursor: 'pointer' }}>
              {t('flower.data.latinname')}
              {renderSortIcon('latin_name')}
            </th>
            <th onClick={() => handleSort('added_time')} style={{ cursor: 'pointer' }}>
              {t('flower.data.addedtime')}
              {renderSortIcon('added_time')}
            </th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {filteredFlowers.map(flower => {
            let addedTime = new Date(flower.added_time)

            let date = addedTime.toLocaleDateString('fi')
            let hour = addedTime.toLocaleString('fi', { hour: 'numeric' })
            let minute = addedTime.toLocaleString('fi', { minute: '2-digit' })
            let addedTimeStr = `${date} ${hour}:${minute}`

            return (
              <tr key={flower._id}>
                <td>{flower.name}</td>
                <td>
                  <em>{flower.latin_name}</em>
                </td>
                <td>{addedTimeStr}</td>
                <td>
                  <button id="deleteFlowerButton" onClick={() => deleteFlower(flower)}>
                    {t('button.delete')}
                  </button>
                </td>
              </tr>
            )
          })}
        </tbody>
      </table>
    </div>
  )
}

export default FlowerList
