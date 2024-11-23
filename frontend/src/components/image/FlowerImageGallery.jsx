import React, { useState } from "react"
import { Button } from 'react-bootstrap'
import { useTranslation } from "react-i18next" 

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
				<div class="row">
					{images.map((image, index) => (
						<div class="col-lg-4 col-md-12 mb-4 mb-lg-0">
							<img className="d-block w-100" src={image.url} alt={`Slide ${index + 1}`} />
						</div>
					))}
				</div>
			)}
    </div>
  )
}

export default FlowerImageGallery
