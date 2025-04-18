import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import SiteService from '../services/sites';
import flowerService from '../services/flowers';
import ImageService from '../services/images';
import FlowerForm from '../components/FlowerForm';
import SiteFlexbox from '../components/SiteFlexbox';
import AddImage from '../components/image/AddImage';
import SiteImagesCarousel from '../components/image/SiteImagesCarousel';

import { useTranslation } from "react-i18next";

const SitePage = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [site, setSite] = useState({});
  const [sites, setSites] = useState([]);
  const [flowers, setFlowers] = useState(); 
  const [showAddNewFlower, setShowAddNewFlower] = useState(false);
  const [images, setImages] = useState([]);
  const { t, i18n } = useTranslation();

  useEffect(() => {
    SiteService.get(params.id)
      .then(initialSites => {
        if (params.id) {
          setSite(initialSites.site);
          setSites(initialSites.subsites);
        } else {
          setSites(initialSites);
        }
      })
      .catch(error => {
        console.error('Error:', error);
        navigate('/');
      });
  }, [params.id, navigate]);

  useEffect(() => {
    flowerService.getFlowersBySite(params.id)
      .then(flowers => {
        setFlowers(flowers);
      })
      .catch(error => {
        console.error("Error fetching flowers:", error);
      });
  }, [params.id, navigate]);

  useEffect(() => {
    if (site._id) {
      fetchImages();
    }
  }, [site]);

  const addFlower = flowerObject => {
    flowerService.create(flowerObject)
      .then(returnedFlower => 
        setFlowers(flowers ? flowers.concat(returnedFlower) : [returnedFlower])
      )
      .catch(error => {
        console.log(error);
        alert(t("error.addingfailed"));
      });
  };

  const fetchImages = () => {
    ImageService.getImagesByEntity(site._id)
      .then(imageURLs => {
        console.log("Images after fetching:", imageURLs); 
        setImages(imageURLs);
      })
      .catch(error => console.error("Error fetching images:", error));
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
  
  const createSite = siteObject => {
    SiteService.create(siteObject)
      .then(newSite => {
        setSites(prevSites => (prevSites ? [...prevSites, newSite] : [newSite]));
      })
      .catch(error => {
        const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '');
        alert(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data));
      });
  };

  const deleteSite = siteObject => {
    if (window.confirm(`${t("label.confirmsitedeletion")} ${siteObject.name}?`)) {
      const parentId = siteObject.parent ? siteObject.parent : '';
      SiteService.remove(siteObject._id)
        .then(() => navigate('/site/' + parentId))
        .catch(error => {
          console.error('Error deleting site:', error);
        });
    }
  };

  const handleBack = () => {
    navigate(-1);
  };

  return (
    <>
      {params.id ? (
        <div className="layout-container">
          <header>
            <h1>{site?.name}</h1>
            <p className="site-note">{site?.note}</p>
          </header>
          <div>
            <aside>
              <h3>{t("site.siteflowers")}:</h3>
              <div className="flower-list">
                {Array.isArray(flowers) && flowers.length > 0 ? (
                  flowers.map(flower => (
                    <div key={flower._id} className="flower-card">
                      <h4>{flower.name}</h4>
                      <p>{flower.latinName}</p>
                    </div>
                  ))
                ) : (
                  <p>No flowers found for this site.</p>
                )}
              </div>
              <button
                id="showFlowerAddingFormButton"
                onClick={() => setShowAddNewFlower(!showAddNewFlower)}
                className="custom-button"
              >
                {t("button.addflower")}
              </button>
              {showAddNewFlower && <FlowerForm createFlower={addFlower} siteID={params.id} />}
            </aside>
            <main className="main-container">
              <div className="site-actions">
                <button onClick={handleBack} style={{ marginRight: "0.5rem" }} className="custom-button">{t("button.goback")}</button>
                <button id="deleteSiteButton" onClick={() => deleteSite(site)} className="custom-delete-button">{t("button.deletethissite")}</button>
                <AddImage entity={site} onImageUpload={fetchImages}/>
              </div>
              <SiteFlexbox createSite={createSite} sites={sites} />
              <div className="uploaded-images">
                <SiteImagesCarousel images={images} onDelete={deleteImage} /> 
              </div>
            </main>
          </div>
        </div>
      ) : (
        <div className="content">
          <main className="main-container">
            <SiteFlexbox createSite={createSite} sites={sites} />
          </main>
        </div>
      )}
    </>
  )
}


export default SitePage;
