import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import SiteService from "../services/sites";
import SiteFlexbox from "../components/SiteFlexbox";
import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'

const GrowerSitesPage = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [site, setSite] = useState({});
  const [sites, setSites] = useState([]);
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
      .catch((error) => {
        console.error("Error:", error);
        navigate("/");
      });
  }, [params.siteId, navigate]);

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

  const handleBack = () => {
    navigate(-1);
  };

  return (
    <Container>
      {params.siteId ? (
        <div>
          <h2>
            {site?.name} {t("title.sitesites")}
          </h2>
          <div className="my-2">
            <button
              onClick={handleBack}
              style={{ marginRight: "0.5rem" }}
              className="custom-button"
            >
              {t("button.goback")}
            </button>
            <button
              id="deleteSiteButton"
              onClick={() => deleteSite(site)}
              className="custom-button"
            >
              <i className="bi bi-trash3-fill"> </i>
              {t("button.deletethissite")}
            </button>
          </div>
          <SiteFlexbox createSite={createSite} sites={sites}/>
        </div>
      ) : (
        <div>
          <h2 className="mb-3">{t("title.sites")}</h2>
          <SiteFlexbox createSite={createSite} sites={sites}/>
        </div>
      )}
    </Container>
  );
};

export default GrowerSitesPage;
