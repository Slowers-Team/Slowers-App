import React from "react";
import Carousel from 'react-bootstrap/Carousel';
import { CarouselCaption } from "react-bootstrap";
import { useTranslation } from "react-i18next"; 
import '../../layouts/SiteImagesCarousel.css';

const SiteImagesCarousel = ({ images, onDelete }) => {
  const { t } = useTranslation(); 

  return (
    <div className="site-images-carousel">
      <h2>{t('carousel.siteImages')}</h2> 

      {(!images || images.length === 0) ? (
        <p>{t('carousel.noImages')}</p> 
      ) : (
        <Carousel>
          {images.map((image, index) => (
            <Carousel.Item key={index}>
              <img className="d-block w-100" src={image.url} alt={`Slide ${index + 1}`} />
              <CarouselCaption>
                <button
                  onClick={() => onDelete(image)}
                  className="btn delete-button"
                >
                  {t('button.delete')} 
                </button>
              </CarouselCaption>
            </Carousel.Item>
          ))}
        </Carousel>
      )}
    </div>
  );
};

export default SiteImagesCarousel;
