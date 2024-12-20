import '../../layouts/Retailer.css'
import { useTranslation } from 'react-i18next'
import FlowerModal from '../flower/FlowerModal'
import { useState, useEffect } from "react"
import { Button, Table } from 'react-bootstrap'
import ImageService from '../../services/images'
import '../../App.css'
import { formatTime } from '../../utils.js'

const RetailerFlowerList = ({ flowers }) => {
  const { t, i18n } = useTranslation()
  const [showModal, setShowModal] = useState(false)
  const [currentFlower, setCurrentFlower] = useState("")
  const [searchTerm, setSearchTerm] = useState("")
  const [sortConfig, setSortConfig] = useState({ key: '', direction: '' })
  const [images, setImages] = useState([])

  useEffect(() => {
    const newImages = Promise.all(flowers.map((f) => {
      if (f.favorite_image) {
        return ImageService.getByID(f.favorite_image)
          .then((url) => (
            {flower: f._id, url: url}
          ))
          .catch((error) => console.error("error fetching image for flower:", f, error))
      }
    }))

    newImages.then((imgs) => setImages(imgs.filter((x)=>x)))
  
  }, [flowers])

  
  const handleShow = (flower) => {
    setShowModal(true)
    setCurrentFlower(flower)
  }

  const handleClose = () => {
    setShowModal(false)
    setCurrentFlower("")
  }

  const handleSort = (key) => {
    let direction = 'asc'
    if (sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc'
    }
    setSortConfig({ key, direction })
  }

  const renderSortIcon = (key) => {
    if (sortConfig.key === key) {
      return  (
          sortConfig.direction === 'asc' ? 
            <span id="sort-icon">
              <i className="bi bi-caret-up-fill" id="sort-icon-up"></i>
            </span> 
            : 
            <span id="sort-icon">
              <i className="bi bi-caret-down-fill" id="sort-icon-down"></i>
            </span>
      )
    }
    return (
      <span id="sort-icon">
        <i className="bi bi-caret-down-fill" id="sort-icon-down"></i>
        <i className="bi bi-caret-up-fill" id="sort-icon-up"></i>
      </span>
    )
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
    flower.grower_email.toLowerCase().includes(searchTerm.toLowerCase()) ||
    flower.quantity.toString().includes(searchTerm.toLowerCase()) ||
    new Date(flower.added_time).toLocaleDateString('fi').toLowerCase().includes(searchTerm.toLowerCase()) ||
    new Date(flower.added_time).toLocaleString('fi', { hour: 'numeric', minute: '2-digit' }).toLowerCase().includes(searchTerm.toLowerCase())
  )
  
  return (
    <div className="retailerFlowerList">
      <div className="d-flex justify-content-start mb-3 input-wrapper">
        <input
          type="text"
          placeholder={t('button.Search')}
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>
      <table id="retailerFlowerList" className="table table-hover align-middle">
        <thead>
          <tr>
            <th>{t('flower.data.image')}</th>
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
            <th onClick={() => handleSort('grower_email')} style={{ cursor: 'pointer' }}>
              {t('flower.data.grower')}
              {renderSortIcon('grower_email')}
            </th>
            <th onClick={() => handleSort('quantity')} style={{ cursor: 'pointer' }}>
              {t('flower.data.qty')}
              {renderSortIcon('quantity')}
            </th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {filteredFlowers.map(flower => {
            return (
              <tr key={flower._id}>
                <td className='image-cell'>
                  <div className='image-container'>
                    {images.find((o) => o.flower === flower._id)?.url && 
                    <img src={images.find((o) => o.flower === flower._id)?.url} alt={flower.name} />
                    }
                  </div>
                </td>
                <td>{flower.name}</td>
                <td>
                  <em>{flower.latin_name}</em>
                </td>
                <td>{formatTime(flower.added_time)}</td>
                <td>{flower.grower_email}</td>
                <td>{flower.quantity}</td>
                <td>
                  <button id='showFlowerPageButton' className="custom-button" onClick={() => handleShow(flower)}>
                  <i className="bi bi-info-circle-fill"></i>
                  </button>
                </td>
              </tr>
            )
          })}
        </tbody>
      </table>
      <FlowerModal show={showModal} handleClose={handleClose} flower={currentFlower}/>
    </div>
  )
}

export default RetailerFlowerList
