import { Outlet, NavLink } from 'react-router-dom'
import './Retailer.css'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'


const tabBar = () => {
    const { t, i18n } = useTranslation()
    return (
      <div className='tab-bar'>
        <div>
          <Nav variant='tabs' defaultActiveKey="/business_owner">
            <Nav.Item>
              <Nav.Link className="menu-tab" as={NavLink} end to="/business_owner"> 
                {t('menu.businessowner')}
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