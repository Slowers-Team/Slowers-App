import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import SiteService from '../services/sites'
import flowerService from '../services/flowers'
import FlowerForm from '../components/FlowerForm'
import SiteFlexbox from '../components/SiteFlexbox'

import { useTranslation } from "react-i18next"

const SitePage = () => {
  const params = useParams()
  const navigate = useNavigate()
  const [site, setSite] = useState({})
  const [sites, setSites] = useState([])
  const [flowers, setFlowers] = useState() 
  const [showAddNewFlower, setShowAddNewFlower] = useState(false)
  const { t, i18n } = useTranslation()

  useEffect(() => {
    SiteService.get(params.id)
      .then(initialSites => {
        if (params.id) {
          setSite(initialSites.site)
          setSites(initialSites.subsites)
        } else {
          setSites(initialSites)
        }
      })
      .catch(error => {
        console.error('Error:', error)
        navigate('/')
      })
  }, [params.id, navigate])

  useEffect(() => {
        flowerService
          .getFlowersBySite(params.id)
          .then(flowers => {
            setFlowers(flowers);
          })
          .catch(error => {
            console.error("Error fetching flowers:", error);
          });
      }, [params.id, navigate]);

  const addFlower = flowerObject => {
    flowerService
      .create(flowerObject)
      .then(returnedFlower => 
        setFlowers(flowers ? flowers.concat(returnedFlower) :
          [returnedFlower])
      )
      .catch(error => {
      console.log(error)
      alert(t("error.addingfailed"))
    })
  }

  const createSite = siteObject => {
    SiteService.create(siteObject)
      .then(newSite => {
        setSites(prevSites => (prevSites ? [...prevSites, newSite] : [newSite]))
      })
      .catch(error => {
        const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '')
        alert(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data))
      })
  }

  const deleteSite = siteObject => {
    if (window.confirm(`${t("label.confirmsitedeletion")} ${siteObject.name}?`)) {
      const parentId = siteObject.parent ? siteObject.parent : ''
      SiteService.remove(siteObject._id)
        .then(() => navigate('/site/' + parentId))
        .catch(error => {
          console.error('Error deleting site:', error)
        })
    }
  }

  const handleBack = () => {
    navigate(-1)
  }

  return (
    <>
      {params.id ? (
        <div className="layout-container">
          <header className="header">
            <h1>{site?.name}</h1>
            <p className="site-note">{site?.note}</p>
          </header>
          <div className="content">
            <aside className="side-container">
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
                className="add-flower-button"
              >
                {t("button.addflower")}
              </button>
              {showAddNewFlower && <FlowerForm createFlower={addFlower} siteID={params.id} />}
            </aside>
            <main className="main-container">
              <div className="site-actions">
                <button onClick={handleBack} className="action-button">{t("button.goback")}</button>
                <button id="deleteSiteButton" onClick={() => deleteSite(site)} className="action-button">{t("button.deletethissite")}</button>
              </div>
              <SiteFlexbox createSite={createSite} sites={sites} />
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
  );
}

export default SitePage
