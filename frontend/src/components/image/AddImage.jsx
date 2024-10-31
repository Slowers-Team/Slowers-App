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
    const [uploadedImageName, setUploadedImageName] = useState("")
    const [uploadedImage, setUploadedImage] = useState()


    useEffect(() => {
      setID(entity._id)
    }, [entity])

    useEffect(() => {
      if (!uploadedImageName) {
        return
      } else {
        ImageService.get(uploadedImageName)        
        .then(data => setUploadedImage(data))
      }
    }, [uploadedImageName])

    const showForm = () => {
      setShow(true)
    }

    const hide = () => {
      setShow(false)
      setUploadedImageName("")
      setUploadedImage(null)
      setMessage("")
    }
  
    const createImage = imageObject => {
      setUploadedImageName("")
      setMessage("")

      ImageService.create({ ...imageObject, entity: id })
        .then(data => {
          console.info("Image upload succesful:", data)
          setMessage(t("alert.imageuploaded"))
          setUploadedImageName(data._id + "." + data.file_format)
        })
        .catch(error => {
          const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '')
          console.error("Image upload failed:", error)
          setMessage(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data))
        })
    }

    return (
      <>
        <Button variant="secondary" onClick={showForm}>{t("button.addimage")}</Button>
        <Modal size="l" show={show} onHide={hide}>
          <Modal.Header closeButton>
            <Modal.Title>{t("image.title")}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            { !uploadedImage
              ? <ImageForm createImage={createImage}/>
              : <img width={100} src={uploadedImage}/>
            }
          </Modal.Body>
          <Modal.Footer>
            { message }
          </Modal.Footer>
        </Modal>
      </>
    )
}

export default AddImage
