import { Outlet, NavLink } from 'react-router-dom'
//import './Grower.css'
import { useTranslation } from 'react-i18next'


const GrowerLayout = () => {
  const { t, i18n } = useTranslation()
  return (
    <div>
      <header>
        <h2>{t('menu.grower')}</h2>
      </header>
      <div>
        <main className="main-container">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

export default GrowerLayout
