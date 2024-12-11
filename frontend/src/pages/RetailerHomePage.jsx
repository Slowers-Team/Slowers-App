import { Container } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'

const RetailerHomePage = () => {
  const { t, i18n } = useTranslation()
  return (
    <Container>
      <h2>{t('title.home')}</h2>
    </Container>
  )
}

export default RetailerHomePage
