import { useParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { useTranslation } from "react-i18next";
import SiteService from "../services/sites";
import SiteMasonry from "../components/SiteMasonry";
import AddSite from "../components/AddSite";

const GrowerHomePage = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [site, setSite] = useState();
  const [sites, setSites] = useState([]);
  const { t, i18n } = useTranslation();

  useEffect(() => {
    SiteService.get(params.siteId)
      .then((initialSites) => {
        if (params.siteId) {
          setSite(initialSites.site)
          setSites(initialSites.subsites)
        } else {
          setSites(initialSites)
        }
      })
  }, [params.siteId])

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
            navigate("/grower/" + parentId);
          } else {
            navigate("/grower");
          }
        })
        .catch((error) => {
          console.error("Error deleting site:", error);
        });
    }
  };

  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-12">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              {params.siteId ? (
                <h2>
                  {site?.name} {t("title.sitehome")}{" "}
                </h2>
              ) : (
                <h2>{t("title.growerhome")}</h2>
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
                  className="custom-delete-button"
                >
                  <i className="bi bi-trash3-fill"> </i>
                  {t("button.deletethissite")}
                </button>
                )}
              </div>
              <div>
                {sites && sites.length > 0 && (
                  <h3 className="my-3">{t("title.sites")}</h3>
                )}
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
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GrowerHomePage;
