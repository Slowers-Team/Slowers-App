import { Modal, Button, Tabs, Tab } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'

const FlowerModal = ({ show, handleClose, flower, deleteFlower }) => {
  const { t } = useTranslation()

  const handleFlowerDelete = (flower) => {
    deleteFlower(flower)
    handleClose()
  }


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
          <Tab eventKey="info" title={t('menu.info')}>
            <div>
              <h3>{t('menu.info')}</h3>
              <p>{t('flower.data.name')}: {flower.name}</p>
              <p>{t('flower.data.latinname')}: {flower.latin_name}</p>
              <p>{t('flower.data.addedtime')}: {flower.added_time}</p>
              <p>{t('flower.data.site')}: {flower.site}</p>
              <button id="deleteFlowerButton" onClick={() => handleFlowerDelete(flower)}>
                {t('button.delete')}
              </button>
            </div>
          </Tab>
          <Tab eventKey="pictures" title={t('menu.pictures')}>
            <div>
              <h3>{t('menu.pictures')}</h3>
            </div>
          </Tab>
          <Tab eventKey="lifecycle" title={t('menu.lifecycle')}>
            <div>
              <h3>{t('menu.lifecycle')}</h3>
            </div>
          </Tab>
        </Tabs>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          {t('button.close')}
        </Button>
      </Modal.Footer>
    </Modal>
  )
}

export default FlowerModal
