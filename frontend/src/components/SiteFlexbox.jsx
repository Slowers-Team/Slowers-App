import { useState } from 'react'
import { Link } from 'react-router-dom'

import NewSiteForm from './NewSiteForm'
import '../Misc.css'

const SiteFlexbox = ({ createSite, sites }) => {
  const [showAddNewSite, setShowAddNewSite] = useState(false)

  return (
    <div className={'flexbox'}>
      <div className={'flexGap'}>
        <div className={'box'}>
          Flowers
        </div>
          {sites && (
            sites.map(site => (
              <div className="box" key={site._id}>
                Name: <Link to={`/site/${site._id}`}>{site.name}</Link>
                <p>Note: {site.note}</p>
              </div>
            ))
          )}
        <div className={'box'}>
          <button id="addNewSiteButton" onClick={() => setShowAddNewSite(!showAddNewSite)}>Add new site</button>
          {showAddNewSite && <NewSiteForm createSite={createSite} />}
        </div>
      </div>
    </div>
    )
}

export default SiteFlexbox
