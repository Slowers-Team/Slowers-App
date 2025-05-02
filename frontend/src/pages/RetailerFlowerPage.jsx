import RetailerFlowerList from '../components/retailer/RetailerFlowerList'
import flowerService from '../services/flowers'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import WideCenteredCard from '../components/WideCenteredCard'
const RetailerFlowerPage = () => {
  const [flowers, setFlowers] = useState([])
  const { t, i18n } = useTranslation()

  useEffect(() => {
    flowerService.getAll().then(initialFlowers => setFlowers(initialFlowers))
  }, [])

  return (
    <WideCenteredCard>
      <h2>{t('title.flowers')}</h2>
      {flowers && <RetailerFlowerList flowers={flowers} />}
    </WideCenteredCard>
  )
}

export default RetailerFlowerPage
