import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import SiteService from '../services/sites'
import flowerService from '../services/flowers'
import FlowerForm from '../components/FlowerForm'
import SiteFlexbox from '../components/SiteFlexbox'

const SitePage = () => {
  const params = useParams()
  const navigate = useNavigate()
  const [site, setSite] = useState({})
  const [sites, setSites] = useState([])
  const [flowers, setFlowers] = useState()
  const [showAddNewFlower, setShowAddNewFlower] = useState(false)

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
    if (params.id) {
      flowerService.getFlowersForSite(params.id)
        .then(fetchedFlowers => {
          setFlowers(fetchedFlowers);
        })
        .catch(error => {
          console.error('Error fetching flowers:', error);
        });
    }
  }, [params.id])

  const addFlower = flowerObject => {
    flowerService.create(flowerObject).catch(error => {
      console.log(error)
      alert('Adding failed')
    })
  }

  const createSite = siteObject => {
    SiteService.create(siteObject)
      .then(newSite => {
        setSites(prevSites => (prevSites ? [...prevSites, newSite] : [newSite]))
      })
      .catch(error => {
        alert('Error: ' + error.response.data)
      })
  }

  const deleteSite = siteObject => {
    if (window.confirm(`Are you sure you want to delete site ${siteObject.name}?`)) {
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
            <p>{site?.note}</p>
          </header>
          <div className="content">
            <aside className="side-container">
              <button
                id="showFlowerAddingFormButton"
                onClick={() => setShowAddNewFlower(!showAddNewFlower)}
              >
                Add a new flower
              </button>
              {showAddNewFlower && <FlowerForm createFlower={addFlower} siteID={params.id} />}
            </aside>
            <main className="main-container">
              <button onClick={handleBack}>Go back</button>
              <button id="deleteSiteButton" onClick={() => deleteSite(site)}>
                Delete this site
              </button>
              <SiteFlexbox createSite={createSite} sites={sites} />
              <div>
                <h2>{t("label.flowers")}</h2>
                <ul>
                  {flowers.map(flower => (
                    <li key={flower._id}>
                      <h3>{flower.name}</h3>
                      <p>{flower.latin_name}</p>
                    </li>
                  ))}
                </ul>
              </div>
            </main>
          </div>
        </div>
      ) : (
        <>
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