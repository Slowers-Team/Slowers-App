import { useTranslation } from 'react-i18next'
import { useRouteLoaderData } from 'react-router-dom' 

const GrowerHomePage = () => {
  const { t, i18n } = useTranslation()
  const { site } = useRouteLoaderData("site")

  return (
    <>
      {site ? (
        <h2>{site?.name} {t('title.sitehome')} </h2>
      ) : (
        <h2>{t('title.home')}</h2>
      )}
      {site?.note && <p className='mx-1'>{t('site.data.note')} : {site?.note}</p>}
    </>
  )
}

export default GrowerHomePage
