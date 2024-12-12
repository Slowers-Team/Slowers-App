import { useParams } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Container } from 'react-bootstrap'
import siteService from '../services/sites'
import ImageService from '../services/images'
import ImageGallery from '../components/image/ImageGallery'
import AddImage from '../components/image/AddImage'

const GrowerImagesPage = () => {
  const params = useParams()
  const [site, setSite] = useState(null)
  const [images, setImages] = useState([])
  const { t } = useTranslation()

  useEffect(() => {
    if (params.siteId) {
      siteService.get(params.siteId).then(initialSite => setSite(initialSite.site))
    }
  }, [])

  useEffect(() => {
    if (params.siteId && site) {
     fetchImages()
    }
  }, [params.siteId, site])

  const fetchImages = () => {
    ImageService.getImagesByEntity(params.siteId)
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
      if (window.confirm(`${t('image.confirmimagedeletion')}?`)) {
        ImageService.deleteImage(imageObject._id).then(() => {setImages([])})
        .catch(error => {
          console.error('Error deleting image:', error)
          alert(t('error.erroroccured'))
        })
        ImageService.clearFavorite(site._id, "site")
        .then(_ => {
          console.log("cleared")
          updateSite({...site, favorite_image: null})
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
        if (imageObject._id === site.favorite_image) {
            const newFavoriteImage = updatedImages[0]?._id || null
            favoriteImage(newFavoriteImage)
        } else {
            favoriteImage(site.favorite_image)
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

  const markFavorite = (images, id = null) => {
    if (id) {
      updateSite({...site, favorite_image: id})
    }

    const fav = id ?? site?.favorite_image
    setImages(images.map(
      (img) => fav === img._id 
        ? {...img, favorite: true}
        : {...img, favorite: false}
    ))
  }

  const favoriteImage = imageID => {
    console.log("Favorite image:", imageID) 
    if (!imageID) {
      console.error("Image object is undefined or missing id")
      alert(t('error.erroroccured'))
      return
    }
    const response = ImageService.setFavorite(site._id, "site", imageID)
    markFavorite(images, imageID)

    console.log(response)
  }

  const updateSite = SiteObject => {
    setSite(SiteObject)
  }
  
  return (
    <Container>
    {site && (
      <div>
        <h2>{site?.name} {t('title.siteimages')}</h2>
        <AddImage entity={site} onImageUpload={onImageUpload} />
        <ImageGallery isGrower={true} images={images} deleteImage={deleteImage} favoriteImage={favoriteImage}/>
      </div>
    )}
    </Container>
  )  
}

export default GrowerImagesPage

