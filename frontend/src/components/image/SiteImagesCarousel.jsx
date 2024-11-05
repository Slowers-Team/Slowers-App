import React from "react";
import Carousel from 'react-bootstrap/Carousel';
import 'bootstrap/dist/css/bootstrap.min.css';
import'../../layouts/SiteImagesCarousel.css';
import { CarouselCaption } from "react-bootstrap";

const SiteImagesCarousel = ({ images, onDelete }) => {
  if (!images || images.length === 0) {
    return <p>No images uploaded for this site.</p>;
  }

  return (
    <Carousel>
      {images.map((image, index) => (
        <Carousel.Item key={index}>
          <img className="d-block w-100" src={image.url} alt={`Slide ${index + 1}`} />
          <CarouselCaption>
            <button
              onClick={() => {
                console.log("Delete button clicked for image ID:", image.id);
                onDelete(image._id)}}
              className="btn"
            >
              Delete
            </button>
          </CarouselCaption>
        </Carousel.Item>
      ))}
    </Carousel>
  );
};

export default SiteImagesCarousel;
