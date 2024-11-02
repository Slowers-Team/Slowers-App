import { useTranslation } from 'react-i18next'
import { Navbar, Nav, NavDropdown } from 'react-bootstrap'
import LangSelect from './components/LangSelect'
import { Link } from 'react-router-dom'

export const NavigationBar = ({ isLoggedIn, handleLogout }) => {
    const { t, i18n } = useTranslation()

    return (
        <>
        <Navbar collapseOnSelect expand="lg" bg="light">
            <Navbar.Brand>
              <h1 className="mx-3 text-center">Slowers</h1>
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">
              <Nav className="me-auto justify-content-start">
                <Nav.Link as={Link} to="/">
                  {t('menu.home')}
                </Nav.Link>
                {!isLoggedIn && (
                  <Nav.Link as={Link} to="/login">
                    {t('menu.login')}
                  </Nav.Link>
                )}
                {!isLoggedIn && (
                  <Nav.Link as={Link} to="/register">
                    {t('menu.register')}
                  </Nav.Link>
                )}
                {isLoggedIn && (
                  <NavDropdown title={t('menu.role')} id="collasible-nav-dropdown">
                    <Nav.Link as={Link} to="/retailer">
                      {t('menu.retailer')}
                    </Nav.Link>
                    <Nav.Link as={Link} to="/grower">
                      {t('menu.grower')}
                    </Nav.Link>
                  </NavDropdown>
                )}
                {isLoggedIn && (
                  <Nav.Link as={Link} onClick={handleLogout}>
                    {t('menu.logout')}
                  </Nav.Link>
                )}
              </Nav>
              <Nav className="ms-auto">
                {isLoggedIn && (
                  <Nav.Link as={Link} to="/user">
                    {t('menu.profile')}
                  </Nav.Link>
                )}
                <Nav.Link as={Link} to="/terms">
                  {t('menu.terms')}
                </Nav.Link>
                <NavDropdown title={t('menu.language')} id="collasible-nav-dropdown">
                  <LangSelect />
                </NavDropdown>
              </Nav>
            </Navbar.Collapse>
          </Navbar>
        </>
    )
}

export default NavigationBar