import { useEffect, useState } from 'react'
import ImageService from '../../services/images'
import AddImage from './AddImage'
import FlowerImageGallery from './FlowerImageGallery'

const FlowerImageTab = ({ isGrower, flower }) => {
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
        console.log("Deleting image:")
    }
    return (
      <div>
        {isGrower && <AddImage entity={flower}/>}
        <FlowerImageGallery images={images} deleteImage={deleteImage}/>
        </div>
    )
}

export default FlowerImageTab
