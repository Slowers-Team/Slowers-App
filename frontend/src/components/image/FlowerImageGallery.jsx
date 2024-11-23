import { useState } from "react"
import { useTranslation } from "react-i18next"
import { Button } from "react-bootstrap"
import CloseButton from 'react-bootstrap/CloseButton'
import 'bootstrap-icons/font/bootstrap-icons.css'
import '../../Misc.css'

const FlowerImageGallery = ({ images, deleteImage, favoriteImage }) => {
  const { t } = useTranslation() 
  const [activeIndex, setActiveIndex] = useState(0)
	const [selectedFavoriteIndex, setSelectedFavoriteIndex] = useState(0)

  if (activeIndex >= images.length && images.length > 0) {
    setActiveIndex(0)
  }

	const handleFavoriteSelect = (selectedIndex, imageObject) => {
		setSelectedFavoriteIndex(selectedIndex)
		favoriteImage(imageObject)
	}
	
  return (
    <div className="m-2">
			{(!images || images.length === 0) ? (
					<p>No flower images</p> 
			) : (
				<div>
					{images.map((image, index) => (
						<div class="image-box">
							<img className="d-block w-100" src={image.url} alt={`Slide ${index + 1}`} />
							<div className="image-buttons">
								<Button variant="dark" className="delete-button" onClick={() => deleteImage(image)}><i className="bi bi-trash"></i></Button>
								<Button variant="dark" onClick={() => handleFavoriteSelect(index, image)} className={`favourite-button ${selectedFavoriteIndex === index ? "selected" : ""}`} disabled={selectedFavoriteIndex !== null && selectedFavoriteIndex == index}>
                  <i className={`bi bi-star-fill ${selectedFavoriteIndex === index ? "text-warning" : ""}`}></i>
								</Button>
								</div>
						</div>
					))}
				</div>
			)}
    </div>
  )
}

export default FlowerImageGallery
