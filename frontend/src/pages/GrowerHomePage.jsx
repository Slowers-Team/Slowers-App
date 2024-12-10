import { useTranslation } from 'react-i18next'
import { useRouteLoaderData, useLoaderData, useParams } from 'react-router-dom' 
import { useState, useEffect } from 'react'; 
import ImageService from "../services/images";
import SiteImagesCarousel from "../components/image/SiteImagesCarousel";

const GrowerHomePage = () => {
  const { t, i18n } = useTranslation()
  const { site } = useLoaderData() ?? useRouteLoaderData("site") // very ugly hack that makes site "false"
  const params = useParams()
  const [images, setImages] = useState([]);

  useEffect(() => {
    if (site) {
      fetchImages();
    }
  }, []);
  
 const fetchImages = () => {
  ImageService.getImagesByEntity(params.siteId)
    .then(imageURLs => {
      console.log('Images after fetching:', imageURLs)
      setImages(imageURLs)
    })
    .catch(error => console.error('Error fetching images:', error))
  }

  return (
    <>
      {site ? (
        <h2>{site?.name} {t('title.sitehome')} </h2>
      ) : (
        <h2>{t("title.home")}</h2>
      )}
      {site?.note && (
        <p className="mx-1">
          {t("site.data.note")} : {site?.note}
        </p>
      )}
      {params.siteId && images && images.length > 0 ? (
        <div className="carousel-wrapper">
          <SiteImagesCarousel images={images} />
        </div>
      ) : null }
    </>
  );
};

export default GrowerHomePage;
