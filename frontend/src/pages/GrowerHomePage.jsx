import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'
import SiteService from "../services/sites";
import ImageService from "../services/images";
import SiteImagesCarousel from "../components/image/SiteImagesCarousel";
import SiteFlexbox from "../components/SiteFlexbox";
import SiteMasonry from "../components/SiteMasonry";
import AddSite from "../components/AddSite";

const GrowerHomePage = () => {
  const params = useParams();
  const [site, setSite] = useState();
  const [sites, setSites] = useState([]);
  const [images, setImages] = useState([]);
  const { t, i18n } = useTranslation();

  useEffect(() => {
    SiteService.get(params.siteId)
      .then((initialSites) => {
        if (params.siteId) {
          setSite(initialSites.site);
          setSites(initialSites.subsites);
        } else {
          setSites(initialSites);
        }
      })
      fetchImages();
  }, [params.siteId]);

  const createSite = (siteObject) => {
    SiteService.create(siteObject)
      .then((newSite) => {
        setSites((prevSites) =>
          prevSites ? [...prevSites, newSite] : [newSite],
        );
      })
      .catch((error) => {
        const key =
          "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, "");
        alert(
          t("error.error") +
            ": " +
            (i18n.exists(key) ? t(key) : error.response.data),
        );
      });
  };

  const deleteSite = (siteObject) => {
    if (
      window.confirm(`${t("label.confirmsitedeletion")} ${siteObject.name}?`)
    ) {
      const parentId = siteObject.parent ? siteObject.parent : "";
      SiteService.remove(siteObject._id)
        .then(() => {
          if (parentId !== null && parentId !== "") {
            navigate("/grower/" + parentId + "/sites");
          } else {
            navigate("/grower/sites");
          }
        })
        .catch((error) => {
          console.error("Error deleting site:", error);
        });
    }
  };
 const fetchImages = () => {
  ImageService.getImagesByEntity(params.siteId)
    .then(imageURLs => {
      console.log('Images after fetching:', imageURLs)
      setImages(imageURLs)
    })
    .catch(error => console.error('Error fetching images:', error))
  }

  return (
    <Container>
      {params.siteId ? (
        <h2>
          {site?.name} {t("title.sitehome")}{" "}
        </h2>
      ) : (
        <h2>{t("title.home")}</h2>
      )}
      {site?.note && (
        <p className="mx-1">
          {site?.note}
        </p>
      )}
      <div className="d-flex gap-2">
        <AddSite createSite={createSite} />
        {params.siteId && (
        <button
          id="deleteSiteButton"
          onClick={() => deleteSite(site)}
          className="custom-button"
        >
          <i class="bi bi-trash3-fill"> </i>
          {t("button.deletethissite")}
        </button>
        )}
      </div>
      <div>
        <h3 className="my-3">{t("title.sites")}</h3>
        <SiteMasonry sites={sites}/>
      </div>
      {/*
      {params.siteId && images && images.length > 0 ? (
        <div className="info-container">
          <h3 className="my-3">Site's images</h3>
          {site?.note && (
            <p className="mx-1">
              {t("site.data.note")} : {site?.note}
            </p>
          )}
          <div className="carousel-wrapper">
            <SiteImagesCarousel images={images} />
          </div>
        </div>
      ) : null }
      */}
    </Container>
  );
};

export default GrowerHomePage;
