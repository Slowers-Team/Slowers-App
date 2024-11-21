import { Modal, Button, Tabs, Tab } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import AddImage from './image/AddImage'
import VisibilityButton from './VisibilityButton'

const FlowerModal = ({ show, handleClose, flower, deleteFlower, updateFlower }) => {
  const { t } = useTranslation()

  const handleFlowerDelete = (flower) => {
    if (deleteFlower) {
      deleteFlower(flower)
    }
    handleClose()
  }

  const addedTime = (flower) => {
    let addedTime = new Date(flower.added_time)

    let date = addedTime.toLocaleDateString('fi')
    let time = addedTime.toLocaleTimeString('fi', { hour: '2-digit', minute: '2-digit' })
    let addedTimeStr = `${date} ${time}`

    return addedTimeStr
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
          mountOnEnter={true}
          unmountOnExit={true}
          >
          <Tab eventKey="info" title={t('menu.info')}>
            <div>
              <h3>{t('menu.info')}</h3>
              <p>{t('flower.data.name')}: {flower.name}</p>
              <p>{t('flower.data.latinname')}: {flower.latin_name}</p>
              <p>{t('flower.data.addedtime')}: {addedTime(flower)}</p>
              <p>{t('flower.data.site')}: {flower.site_name}</p>
              <p>{t('flower.visible.long')}: {flower.visible 
                    ? t('flower.visible.true') 
                    : t('flower.visible.false')}
              <VisibilityButton flower={flower} updateFlower={updateFlower}/>
              </p>
              {deleteFlower && (
                <button id="deleteFlowerButton" onClick={() => handleFlowerDelete(flower)}>
                  {t('button.delete')}
                </button>
              )}
            </div>
          </Tab>
          <Tab eventKey="pictures" title={t('menu.pictures')}>
            <div>
              <h3>{t('menu.pictures')}</h3>
              <AddImage entity={flower}/>
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
