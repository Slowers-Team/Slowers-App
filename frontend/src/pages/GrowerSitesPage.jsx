import { useParams, useNavigate, useRouteLoaderData, useLoaderData, useFetcher } from 'react-router-dom'
import SiteService from '../services/sites'
import SiteFlexbox from '../components/SiteFlexbox'
import { useTranslation } from "react-i18next"

const GrowerSitesPage = () => {
  const params = useParams()
  const sitePage = Boolean(params.siteId)
  const navigate = useNavigate()
  const site = sitePage ? useRouteLoaderData("site").site : null
  const sites = sitePage 
    ? useRouteLoaderData("site").subsites
    : useLoaderData()
  const { t, i18n } = useTranslation()
  const fetcher = useFetcher()

  const createSite = siteObject => {
    SiteService.create(siteObject)
      .then(_ => {
        fetcher.submit({}, {action: "/grower/sites", method: "post"})
      })
      .catch(error => {
        const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '')
        alert(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data))
      })
  }

  const deleteSite = siteObject => {
    if (window.confirm(`${t("label.confirmsitedeletion")} ${siteObject.name}?`)) {
      const parentId = siteObject.parent ? siteObject.parent : ''
      SiteService.remove(siteObject._id)
        .then(() => {
          const redirect = (parentId !== null && parentId !== '')
            ? '/grower/' + parentId + '/sites'
            : '/grower/sites'
          fetcher.submit({redirect: redirect}, {action: "/grower", method: "delete"})
        })
        .catch(error => {
          console.error('Error deleting site:', error)
        })
    }
  }

  const handleBack = () => {
    navigate(-1)
  }

  return (
    <>
      {params.siteId ? (
        <div>
          <h2>{site?.name} {t("title.sitesites")}</h2>
          <div className='my-2'>
            <button onClick={handleBack} style={{ marginRight: "0.5rem" }} className="btn btn-light">{t("button.goback")}</button>
            <button id="deleteSiteButton" onClick={() => deleteSite(site)} className="btn btn-light">{t("button.deletethissite")}</button>
          </div>
          <SiteFlexbox createSite={createSite} sites={sites} />
        </div>
      ) : (
        <div>
          <h2 className='mb-3'>{t("title.sites")}</h2>
          <SiteFlexbox createSite={createSite} sites={sites} />
        </div>
      )}
    </>
  )
}

export default GrowerSitesPage
