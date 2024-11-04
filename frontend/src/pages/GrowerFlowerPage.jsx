import { useParams } from 'react-router-dom'
import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
const GrowerFlowerPage = () => {
  const params = useParams()
  const [flowers, setFlowers] = useState()
  const { t, i18n } = useTranslation()

  useEffect(() => {
    if (params.siteId) {
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
      <h2>{t('title.site')}: {params.siteId}</h2>
    ) : (
      <h2>{t('title.home')}</h2>
    )}
      <h2>{t('title.flowers')}</h2>
      { flowers ? (<GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} />) : (<GrowerFlowerList flowers={[]} deleteFlower={deleteFlower} />) }
    </>
  )
}

export default GrowerFlowerPage
