import { useState, useEffect } from 'react'

import SiteService from '../services/sites'
import NewSiteForm from '../components/NewSiteForm'
import SiteFlexbox from '../components/SiteFlexbox'

const SitePage = () => {
    const [sites, setSites] = useState([])
    const [showAddNewSite, setShowAddNewSite] = useState(false)
  
    useEffect(() => {
      SiteService
        .getRoot()
        .then(initialSites => setSites(initialSites))
    }, [])


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
            setSites(l => l.filter(item => item._id !== SiteObject._id));
          })
      }
    }


  return (
    <div>
      <NewSiteForm createSite={createSite} />
      <SiteFlexbox />
    </div>
  )
}

export default SitePage
