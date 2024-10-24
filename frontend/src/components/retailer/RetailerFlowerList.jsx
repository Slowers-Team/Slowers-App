import '../../layouts/Retailer.css'
import { useTranslation } from 'react-i18next'
import FlowerModal from '../FlowerModal.jsx'
import { useState } from "react"
import { Button, Table } from 'react-bootstrap'

const RetailerFlowerList = ({ flowers }) => {
  const { t, i18n } = useTranslation()
  const [showModal, setShowModal] = useState("")
  const [currentFlower, setCurrentFlower] = useState("")

  const handleShow = (flower) => {
    setShowModal(true)
    setCurrentFlower(flower)
  }
  
  const handleClose = () => setShowModal(false)

  return (
    <div className="retailerFlowerList">
      <table id="retailerFlowerList">
        <thead>
          <tr>
            <th>{t('flower.data.name')}</th>
            <th>{t('flower.data.latinname')}</th>
            <th>{t('flower.data.addedtime')}</th>
            <th>{t('flower.data.grower')}</th>
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
                <td>{flower.name}</td>
                <td>
                  <em>{flower.latin_name}</em>
                </td>
                <td>{addedTimeStr}</td>
                <td>{flower.grower_email}</td>
                <td>
                  <button id='showFlowerPageButton' onClick={() => handleShow(flower)}>
                  {t('button.flowerpage')}
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
