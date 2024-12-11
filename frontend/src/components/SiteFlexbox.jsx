import { useState } from 'react'
import { Link } from 'react-router-dom'

import NewSiteForm from './NewSiteForm'
import '../Misc.css'
import placeholderImage from '../assets/images/site-placeholder-image.jpg'

import { useTranslation } from 'react-i18next'

const SiteFlexbox = ({ createSite, sites }) => {
  const [showAddNewSite, setShowAddNewSite] = useState(false)
  const { t, i18n } = useTranslation()

    
  return (
    <div className={'flexbox'}>
      <div className={'flexGap'}>
        {sites &&
          sites.map(site => (
            <div className="box" key={site._id}>
              <img 
                src={site.imageUrl ? `/api/images/${site.imageUrl}` : placeholderImage} 
                alt={site.name} 
                className="site-image" 
              />
              <h3><Link to={`/grower/${site._id}`} className="link-success">{site.name}</Link></h3>
            </div>
          ))}
        <div className={'box'}>
          <button className="custom-button" id="addNewSiteButton" onClick={() => setShowAddNewSite(!showAddNewSite)}>
            + {t('button.addsite')}
          </button>
          {showAddNewSite && <NewSiteForm createSite={createSite} />}
        </div>
      </div>
    </div>
  )
}

export default SiteFlexbox
