import { Outlet, NavLink, Link } from 'react-router-dom'
import './Retailer.css'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'

const tabBar = () => {
  const { t, i18n } = useTranslation()
  return (
    <div className='my-2'>
      <div className="my-3">
        <Nav variant='tabs' defaultActiveKey="/retailer">
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} end to="/retailer"> 
              {t('menu.home')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} end to="/retailer/flowers">
              {t('menu.flowers')}
            </Nav.Link>
          </Nav.Item>
        </Nav>
      </div>
    </div>
  )
}

const RetailerLayout = () => {
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

export default RetailerLayout
