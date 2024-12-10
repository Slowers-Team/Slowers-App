import { useTranslation } from 'react-i18next'
import { useRouteLoaderData, useLoaderData, useParams, defer, Await } from 'react-router-dom' 
import { useState, useEffect, Suspense } from 'react'; 
import ImageService from "../services/images";
import SiteImagesCarousel from "../components/image/SiteImagesCarousel";

const GrowerHomePage = () => {
  const { t, i18n } = useTranslation()
  const { site, images } = useLoaderData() ?? useRouteLoaderData("site")

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
      <Suspense
        fallback={<p>Loading images...</p>}
      >
        <Await
          resolve={images}
          errorElement={
            <p>Error loading images!</p>
          }
        >
          {(imgs) => (
             imgs && imgs.length > 0 ? (
              <div className="carousel-wrapper">
                <SiteImagesCarousel images={imgs} />
              </div>
            ) : null 
          )}
        </Await>
      </Suspense>
      
    </>
  );
};

export default GrowerHomePage;
