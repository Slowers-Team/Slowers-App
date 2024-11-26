import siteService from '../services/sites'

const siteLoader = async ({ params }) => {
  const { site, subsites } = await siteService.get(params.siteId)
  return { site, subsites}
}

const rootSiteLoader = async () => {
  const rootSites = await siteService.get()
  return rootSites
}

export { 
  siteLoader,
  rootSiteLoader,
}
