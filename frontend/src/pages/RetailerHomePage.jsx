import { useTranslation } from 'react-i18next'

const RetailerHomePage = () => {
  const { t, i18n } = useTranslation()
  return (
    <>
      <h2>{t('title.home')}</h2>
    </>
  )
}

export default RetailerHomePage
