import siteService from '../services/sites'

const siteLoader = async ({ params }) => {
  const { site, subsites } = await siteService.get(params.siteId)
  console.log(site, subsites)
  return { site, subsites }
}

export default siteLoader
