import { useTranslation } from 'react-i18next'
import WideCenteredCard from '../components/WideCenteredCard'

const MarketplaceHomePage = () => {
  const { t, i18n } = useTranslation()
  
  return (
    <WideCenteredCard>
      <h2>{t('title.marketplacehome')}</h2>
    </WideCenteredCard>
  )
}

export default MarketplaceHomePage
