import { Outlet, NavLink } from 'react-router-dom'
import './Retailer.css'
import './Business.css'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'
import { Authenticator } from '../Authenticator'
import { useState, useEffect } from 'react'


const TabBar = ({ designation }) => {
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
            {(designation === "owner" || designation === "employee") && (
              <Nav.Item>
                <Nav.Link className="menu-tab" as={NavLink} end to="/business/employees">
                  {t('menu.employees')}
                </Nav.Link>
              </Nav.Item>
            )}
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
  const [designation, setterDesignation] = useState(Authenticator.designation)

  useEffect(() => {
    setterDesignation(Authenticator.designation)
  }, [Authenticator.designation])

  return (
    <div>
    	<TabBar designation={ designation } />
      <div>
        <main className="main-container">
          <Outlet context={{ setterDesignation }}/>
        </main>
      </div>
    </div>
  )
}

export default BusinessLayout