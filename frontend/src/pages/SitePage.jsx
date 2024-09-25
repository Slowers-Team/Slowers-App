import SiteService from '../services/sites'
import NewSiteForm from '../components/NewSiteForm'

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
    </div>
  )
}

export default SitePage
