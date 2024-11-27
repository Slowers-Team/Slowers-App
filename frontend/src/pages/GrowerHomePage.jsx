import { useTranslation } from 'react-i18next'
import { useRouteLoaderData, useLoaderData } from 'react-router-dom' 

const GrowerHomePage = () => {
  const { t, i18n } = useTranslation()
  const { site } = useLoaderData() ?? useRouteLoaderData("site") // very ugly hack that makes site "false"

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
