import { useState, useEffect } from "react"
import { useParams, useNavigate } from "react-router-dom"

import SiteService from "../services/sites"
import flowerService from "../services/flowers"
import FlowerForm from "../components/FlowerForm"
import SiteFlexbox from "../components/SiteFlexbox"

import { useTranslation } from "react-i18next"

const SitePage = () => {
  const params = useParams()
  const navigate = useNavigate()
  const [site, setSite] = useState({})
  const [sites, setSites] = useState([])
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
        console.error("Error:", error)
        navigate("/")
      })
  }, [params.id, navigate])

  const addFlower = flowerObject => {
    flowerService.create(flowerObject).catch(error => {
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
        alert(t("error.error") + ": " + error.response.data)
      })
  }

  const deleteSite = siteObject => {
    if (
      window.confirm(`${t("label.confirmsitedeletion")} ${siteObject.name}?`)
    ) {
      const parentId = siteObject.parent ? siteObject.parent : ""
      SiteService.remove(siteObject._id)
        .then(() => navigate("/site/" + parentId))
        .catch(error => {
          console.error("Error deleting site:", error)
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
            <p>{site?.note}</p>
          </header>
          <div className="content">
            <aside className="side-container">
              <button
                id="showFlowerAddingFormButton"
                onClick={() => setShowAddNewFlower(!showAddNewFlower)}
              >
                {t("button.addflower")}
              </button>
              {showAddNewFlower && (
                <FlowerForm createFlower={addFlower} siteID={params.id} />
              )}
            </aside>
            <main className="main-container">
              <button onClick={handleBack}>{t("button.goback")}</button>
              <button id="deleteSiteButton" onClick={() => deleteSite(site)}>
                {t("button.deletethissite")}
              </button>
              <SiteFlexbox createSite={createSite} sites={sites} />
            </main>
          </div>
        </div>
      ) : (
        <>
          <header className="header">
            <h1>{t('title.rootsites')}</h1>
          </header>
          <div className="content">
            <main className="main-container">
              <SiteFlexbox createSite={createSite} sites={sites} />
            </main>
          </div>
        </>
      )}
    </>
  )
}

export default SitePage
