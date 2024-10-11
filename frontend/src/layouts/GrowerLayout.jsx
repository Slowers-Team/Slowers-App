import { Outlet, NavLink } from 'react-router-dom'
//import './Grower.css'
import { useTranslation } from 'react-i18next'

const navigationBar = () => {
  const { t, i18n } = useTranslation()
  return (
    <div className="nav-container">
      <nav>
        <ul>
          <li>
            <NavLink to="/grower" end>
              {t('menu.home')}
            </NavLink>
          </li>
          <li>
            <NavLink to="/grower/flowers">{t('menu.flowers')}</NavLink>
          </li>
        </ul>
      </nav>
    </div>
  )
}

const GrowerLayout = () => {
  const { t, i18n } = useTranslation()
  return (
    <div className="layout-container">
      <header className="header">
        <h1>Slowers</h1>
      </header>
      <div className="content">
        <aside className="side-container">
          <h1>{t('title.navigation')}</h1>
          {navigationBar()}
        </aside>
        <main className="main-container">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

export default GrowerLayout
