import siteService from '../services/sites'
import flowerService from '../services/flowers'

const siteLoader = async ({ params }) => {
  const { site, subsites } = await siteService.get(params.siteId)
  return { site, subsites}
}

const rootSiteLoader = async () => {
  const rootSites = await siteService.get()
  return rootSites
}

const siteFlowerLoader = async ( { params } ) => {
    return await flowerService.getFlowersBySite(params.siteId)
}

const rootFlowerLoader = async () => {
  return await flowerService.getUserFlowers()
}

export { 
  siteLoader,
  rootSiteLoader,
  siteFlowerLoader,
  rootFlowerLoader
}
