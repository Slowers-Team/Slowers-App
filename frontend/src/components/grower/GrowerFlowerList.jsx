import '../../layouts/Grower.css'
import FlowerModal from '../FlowerModal.jsx'
import { useState, useEffect } from "react"
import { useTranslation } from 'react-i18next'

import placeholderImage from '../../assets/images/site-placeholder-image.jpg'

const GrowerFlowerList = ({ flowers, deleteFlower, setCheckedFlowers, updateFlower}) => {
  const { t, i18n } = useTranslation()
  const [showModal, setShowModal] = useState(false)
  const [currentFlower, setCurrentFlower] = useState("")
  const [checkedFlowers, setLocalCheckedFlowers] = useState([])

  useEffect(() => {
    setCheckedFlowers(checkedFlowers)
  }, [checkedFlowers, setCheckedFlowers])

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

  return (
    <div className="growerFlowerList">
      <table id="growerFlowerList">
        <thead>
          <tr>
            <th>
              <input type="checkbox" onChange={toggleCheckedAll} checked={areAllChecked}/>
            </th>
            <th>Image</th>
            <th>{t('flower.data.name')}</th>
            <th>{t('flower.data.latinname')}</th>
            <th>{t('flower.data.addedtime')}</th>
            <th>{t('flower.data.site')}</th>
            <th>{t('flower.data.qty')}</th>
            <th>{t('flower.visible.short')}</th>
            <th></th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {flowers.map(flower => {
            let addedTime = new Date(flower.added_time)

            let date = addedTime.toLocaleDateString('fi')
            let hour = addedTime.toLocaleString('fi', { hour: 'numeric' })
            let minute = addedTime.toLocaleString('fi', { minute: '2-digit' })
            let addedTimeStr = `${date} ${hour}:${minute}`

            return (
              <tr key={flower._id}>
                <td>
                  <input type="checkbox" checked={checkedFlowers.includes(flower._id)} onChange={() => toggleCheckedFlower(flower._id)}/>
                </td>
                <td>
                  <img src={placeholderImage} className='img-fluid' alt={flower.name} />
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
      <FlowerModal show={showModal} handleClose={handleClose} flower={currentFlower} deleteFlower={deleteFlower} updateFlower={handleUpdate}/>
    </div>
  )
}

export default GrowerFlowerList
