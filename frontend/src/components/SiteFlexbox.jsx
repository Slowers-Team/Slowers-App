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
        <div className={'box'}>
        <Link to={'/flowers'}>{t('title.flowers')}</Link>
        </div>
          {sites && (
            sites.map(site => (
              <div className="box" key={site._id}>
                {t('site.data.name')}: <Link to={'/site/' + site._id}>{site.name}</Link>
                <p>{t('site.data.note')}: {site.note}</p>
              </div>
            ))
          )}
        <div className={'box'}>
          <button id="addNewSiteButton" onClick={() => setShowAddNewSite(!showAddNewSite)}>{t('button.addsite')}</button>
          {showAddNewSite && <NewSiteForm createSite={createSite} />}
        </div>
      </div>
    </div>
    )
}

export default SiteFlexbox
