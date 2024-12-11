import { Modal, Button } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useState } from 'react'
import FlowerForm from '../FlowerForm'

const AddFlower = ({ createFlower, siteID }) => {
    const { t, i18n } = useTranslation()
    const [showModal, setShowModal] = useState(false)
  
    const handleShow = () => {
      setShowModal(true)
    }
    
    const handleClose = () => {
      setShowModal(false)
    }

    return (
      <>
        <Button variant="light" onClick={handleShow}>
          + {t("button.addflower")}
        </Button>
        <Modal size="l" show={showModal} onHide={handleClose}>
          <Modal.Header closeButton>
            <Modal.Title>{t("button.addflower")}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <FlowerForm createFlower={createFlower} siteID={siteID} handleClose={handleClose}/>
          </Modal.Body>
        </Modal>
      </>
    )
  }

export default AddFlower
