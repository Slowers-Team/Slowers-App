import { Outlet, NavLink, Link } from 'react-router-dom'
import './Retailer.css'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'

const tabBar = () => {
  const { t, i18n } = useTranslation()
  return (
    <div className='tab-bar'>
      <div>
        <Nav variant='tabs' defaultActiveKey="/marketplace">
          <Nav.Item>
            <Nav.Link className="menu-tab" as={NavLink} end to="/marketplace"> 
              {t('menu.home')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link className="menu-tab" as={NavLink} end to="/marketplace/flowers">
              {t('menu.marketplace')}
            </Nav.Link>
          </Nav.Item>
        </Nav>
      </div>
    </div>
  )
}

const MarketplaceLayout = () => {
  return (
    <div>
      {tabBar()}
      <div>
        <main className="main-container">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

export default MarketplaceLayout
