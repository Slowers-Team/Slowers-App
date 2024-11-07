import { useParams } from 'react-router-dom'
import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import siteService from '../services/sites'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
const GrowerFlowerPage = () => {
  const params = useParams()
  const [flowers, setFlowers] = useState()
  const [site, setSite] = useState()
  const { t, i18n } = useTranslation()

  useEffect(() => {
    if (params.siteId) {
      siteService.get(params.siteId).then(initialSite => setSite(initialSite.site))
      flowerService.getFlowersBySite(params.siteId).then(initialFlowers => setFlowers(initialFlowers))
    } else {
      flowerService.getUserFlowers().then(initialFlowers => setFlowers(initialFlowers))
    }
  }, [])

  const deleteFlower = flowerObject => {
    if (window.confirm(`${t('label.confirmflowerdeletion')} ${flowerObject.name}?`)) {
      flowerService.remove(flowerObject._id).then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => item._id !== flowerObject._id))
      })
    }
  }

  return (
    <>
    {params.siteId ? (
      <h2>{site?.name} {t('title.siteflowers')}</h2>
    ) : (
      <h2>{t('title.allflowers')}</h2>
    )}
      { flowers ? (<GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} />) : (<GrowerFlowerList flowers={[]} deleteFlower={deleteFlower} />) }
    </>
  )
}

export default GrowerFlowerPage
