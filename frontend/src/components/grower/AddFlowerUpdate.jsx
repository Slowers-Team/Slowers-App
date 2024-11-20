import { Modal, Button } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useState } from 'react'

const AddFlowerUpdate = ({ checkedFlowers }) => {
    const { t } = useTranslation()
    const [showModal, setShowModal] = useState(false)
  
    const handleShow = () => {
      console.log(checkedFlowers)
      setShowModal(true)
    }
    
    const handleClose = () => {
      setShowModal(false)
    }

    return (
      <>
        <Button variant="light" onClick={handleShow}>
          {t("button.update")}
        </Button>
        <Modal size="l" show={showModal} onHide={handleClose}>
          <Modal.Header closeButton>
            <Modal.Title>{t("button.update")}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
          </Modal.Body>
        </Modal>
      </>
    )
  }

export default AddFlowerUpdate
