import { useTranslation } from 'react-i18next'

const GrowerHomePage = () => {
  const { t, i18n } = useTranslation()
  return (
    <>
      <h2>{t('title.home')}</h2>
    </>
  )
}

export default GrowerHomePage
