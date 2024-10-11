import { useState, useEffect } from 'react'
import flowerService from '../services/flowers'
import FlowerList from '../components/FlowerList'
import { useTranslation } from 'react-i18next'

const HomePage = () => {
  const [flowers, setFlowers] = useState([])
  const { t, i18n } = useTranslation()

  useEffect(() => {
    flowerService.getUserFlowers().then(initialFlowers => setFlowers(initialFlowers))
  }, [])

  const deleteFlower = flowerObject => {
    if (
      window.confirm(
        `${t('label.confirmflowerdeletion')} ${flowerObject.name}?`
      )
    ) {
      flowerService.remove(flowerObject._id).then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => item._id !== flowerObject._id))
      })
    }
  }

  return (
    <div>
      <h2>{t('title.flowers')}</h2>
      {flowers && <FlowerList flowers={flowers} deleteFlower={deleteFlower} />}
    </div>
  )
}

export default HomePage
