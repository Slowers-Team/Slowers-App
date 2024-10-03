
import { useState, useEffect } from "react"
import { useParams, useNavigate } from "react-router-dom"

import SiteService from "../services/sites"
import SiteFlexbox from "../components/SiteFlexbox"

const SitePage = () => {
  const params = useParams()
  const navigate = useNavigate()
  const [site, setSite] = useState({})
  const [sites, setSites] = useState([])

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

  const createSite = siteObject => {
    SiteService.create(siteObject)
      .then(newSite => {
        setSites(prevSites => (prevSites ? [...prevSites, newSite] : [newSite]))
      })
      .catch(error => {
        alert("Error: " + error.response.data)
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
      {params.id ? (
        <>
          <h2>{site?.name}</h2>
          <button id="deleteSiteButton" onClick={() => deleteSite(site)}>
            Delete this site
          </button>
          <p>{site?.note}</p>
        </>
      ) : (
        <>
          <h2>Root Sites</h2>
          <p>
            <br />
          </p>
        </>
      )}
      <SiteFlexbox createSite={createSite} sites={sites} />
    </div>
  )
}

export default SitePage
