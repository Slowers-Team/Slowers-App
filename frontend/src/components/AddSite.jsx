import { Modal } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useState } from 'react'
import NewSiteForm from './NewSiteForm'

const AddSite = ({ createSite }) => {
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
        <button className='custom-button' onClick={handleShow}>
          + {t("button.addsite")}
        </button>
        <Modal size="l" show={showModal} onHide={handleClose}>
          <Modal.Header closeButton>
            <Modal.Title>{t("button.addsite")}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <NewSiteForm createSite={createSite} handleClose={handleClose}/>
          </Modal.Body>
        </Modal>
      </>
    )
  }

export default AddSite
