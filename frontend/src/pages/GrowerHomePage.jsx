import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import { useTranslation } from "react-i18next";
import siteService from "../services/sites";
import ImageService from "../services/images";
import SiteImagesCarousel from "../components/image/SiteImagesCarousel";

const GrowerHomePage = () => {
  const params = useParams();
  const [site, setSite] = useState();
  const [images, setImages] = useState([]);
  const { t, i18n } = useTranslation();

  useEffect(() => {
    if (params.siteId) {
      siteService
        .get(params.siteId)
        .then((initialSite) => setSite(initialSite.site));
    }
  }, []);

  useEffect(() => {
    fetchImages()
 }, [params.siteId]);

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
      {params.siteId ? (
        <h2>
          {site?.name} {t("title.sitehome")}{" "}
        </h2>
      ) : (
        <h2>{t("title.home")}</h2>
      )}
      {site?.note && (
        <p className="mx-1">
          {t("site.data.note")} : {site?.note}
        </p>
      )}
      {params.siteId && (
        <div className="carousel-wrapper">
          <SiteImagesCarousel images={images} />
        </div>
      )}
    </>
  );
};

export default GrowerHomePage;
