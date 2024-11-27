import '../App.css'
import { useTranslation } from 'react-i18next'
import { Navbar, Nav, NavDropdown, Button, Offcanvas } from 'react-bootstrap'
import LangSelect from './LangSelect'
import { Link, useLoaderData, Outlet, useFetcher } from 'react-router-dom'
import { useState } from 'react'

export const NavigationBar = () => {
    const { t, i18n } = useTranslation()
    const [showOffCanvas, setShowOffCanvas] = useState(false)
    const {isLoggedIn, username} = useLoaderData()
    const fetcher = useFetcher()

    const handleClose = () => setShowOffCanvas(false)
    const handleShow = () => setShowOffCanvas(!showOffCanvas)

    const handleLogout = () => fetcher.submit({},{action: "/logout", method: "post"})
  
    return (
        <>
        <Navbar expand="sm" bg="light">
            <Button variant="light" className="mx-2" onClick={handleShow}>
              <span className='navbar-toggler-icon'></span>
            </Button>
            <Navbar.Brand as={Link} to="/">
              <h1>Slowers</h1>
            </Navbar.Brand>
              <Nav className="ms-auto mx-2">
                {isLoggedIn && (
                  <NavDropdown title={username} id="collasible-nav-dropdown">
                    <Nav.Link className="text-secondary" as={Link} to="/user">
                      {t("menu.profile")}
                    </Nav.Link>
                    <Nav.Link className="text-secondary" as={Link} onClick={() => { handleLogout(); handleClose(); }}>
                      {t("menu.logout")}
                    </Nav.Link>
                  </NavDropdown>
                )}
                <NavDropdown title={t("menu.language")} id="collasible-nav-dropdown">
                  <LangSelect/>
                </NavDropdown>
              </Nav>
          </Navbar>

          <Offcanvas show={showOffCanvas} onHide={handleClose} className="offcanvas-thin">
            <Offcanvas.Header closeButton>
              <Offcanvas.Title> 
                <h2>Slowers</h2>
              </Offcanvas.Title>
            </Offcanvas.Header>
            <Offcanvas.Body>
            <Nav className="flex-column pe-3">
                <Nav.Link className="text-secondary" as={Link} to="/" onClick={handleClose}>
                  {t("menu.home")}
                </Nav.Link>
                {!isLoggedIn && (
                  <Nav.Link className="text-secondary" as={Link} to="/login" onClick={handleClose}>
                    {t("menu.login")}
                  </Nav.Link>
                )}
                {!isLoggedIn && (
                  <Nav.Link className="text-secondary" as={Link} to="/register" onClick={handleClose}>
                    {t("menu.register")}
                  </Nav.Link>
                )}
                {isLoggedIn && (
                  <Nav.Link className="text-secondary" as={Link} to="/retailer" onClick={handleClose}>
                    {t("menu.retailer")}
                  </Nav.Link>
                )}
                {isLoggedIn && (
                  <Nav.Link className="text-secondary" as={Link} to="/grower" onClick={handleClose}>
                    {t("menu.grower")}
                  </Nav.Link>
                )}
                <Nav.Link className="text-secondary" as={Link} to="/terms" onClick={handleClose}>
                  {t("menu.terms")}
                </Nav.Link>
              </Nav>
            </Offcanvas.Body>
          </Offcanvas>
          <Outlet />
        </>
    )
}

export default NavigationBar
