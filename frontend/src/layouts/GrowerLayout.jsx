import { Outlet, NavLink, Link } from 'react-router-dom'
//import './Grower.css'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'

const tabBar = () => {
  const { t, i18n } = useTranslation()
  return (
    <div className='my-2'>
      <Link to="/grower" className="mx-2 text-secondary text-decoration-none">Placeholder</Link>
      <div className="my-3">
        <Nav variant='tabs' defaultActiveKey="/grower">
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} end to="/" disabled> 
              {t('menu.home')} (Disabled)
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} end to="/grower/flowers">
              {t('menu.flowers')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} end to="/grower">
              Subsites
            </Nav.Link>
          </Nav.Item>
        </Nav>
      </div>
    </div>
  )
}
const GrowerLayout = () => {
  const { t, i18n } = useTranslation()
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

export default GrowerLayout
