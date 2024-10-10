import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
const GrowerFlowerPage = () => {
  const [flowers, setFlowers] = useState([])
  const { t, i18n } = useTranslation()

  useEffect(() => {
    flowerService.getUserFlowers().then(initialFlowers => setFlowers(initialFlowers))
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
      <h2>{t('title.flowers')}</h2>
      {flowers && <GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} />}
    </>
  )
}

export default GrowerFlowerPage
