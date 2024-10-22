import '../../layouts/Grower.css'
import FlowerModal from '../FlowerModal.jsx'
import { useState } from "react"
import { Button, Table } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'

const GrowerFlowerList = ({ flowers, deleteFlower }) => {
  const { t, i18n } = useTranslation()
  const [showModal, setShowModal] = useState("")

  const handleShow = () => setShowModal(true)
  const handleClose = () => setShowModal(false)

  return (
    <div className="growerFlowerList">
      <table id="growerFlowerList">
        <thead>
          <tr>
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
                <td>{flower.name}</td>
                <td>
                  <em>{flower.latin_name}</em>
                </td>
                <td>{addedTimeStr}</td>
                <td>{flower.site_name}</td>
                <td>
                  <button type="button" className="btn btn-light" onClick={handleShow}>
                  Demo modal
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
      <FlowerModal show={showModal} handleClose={handleClose} />
    </div>
  )
}

export default GrowerFlowerList
