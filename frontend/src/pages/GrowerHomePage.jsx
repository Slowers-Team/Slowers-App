import { useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'

const GrowerHomePage = () => {
  const params = useParams()
  const { t, i18n } = useTranslation()
  return (
    <>
    {params.siteId ? (
      <h2>{t('title.site')}: {params.siteId}</h2>
    ) : (
      <h2>{t('title.home')}</h2>
    )}
    </>
  )
}

export default GrowerHomePage
