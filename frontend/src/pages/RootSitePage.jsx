import { useState, useEffect } from 'react'

import SiteService from '../services/sites'
import SiteFlexbox from '../components/SiteFlexbox'

const RootSitePage = () => {
  const [rootSites, setRootSites] = useState([])

  useEffect(() => {
    SiteService
      .getRoot()
      .then(initialSites => setRootSites(initialSites))
  }, [])
  
  const createSite = SiteObject => {
      SiteService
        .create(SiteObject)
        .then(newSite => {
          setRootSites(prevSites => prevSites ? [...prevSites, newSite] : [newSite])
        })
        .catch(error => {
          alert('Error: ' + error.response.data)
        })
  }

  return (
    <div>
      <h2>Root Sites</h2>
      <SiteFlexbox createSite={createSite} sites={rootSites}/>
    </div>
  )
}
export default RootSitePage
