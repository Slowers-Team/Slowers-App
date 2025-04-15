import { Outlet, NavLink } from 'react-router-dom'
import './Retailer.css'
import './Business.css'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'


const tabBar = () => {
    const { t, i18n } = useTranslation()
    return (
      <div className='tab-bar'>
        <div>
          <Nav variant='tabs' defaultActiveKey="/business">
            <Nav.Item>
              <Nav.Link className="menu-tab" as={NavLink} end to="/business">
                {t('menu.business')}
              </Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link className="menu-tab" as={NavLink} end to="/business/employees">
                {t('menu.employees')}
              </Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link className="menu-tab" as={NavLink} end to="/business/grower/sites">
                {t('menu.sites')}
              </Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link className="menu-tab" as={NavLink} end to="/business/grower/flowers">
                {t('menu.flowers')}
              </Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link className="menu-tab" as={NavLink} end to="/business/retailer">
                {t('menu.retailer')}
              </Nav.Link>
            </Nav.Item>
          </Nav>
        </div>
      </div>
    )
  }

const BusinessLayout = () => {
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

export default BusinessLayout