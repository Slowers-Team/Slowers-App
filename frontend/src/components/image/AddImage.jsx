import { Modal, Button, Tabs, Tab } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import ImageService from '../../services/images'
import ImageForm from './ImageForm'

const AddImage = ({ entityID, entityName }) => {
    const createImage = imageObject => {
      ImageService.create({ ...imageObject, entity: entityID })
        .then(data => console.log(data))
        .catch(error => {
          const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '')
          alert(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data))
        })
    }

    return (
      <>
        <ImageForm createImage={createImage}/>
      </>
    )
}

export default AddImage
