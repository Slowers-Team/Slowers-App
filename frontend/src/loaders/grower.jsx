import siteService from '../services/sites'
import flowerService from '../services/flowers'
import imageService from '../services/images'
import { defer } from 'react-router-dom'

const siteLoader = async ({ params }) => {
  let { site, subsites } = await siteService.get(params.siteId)
  let images = imageService.getImagesByEntity(params.siteId) 
  return defer({ site: site, subsites: subsites, images: images})
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
