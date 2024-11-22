import React, { useState } from "react"
import { Button } from 'react-bootstrap'
import { useTranslation } from "react-i18next" 

const FlowerImageGallery = ({ images, onDelete }) => {
  const { t } = useTranslation() 
  const [activeIndex, setActiveIndex] = useState(0)

  if (activeIndex >= images.length && images.length > 0) {
    setActiveIndex(0)
  }

  return (
    <div>
			{(!images || images.length === 0) ? (
					<p>{t('carousel.noImages')}</p> 
			) : (
				<div class="row">
					{images.map((image, index) => (
						<div class="col-lg-4 col-md-12 mb-4 mb-lg-0">
							<img className="d-block w-100" src={image.url} alt={`Slide ${index + 1}`} />
							<Button onClick={() => onDelete(image)}>
								{t('button.delete')}
							</Button>
						</div>
					))}
				</div>
			)}
    </div>
  )
}

export default FlowerImageGallery
