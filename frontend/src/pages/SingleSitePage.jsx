import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'

import SiteService from '../services/sites'
import SiteFlexbox from '../components/SiteFlexbox'

const SingleSitePage = () => {
  let params = useParams()
  const navigate = useNavigate()
  const [site, setSite] = useState({})
  const [subsites, setSubsites] = useState([])

  useEffect(() => {
    const fetchSiteData = async () => {
      try {
        const initialSite = await SiteService
        .get(params.id)
        .then(initialSite => {
          setSite(initialSite.site)
          setSubsites(initialSite.subsites)
        })
      } catch (error) {
        console.error('Error fetching site:', error)
        navigate('/')
      }
    }

    fetchSiteData()

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

  const deleteSite = async (SiteObject) => {
    if (window.confirm(`Are you sure you want to delete site ${SiteObject.name} and its subsites?`)) {
      try {
        const parentId = SiteObject.parent ? SiteObject.parent : ''
        await SiteService.remove(SiteObject._id)
        navigate('/site/' + parentId)
        
      } catch (error) {
        console.error('Error deleting site:', error)
      }
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
