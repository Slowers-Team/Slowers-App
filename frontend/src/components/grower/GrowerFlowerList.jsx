import '../../layouts/Grower.css'
import FlowerModal from '../FlowerModal.jsx'
import { useState, useEffect } from "react"
import { useTranslation } from 'react-i18next'

const GrowerFlowerList = ({ flowers, deleteFlower, setCheckedFlowers}) => {
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

  const areAllSelected = checkedFlowers.length === flowers.length

  const toggleSelectAll = () => {
    if (areAllSelected) {
      setLocalCheckedFlowers([])
    } else {
      setLocalCheckedFlowers(flowers.map(flower => flower._id))
    }
  }

  const toggleSelectFlower = (flower) => {
    setLocalCheckedFlowers((prevSelected) =>
      prevSelected.includes(flower)
        ? prevSelected.filter((id) => id !== flower)
        : [...prevSelected, flower]
    )
  }

  return (
    <div className="growerFlowerList">
      <table id="growerFlowerList">
        <thead>
          <tr>
            <th>
              <input type="checkbox" onChange={toggleSelectAll} checked={areAllSelected}/>
            </th>
            <th>{t('flower.data.name')}</th>
            <th>{t('flower.data.latinname')}</th>
            <th>{t('flower.data.addedtime')}</th>
            <th>{t('flower.data.site')}</th>
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
                  <input type="checkbox" checked={checkedFlowers.includes(flower._id)} onChange={() => toggleSelectFlower(flower._id)}/>
                </td>
                <td>{flower.name}</td>
                <td>
                  <em>{flower.latin_name}</em>
                </td>
                <td>{addedTimeStr}</td>
                <td>{flower.site_name}</td>
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
      <FlowerModal show={showModal} handleClose={handleClose} flower={currentFlower} deleteFlower={deleteFlower}/>
    </div>
  )
}

export default GrowerFlowerList
