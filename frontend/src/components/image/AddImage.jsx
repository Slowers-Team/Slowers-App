import { Modal, Button } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useState, useEffect } from 'react'
import ImageService from '../../services/images'
import ImageForm from './ImageForm'

const AddImage = ({ entity }) => {
    const [show, setShow] = useState(false)
    const [id, setID] = useState("")
    const [message, setMessage] = useState("")
    const { t, i18n } = useTranslation()

    useEffect(() => {
      setID(entity._id)
    }, [entity])


    const showForm = () => {
      setShow(true)
    }

    const hide = () => {
      setShow(false)
      setMessage("")
    }
  
    const createImage = imageObject => {
      setMessage("")

      ImageService.create({ ...imageObject, entity: id })
        .then(data => {
          console.info("Image upload succesful:", data)
          setMessage(t("alert.imageuploaded"))
        })
        .catch(error => {
          const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '')
          console.error("Image upload failed:", error)
          setMessage(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data))
        })
    }

    return (
      <>
        <Button className='mx-2' variant="light" onClick={showForm}>{t("button.addimage")}</Button>
        <Modal size="l" show={show} onHide={hide}>
          <Modal.Header closeButton>
            <Modal.Title>{t("image.title")}</Modal.Title>
          </Modal.Header>
          <Modal.Body>

        <ImageForm createImage={createImage}/>
          </Modal.Body>
          { message &&
            <Modal.Footer>
              { message }
            </Modal.Footer>
          }
        </Modal>
      </>
    )
}

export default AddImage
