import '../../layouts/Grower.css'
import ImageService from '../../services/images'
import '../image/FlowerListImage.css'
import FlowerModal from '../FlowerModal.jsx'
import { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'

const GrowerFlowerList = ({ flowers, deleteFlower, modifyFlower, setCheckedFlowers, updateFlower }) => {
  const { t, i18n } = useTranslation()
  const [showModal, setShowModal] = useState(false)
  const [currentFlower, setCurrentFlower] = useState("")
  const [checkedFlowers, setLocalCheckedFlowers] = useState([])
  const [images, setImages] = useState([])
  const [sortConfig, setSortConfig] = useState({ key: '', direction: '' })
  const [searchTerm, setSearchTerm] = useState("")

  useEffect(() => {
    setCheckedFlowers(checkedFlowers)
  }, [checkedFlowers, setCheckedFlowers])

  useEffect(() => {
    const newImages = Promise.all(flowers.map((f) => {
      if (f.favorite_image) {
        return ImageService.getByID(f.favorite_image)
          .then((url) => (
            {flower: f._id, url: url}
          ))
          // .catch((error) => console.error("bad", f, error))
          .catch((_) => {})
      }
    }))

    newImages.then((imgs) => setImages(imgs.filter((x)=>x)))
  
  }, [flowers])

  console.log(images)
  const handleShow = (flower) => {
    setShowModal(true)
    setCurrentFlower(flower)
  }
  
  const handleClose = () => {
    setShowModal(false)
    setCurrentFlower("")
  }

  const handleUpdate = (flowerObject) => {
    setCurrentFlower(flowerObject)
    updateFlower(flowerObject)
  }

  const areAllChecked = checkedFlowers.length === flowers.length

  const toggleCheckedAll = () => {
    if (areAllChecked) {
      setLocalCheckedFlowers([])
    } else {
      setLocalCheckedFlowers(flowers.map(flower => flower._id))
    }
  }

  const toggleCheckedFlower = (flower) => {
    setLocalCheckedFlowers((prevChecked) => (
      prevChecked.includes(flower) ? prevChecked.filter(id => id !== flower) : [...prevChecked, flower]
    ))
  }

  const handleSort = (key) => {
    let direction = 'asc'
    if (sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc'
    }
    setSortConfig({ key, direction })
  }

  const sortedFlowers = [...flowers].sort((a, b) => {
    if (!sortConfig.key) return 0

    const aValue = a[sortConfig.key]
    const bValue = b[sortConfig.key]

    if (aValue < bValue) return sortConfig.direction === 'asc' ? -1 : 1
    if (aValue > bValue) return sortConfig.direction === 'asc' ? 1 : -1
    return 0
  })

  const renderSortIcon = (key) => {
    if (sortConfig.key === key) {
      return sortConfig.direction === 'asc' ? ' \u25B2' : ' \u25BC'
    }
    return ' \u25BE'
  }

  const filteredFlowers = sortedFlowers.filter(flower => 
    flower.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    flower.latin_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    flower.site_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    flower.quantity.toString().includes(searchTerm.toLowerCase()) ||
    new Date(flower.added_time).toLocaleDateString('fi').toLowerCase().includes(searchTerm.toLowerCase()) ||
    new Date(flower.added_time).toLocaleString('fi', { hour: 'numeric', minute: '2-digit' }).toLowerCase().includes(searchTerm.toLowerCase()) ||
    (flower.visible ? t('flower.visible.true') : t('flower.visible.false')).toLowerCase().includes(searchTerm.toLowerCase())
  )

  return (
    <div className="growerFlowerList">
      <div className="d-flex justify-content-start mb-3 input-wrapper">
        <input
          type="text"
          placeholder={t('button.Search')}
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>
      <table id="growerFlowerList" className="table table-hover">
        <thead>
          <tr>
            <th>
              <input type="checkbox" onChange={toggleCheckedAll} checked={areAllChecked} />
            </th>
            <th/>
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
            <th onClick={() => handleSort('site_name')} style={{ cursor: 'pointer' }}>
              {t('flower.data.site')}
              {renderSortIcon('site_name')}
            </th>
            <th onClick={() => handleSort('quantity')} style={{ cursor: 'pointer' }}>
              {t('flower.data.qty')}
              {renderSortIcon('quantity')}
            </th>
            <th onClick={() => handleSort('visible')} style={{ cursor: 'pointer' }}>
              {t('flower.visible.short')}
              {renderSortIcon('visible')}
            </th>
            <th></th>
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
                <td>
                  <input type="checkbox" checked={checkedFlowers.includes(flower._id)} onChange={() => toggleCheckedFlower(flower._id)} />
                </td>
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
                <td>{addedTimeStr}</td>
                <td>{flower.site_name}</td>
                <td>{flower.quantity}</td>
                <td>{flower.visible 
                    ? t('flower.visible.true') 
                    : t('flower.visible.false')}</td>
                <td>
                  <button id='showFlowerPageButton' onClick={() => handleShow(flower)}>
                  {t('button.flowerpage')}
                  </button>
                </td>
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
      <FlowerModal show={showModal} handleClose={handleClose} flower={currentFlower} deleteFlower={deleteFlower} updateFlower={handleUpdate} modifyFlower={modifyFlower}/>
    </div>
  )
}

export default GrowerFlowerList
