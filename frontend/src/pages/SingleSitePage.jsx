import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'

import SiteService from '../services/sites'
import SiteFlexbox from '../components/SiteFlexbox'

const SingleSitePage = () => {
  const params = useParams()
  const navigate = useNavigate()
  const [site, setSite] = useState({})
  const [subsites, setSubsites] = useState([])

  useEffect(() => {
    SiteService.get(params.id)
      .then(initialSite => {
        setSite(initialSite.site)
        setSubsites(initialSite.subsites)
      })
      .catch(error => {
        console.error("Error fetching site:", error)
        navigate("/")
      })
  }, [params.id, navigate])

  const createSite = SiteObject => {
      SiteService
          .create(SiteObject)
          .then(newSite => {
              setSubsites(prevSites => prevSites ? [...prevSites, newSite] : [newSite])
          })
          .catch(error => {
              alert('Error: ' + error.response.data)
          })
  }

  const deleteSite = siteObject => {
    if (
      window.confirm(`Are you sure you want to delete site ${siteObject.name}?`)
    ) {
      const parentId = siteObject.parent ? siteObject.parent : ""
      SiteService.remove(siteObject._id)
        .then(() => navigate("/site/" + parentId))
        .catch(error => {
          console.error("Error deleting site:", error)
        })
    }
  }

  return (
    <div>
      <h2>{site.name}</h2>
      <button id="deleteSiteButton" onClick={() => deleteSite(site)}>Delete this site</button>
      <p> {site.note} </p>
      <SiteFlexbox createSite={createSite} sites={subsites}/>
    </div>
  )
}

export default SingleSitePage
