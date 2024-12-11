import { Outlet, NavLink, Link, useParams } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Nav } from 'react-bootstrap'
import siteService from '../services/sites'
import SiteBreadCrumbs from '../components/SiteBreadCrumbs' 

const tabBar = () => {
  const params = useParams()
  const [site, setSite] = useState()
  const { t, i18n } = useTranslation()

  useEffect(()=>{
    if (params.siteId && site?._id != params.siteId) {
      siteService.get(params.siteId).then((s) => setSite(s)).catch((err) => console.error(err))
    }
  },[params.siteId])

  return (
    <div className='my-2'>
      {site ? <SiteBreadCrumbs props={site} /> : <></>}
      <Link to="/grower" className="mx-2 text-secondary text-decoration-none">Placeholder</Link>
      <div className="my-3">
      {params.siteId ? (
        <Nav variant='tabs' defaultActiveKey="/grower">
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} end to={`/grower/${params.siteId}`}> 
              {t('menu.home')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} to={`/grower/${params.siteId}/flowers`}>
              {t('menu.flowers')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link className="text-success" as={NavLink} to={`/grower/${params.siteId}/sites`}>
              {t('menu.sites')}
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
           <Nav.Link className="text-success" as={NavLink} to={`/grower/${params.siteId}/images`}>
              {t('menu.images')}
            </Nav.Link>
          </Nav.Item>
        </Nav>
      ) : (
        <Nav variant='tabs' defaultActiveKey="/grower">
        <Nav.Item>
          <Nav.Link className="text-success" as={NavLink} end to="/grower"> 
            {t('menu.home')}
          </Nav.Link>
        </Nav.Item>
        <Nav.Item>
          <Nav.Link className="text-success" as={NavLink} to="/grower/flowers">
            {t('menu.flowers')}
          </Nav.Link>
        </Nav.Item>
        <Nav.Item>
          <Nav.Link className="text-success" as={NavLink} to="/grower/sites">
            {t('menu.sites')}
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
