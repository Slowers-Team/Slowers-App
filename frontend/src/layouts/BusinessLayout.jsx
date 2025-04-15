import { Outlet, NavLink } from 'react-router-dom'
import './Retailer.css'
import './Business.css'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'
import { Authenticator } from '../Authenticator'
import { useState, useEffect } from 'react'


const TabBar = ({ designation, businessType }) => {
    const { t, i18n } = useTranslation()
    const accessToEmployeesTab = designation === "owner" || designation === "employee"
    const accessToRetailerTab = businessType === "retailer"
    return (
      <div className='tab-bar'>
        <div>
          <Nav variant='tabs' defaultActiveKey="/business">
            <Nav.Item>
              <Nav.Link className="menu-tab" as={NavLink} end to="/business">
                {t('menu.business')}
              </Nav.Link>
            </Nav.Item>
            { accessToEmployeesTab && (
              <Nav.Item>
                <Nav.Link className="menu-tab" as={NavLink} end to="/business/employees">
                  {t('menu.employees')}
                </Nav.Link>
              </Nav.Item>
            )}
            {( accessToEmployeesTab && accessToRetailerTab) && (
              <Nav.Item>
                <Nav.Link className="menu-tab" as={NavLink} end to="/business/retailer">
                  {t('menu.retailer')}
                </Nav.Link>
              </Nav.Item>
            )}
          </Nav>
        </div>
      </div>
    )
  }

const BusinessLayout = () => {
  const [designation, setterDesignation] = useState(Authenticator.designation)
  const [businessType, setterBusinessType] = useState(Authenticator.businessType)

  useEffect(() => {
    setterDesignation(Authenticator.designation)
    setterBusinessType(Authenticator.businessType)
  }, [Authenticator.designation, Authenticator.businessType])

  return (
    <div>
    	<TabBar designation={designation} businessType={ businessType } />
      <div>
        <main className="main-container">
          <Outlet context={{ setterDesignation, setterBusinessType }}/>
        </main>
      </div>
    </div>
  )
}

export default BusinessLayout