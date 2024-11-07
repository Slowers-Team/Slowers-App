import { useParams } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import siteService from '../services/sites'

const GrowerHomePage = () => {
  const params = useParams()
  const [site, setSite] = useState()
  const { t, i18n } = useTranslation()

  useEffect(() => {
    if (params.siteId) {
      siteService.get(params.siteId).then(initialSite => setSite(initialSite.site))
    }
  }, [])

  return (
    <>
    {params.siteId ? (
      <h2>{site?.name} {t('title.sitehome')} </h2>
    ) : (
      <h2>{t('title.home')}</h2>
    )}
    </>
  )
}

export default GrowerHomePage
