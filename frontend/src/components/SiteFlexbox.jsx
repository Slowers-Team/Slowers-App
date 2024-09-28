import { useState } from 'react'

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
                Name: {site.name}
                <p>Note: {site.note}</p>
                <p>ID: {site._id}</p>
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
