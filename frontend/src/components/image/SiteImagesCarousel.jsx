import React, { useState } from "react"
import Carousel from 'react-bootstrap/Carousel'
import { CarouselCaption } from "react-bootstrap"
import { useTranslation } from "react-i18next" 
import '../../layouts/SiteImagesCarousel.css'

const SiteImagesCarousel = ({ images, onDelete }) => {
  const { t } = useTranslation() 
  const [activeIndex, setActiveIndex] = useState(0)

  if (activeIndex >= images.length && images.length > 0) {
    setActiveIndex(0)
  }

  const handleSelect = (selectedIndex) => {
    setActiveIndex(selectedIndex)
  }

  return (
    <div className="site-images-carousel">
      <h3>{t('carousel.siteImages')}</h3> 

      {(!images || images.length === 0) ? (
        <p>{t('carousel.noImages')}</p> 
      ) : (
        <Carousel activeIndex={activeIndex} onSelect={handleSelect}>
          {images.map((image, index) => (
            <Carousel.Item key={image._id || index}>
              <img className="d-block w-100" src={image.url} alt={`Slide ${index + 1}`} />
              <CarouselCaption>
                <button onClick={() => onDelete(image)} className="btn delete-button">
                  {t('button.delete')}
                </button>
              </CarouselCaption>
            </Carousel.Item>
          ))}
        </Carousel>
      )}
    </div>
  )
}

export default SiteImagesCarousel
