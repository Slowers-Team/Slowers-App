import { Container } from 'react-bootstrap'
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
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-12">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2>{t('title.flowers')}</h2>
              {flowers && <RetailerFlowerList flowers={flowers} />}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default RetailerFlowerPage
