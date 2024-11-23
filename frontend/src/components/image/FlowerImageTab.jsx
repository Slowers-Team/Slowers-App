import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import ImageService from '../../services/images'
import AddImage from './AddImage'
import FlowerImageGallery from './FlowerImageGallery'

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
      if (window.confirm(`${t("Confirm image deletion")}?`)) {
        ImageService.deleteImage(imageObject._id)
          .then(() => {
            setImages(l => l.filter(item => item._id !== imageObject._id))
            alert(t("Image deleted"))
          })
          .catch(error => {
            console.error('Error deleting image:', error)
            alert(t("Error"))
          })
      }
    }

    const onImageUpload = () => {
      fetchImages()
    }

    return (
      <div>
        {isGrower && <AddImage entity={flower} onImageUpload={onImageUpload}/>}
        <FlowerImageGallery images={images} deleteImage={deleteImage}/>
        </div>
    )
}

export default FlowerImageTab
