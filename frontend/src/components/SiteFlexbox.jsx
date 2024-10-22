import { useState } from 'react'
import { Link } from 'react-router-dom'

import NewSiteForm from './NewSiteForm'
import '../Misc.css'

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
              <h3><Link to={'/site/' + site._id} className='link-success'>{site.name}</Link></h3>
              <p>{t('site.data.note')}:<br />{site.note}</p>
            </div>
          ))}
        <div className={'box'}>
          <button id="addNewSiteButton" onClick={() => setShowAddNewSite(!showAddNewSite)} className='btn btn-light'>
            {t('button.addsite')}
          </button>
          {showAddNewSite && <NewSiteForm createSite={createSite} />}
        </div>
      </div>
    </div>
  )
}

export default SiteFlexbox
