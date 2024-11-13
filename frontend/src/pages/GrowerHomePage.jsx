import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import siteService from '../services/sites';
import ImageService from '../services/images';
import SiteImagesCarousel from '../components/image/SiteImagesCarousel';
import AddImage from '../components/image/AddImage';

const GrowerHomePage = () => {
  const params = useParams();
  const [site, setSite] = useState(null);
  const [images, setImages] = useState([]);
  const { t } = useTranslation();

  useEffect(() => {
    if (params.siteId) {
      siteService.get(params.siteId).then(initialSite => setSite(initialSite.site))
    }
  }, []);

  useEffect(() => {
     fetchImages();
  }, [params.siteId]);

  const fetchImages = () => {
    ImageService.getImagesByEntity(params.siteId)
      .then(imageURLs => {
        console.log('Images after fetching:', imageURLs);
        setImages(imageURLs);
      })
      .catch(error => console.error('Error fetching images:', error));
  };

  const onImageUpload = () => {
    fetchImages();
  };
 
  const deleteImage = imageObject => {
    console.log("Deleting image:", imageObject); 
    if (!imageObject || !imageObject._id) {
      console.error("Image object is undefined or missing id");
      return;
    }
    if (window.confirm(`${t("Confirm image deletion")}?`)) {
      ImageService.deleteImage(imageObject._id)
        .then(() => {
          setImages(l => l.filter(item => item._id !== imageObject._id));
          alert(t("Image deleted"));
        })
        .catch(error => {
          console.error('Error deleting image:', error);
          alert(t("Error"));
        });
    }
  };
  
  return (
    <>
      {params.siteId && site ? (
        <main className="main-container">
          <div className="site-actions">
            <AddImage entity={site} onImageUpload={onImageUpload} />
          </div>
          <div className="uploaded-images">
            <SiteImagesCarousel images={images} onDelete={deleteImage} />
          </div>
        </main>
      ) : (
        <h2>{t('title.home')}</h2>
      )}
    </>
  );
};

export default GrowerHomePage;
