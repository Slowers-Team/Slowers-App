import { Modal, Button, Tabs, Tab } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'

const FlowerModal = ({ show, handleClose, flower }) => {
  const { t } = useTranslation()

  return (
    <Modal size="xl" show={show} onHide={handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>{flower.name}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Tabs
          defaultActiveKey="info"
          id="uncontrolled-tab-example"
          className="mb-3"
          >
          <Tab eventKey="info" title="Info">
            <div>
              <p>{t('flower.data.name')}: {flower.name}</p>
              <p>{t('flower.data.latinname')}: {flower.latin_name}</p>
              <p>{t('flower.data.addedtime')}: {flower.added_time}</p>
              <p>{t('flower.data.site')}: {flower.site}</p>
            </div>
          </Tab>
          <Tab eventKey="pictures" title="Pictures">
            Tab content for Pictures
          </Tab>
        </Tabs>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          CLOSE
        </Button>
      </Modal.Footer>
    </Modal>
  )
}

export default FlowerModal
