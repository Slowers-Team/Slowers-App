import { Modal, Button } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'

const FlowerModal = ({ show, handleClose }) => {
  const { t } = useTranslation()

  return (
    <Modal show={show} onHide={handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>TITLE</Modal.Title>
      </Modal.Header>
      <Modal.Body>TEXT</Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          CLOSE
        </Button>
      </Modal.Footer>
    </Modal>
  )
}

export default FlowerModal
