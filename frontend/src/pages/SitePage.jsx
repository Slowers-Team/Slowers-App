import SiteService from '../services/sites'
import NewSiteForm from '../components/NewSiteForm'
import SiteFlexbox from '../components/SiteFlexbox'

const SitePage = () => {

    const createNewSite = SiteObject => {
        SiteService
            .create(SiteObject)
            .catch(error => {
                alert('Error: ' + error.response.data)
            })
    }

  return (
    <div>
      <NewSiteForm createNewSite={createNewSite} />
      <SiteFlexbox />
    </div>
  )
}

export default SitePage
