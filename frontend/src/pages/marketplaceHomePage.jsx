import { Container } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'

const MarketplaceHomePage = () => {
  const { t, i18n } = useTranslation()
  return (
    <Container>
      <h2>{t('title.marketplacehome')}</h2>
    </Container>
  )
}

export default MarketplaceHomePage
