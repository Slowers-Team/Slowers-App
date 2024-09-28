import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'

import SiteService from '../services/sites'
import SiteFlexbox from '../components/SiteFlexbox'

const SingleSitePage = () => {
  let params = useParams()
  
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
      }
    };

    fetchSiteData()

  }, [params.id])

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

  const deleteSite = SiteObject => {
    if (window.confirm(`Are you sure you want to delete site ${SiteObject.name} and its subsites?`)) {
      SiteService
        .remove(SiteObject._id)
        .then(response => {
          console.log(response)
          setSite(l => l.filter(item => item._id !== SiteObject._id));
        })
    }
  }

  return (
    <div>
      <h2>{site.name}</h2>
      <SiteFlexbox createSite={createSite} sites={subsites}/>
    </div>
  )
}

export default SingleSitePage
