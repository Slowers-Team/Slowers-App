import RetailerFlowerList from '../components/retailer/RetailerFlowerList'
import flowerService from '../services/flowers'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
const RetailerFlowerPage = () => {
  const [flowers, setFlowers] = useState([])
  const { t, i18n } = useTranslation()

  useEffect(() => {
    flowerService.getAll().then(initialFlowers => setFlowers(initialFlowers))
  }, [])

  return (
    <>
      <h2>{t('title.flowers')}</h2>
      {flowers && <RetailerFlowerList flowers={flowers} />}
    </>
  )
}

export default RetailerFlowerPage
