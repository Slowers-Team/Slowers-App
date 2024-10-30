import { Modal, Button } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useState, useEffect } from 'react'
import ImageService from '../../services/images'
import ImageForm from './ImageForm'

const AddImage = ({ entity }) => {
    const [show, setShow] = useState(false)
    const [id, setID] = useState("")
    const { t, i18n } = useTranslation()

    useEffect(() => {
      setID(entity._id)
    }, [entity])

    const toggleVisibility = () => {
      setShow(!show)
    }
  
    const createImage = imageObject => {
      ImageService.create({ ...imageObject, entity: id })
        .then(_ => {
          alert(t("alert.imageuploaded"))
          toggleVisibility()
        })
        .catch(error => {
          const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '')
          alert(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data))
        })
    }

    return (
      <>
        <Button variant="secondary" onClick={toggleVisibility}>{t("button.addimage")}</Button>
        <Modal size="l" show={show} onHide={toggleVisibility}>
          <Modal.Header closeButton>
            <Modal.Title>{t("image.title")}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <ImageForm createImage={createImage}/>
          </Modal.Body>
        </Modal>
      </>
    )
}

export default AddImage
