import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import ImageService from '../../services/images'
import FlowerService from '../../services/flowers'
import AddImage from './AddImage'
import ImageGallery from './ImageGallery'

const FlowerImageTab = ({ isGrower, flower, updateFlower }) => {
    const { t } = useTranslation()
    const [images, setImages] = useState([])

    useEffect(() => {
      if (flower?._id) {
        fetchImages();
      }
    }, [flower]);
    

    const markFavorite = (images, id = null) => {
      if (id) {
        updateFlower({...flower, favorite_image: id})
      }

      const fav = id ?? flower?.favorite_image
      setImages(images.map(
        (img) => fav === img._id 
          ? {...img, favorite: true}
          : {...img, favorite: false}
      ))
    }

    const fetchImages = () => {
        ImageService.getImagesByEntity(flower._id)
          .then(fetchedImages => {
            console.log('Images after fetching:', fetchedImages)
            markFavorite(fetchedImages)
          })
          .catch(error => console.error('Error fetching images:', error))
    }

    const deleteImage = imageObject => {
      console.log("Deleting image:", imageObject) 
      if (!imageObject || !imageObject._id) {
        console.error("Image object is undefined or missing id")
        return
      }
      if (images.length === 1) {
        let confirmMessage = `${t('image.confirmimagedeletion')}?`
        if (flower.visible) {
            confirmMessage = `${t('image.flowerwillbehidden')}: ` + confirmMessage
        }
        if (window.confirm(confirmMessage)) {
          if (flower.visible) {
          FlowerService.toggleVisibility(flower._id).then(() => {updateFlower({...flower, visible: false})})
          }
          ImageService.deleteImage(imageObject._id).then(() => {setImages([])})
          .catch(error => {
            console.error('Error deleting image:', error)
            alert(t('error.erroroccured'))
          })
          ImageService.clearFavorite(flower._id, "flower")
          .then(_ => {
            console.log("cleared")
            updateFlower({...flower, favorite_image: null})
          })
          .catch(error => {
            console.error('Error clearing favorite image:', error)
            alert(t('error.erroroccured'))
          })

        }
        return
      }
      if (window.confirm(`${t('image.confirmimagedeletion')}?`)) {
        ImageService.deleteImage(imageObject._id)
        .then(() => {
          const updatedImages = images.filter(item => item._id !== imageObject._id)
          setImages(updatedImages)
          if (imageObject._id === flower.favorite_image) {
              const newFavoriteImage = updatedImages[0]?._id || null
              favoriteImage(newFavoriteImage)
          } else {
              favoriteImage(flower.favorite_image)
          }
        })
          .catch(error => {
            console.error('Error deleting image:', error)
            alert(t('error.erroroccured'))
          })
      }
    }

    const onImageUpload = (newImageID) => {
      if (images.length === 0) {
        favoriteImage(newImageID)
      }

      fetchImages()
    }

    const favoriteImage = imageID => {
      console.log("Favorite image:", imageID) 
      if (!imageID) {
        console.error("Image object is undefined or missing id")
        alert(t('error.erroroccured'))
        return
      }
      const response = ImageService.setFavorite(flower._id, "flower", imageID)
      markFavorite(images, imageID)

      console.log(response)
    }

    return (
      <div>
        {isGrower && <AddImage entity={flower} onImageUpload={onImageUpload}/>}
        <ImageGallery isGrower={isGrower} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage} type="flower"/>
      </div>
    )
}

export default FlowerImageTab
