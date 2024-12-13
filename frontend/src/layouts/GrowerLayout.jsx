import { Outlet, NavLink, Link } from 'react-router-dom'
import { useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'

const tabBar = () => {
  const params = useParams()
  const { t, i18n } = useTranslation()

  return (
    <div className='tab-bar'>
      <div>
      {params.siteId ? (
        <Nav variant='tabs' defaultActiveKey="/grower">
          <Nav.Item>
            <Nav.Link className="menu-tab" as={NavLink} end to={`/grower/${params.siteId}`}> 
              {t('menu.home')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link className="menu-tab" as={NavLink} to={`/grower/${params.siteId}/flowers`}>
              {t('menu.flowers')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
           <Nav.Link className="menu-tab" as={NavLink} to={`/grower/${params.siteId}/images`}>
              {t('menu.images')}
            </Nav.Link>
          </Nav.Item>
        </Nav>
      ) : (
        <Nav variant='tabs' defaultActiveKey="/grower">
        <Nav.Item>
          <Nav.Link className="menu-tab" as={NavLink} end to="/grower"> 
            {t('menu.home')}
          </Nav.Link>
        </Nav.Item>
        <Nav.Item>
          <Nav.Link className="menu-tab" as={NavLink} to="/grower/flowers">
            {t('menu.flowers')}
          </Nav.Link>
        </Nav.Item>
      </Nav>
      )}
      </div>
    </div>
  )
}
const GrowerLayout = () => {
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
