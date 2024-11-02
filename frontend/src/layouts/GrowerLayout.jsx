import { Outlet, NavLink } from 'react-router-dom'
import { useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'

const navigationBar = () => {
  const params = useParams()
  const { t, i18n } = useTranslation()

  return (
    <div className="nav-container">
      {params.siteId ? (
        <nav>
          <ul>
            <li>
              <NavLink to={`/grower/${params.siteId}`} end>
                {t('menu.home')}
              </NavLink>
            </li>
            <li>
              <NavLink to={`/grower/${params.siteId}/flowers`}>{t('menu.flowers')}</NavLink>
            </li>
            <li>
              <NavLink to={`/grower/${params.siteId}/sites`}>{t('menu.sites')}</NavLink>
            </li>
          </ul>
        </nav>
      ) : (
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
            <li>
              <NavLink to="/grower/sites">{t('menu.sites')}</NavLink>
            </li>
          </ul>
        </nav>
      )}
    </div>
  )
}

const GrowerLayout = () => {
  const { t, i18n } = useTranslation()
  return (
    <div className="layout-container">
      <header className="header">
        <h2>{t('menu.grower')}</h2>
      </header>
      <div className="content">
        <aside className="side-container">
          <h2>{t('title.navigation')}</h2>
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
