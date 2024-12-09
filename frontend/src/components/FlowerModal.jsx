import { Modal, Button, Tabs, Tab } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import FlowerInfoTab from './FlowerInfoTab'
import FlowerImageTab from './image/FlowerImageTab'


const FlowerModal = ({ show, handleClose, flower, deleteFlower, updateFlower, modifyFlower }) => {
  const { t } = useTranslation()

  const isGrower = Boolean(deleteFlower && updateFlower && modifyFlower)

  return (
    <Modal size="lg" show={show} onHide={handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>{flower.name}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Tabs defaultActiveKey="info" className="mb-3" mountOnEnter={true} unmountOnExit={true}>
          <Tab eventKey="info" title={t('menu.info')}>
            <div>
              <FlowerInfoTab isGrower={isGrower} flower={flower} deleteFlower={deleteFlower} updateFlower={updateFlower} modifyFlower={modifyFlower} handleClose={handleClose}/>
            </div>
          </Tab>
          <Tab eventKey="images" title={t('menu.images')}>
            <div>
              <FlowerImageTab isGrower={isGrower} flower={flower} updateFlower={updateFlower}/>
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
