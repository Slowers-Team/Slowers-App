import { useState } from "react"
import { useTranslation } from "react-i18next" 
import CloseButton from 'react-bootstrap/CloseButton'
import '../../Misc.css'

const FlowerImageGallery = ({ images, deleteImage }) => {
  const { t } = useTranslation() 
  const [activeIndex, setActiveIndex] = useState(0)

  if (activeIndex >= images.length && images.length > 0) {
    setActiveIndex(0)
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
							<CloseButton className="delete-button" onClick={() => deleteImage(image)} />							
						</div>
					))}
				</div>
			)}
    </div>
  )
}

export default FlowerImageGallery
