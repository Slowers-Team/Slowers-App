import { useTranslation } from 'react-i18next'

const MarketplaceHomePage = () => {
  const { t, i18n } = useTranslation()
  
  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-12">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2>{t('title.marketplacehome')}</h2>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default MarketplaceHomePage
