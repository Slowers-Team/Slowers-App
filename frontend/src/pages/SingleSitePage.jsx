import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'

import SiteService from '../services/sites'
import NewSiteForm from '../components/NewSiteForm'
import SiteFlexbox from '../components/SiteFlexbox'

const SingleSitePage = () => {
  let params = useParams()
  
  const [site, setSite] = useState({})
  const [showAddNewSite, setShowAddNewSite] = useState(false)

  useEffect(() => {
    SiteService
      .get(params.id)
      .then(initialSite => setSite(initialSite))
  }, [])

  console.log(site)

  const createSite = SiteObject => {
      SiteService
          .create(SiteObject)
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
      <SiteFlexbox createSite={createSite}/>
    </div>
  )
}

export default SingleSitePage
