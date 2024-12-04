import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import ImageService from '../../services/images'
import AddImage from './AddImage'
import ImageGallery from './ImageGallery'

const FlowerImageTab = ({ isGrower, flower }) => {
    const { t } = useTranslation()
    const [images, setImages] = useState([])

    useEffect(() => {
        fetchImages()
    }, [])

    const fetchImages = () => {
        ImageService.getImagesByEntity(flower._id)
          .then(imageURLs => {
            console.log('Images after fetching:', imageURLs)
            setImages(imageURLs)
          })
          .catch(error => console.error('Error fetching images:', error))
    }

    const deleteImage = imageObject => {
      console.log("Deleting image:", imageObject) 
      if (!imageObject || !imageObject._id) {
        console.error("Image object is undefined or missing id")
        return
      }
      if (window.confirm(`${t('image.confirmimagedeletion')}?`)) {
        ImageService.deleteImage(imageObject._id)
          .then(() => {
            setImages(l => l.filter(item => item._id !== imageObject._id))
          })
          .catch(error => {
            console.error('Error deleting image:', error)
            alert(t('error.erroroccured'))
          })
      }
    }

    const onImageUpload = () => {
      fetchImages()
    }

    const favoriteImage = imageObject => {
      console.log("Favorite image:", imageObject) 
      if (!imageObject || !imageObject._id) {
        console.error("Image object is undefined or missing id")
        alert(t('error.erroroccured'))
        return
      }
    }

    return (
      <div>
        <h3>{t('menu.images')}</h3>
        {isGrower && <AddImage entity={flower} onImageUpload={onImageUpload}/>}
        <ImageGallery isGrower={isGrower} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} type="flower"/>
      </div>
    )
}

export default FlowerImageTab
