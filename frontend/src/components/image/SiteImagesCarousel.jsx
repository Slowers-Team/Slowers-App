import React, { useState } from "react"
import Carousel from 'react-bootstrap/Carousel'
import { CarouselCaption } from "react-bootstrap"
import { useTranslation } from "react-i18next" 
import '../../layouts/SiteImagesCarousel.css'

const SiteImagesCarousel = ({ images }) => {
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
      <Carousel activeIndex={activeIndex} onSelect={handleSelect} variant="dark">
        {images.map((image, index) => (
          <Carousel.Item key={image._id || index}>
            <img className="d-block w-100" src={image.url} alt={`Slide ${index + 1}`} />
          </Carousel.Item>
        ))}
      </Carousel>
    </div>
  )
}

export default SiteImagesCarousel
