import './App.css'
import ProtectedRoute from './ProtectedRoute'
import RegisterPage from './pages/RegisterPage'
import TermsPage from './pages/TermsPage'
import HomePage from './pages/HomePage'
import LogInPage from './pages/LogInPage'
import UserPage from './pages/UserPage'
import RetailerHomePage from './pages/RetailerHomePage'
import RetailerFlowerPage from './pages/RetailerFlowerPage'
import RetailerLayout from './layouts/RetailerLayout'
import GrowerLayout from './layouts/GrowerLayout'
import GrowerHomePage from './pages/GrowerHomePage'
import GrowerFlowerPage from './pages/GrowerFlowerPage'
import GrowerSitesPage from './pages/GrowerSitesPage'
import LangSelect from './components/LangSelect'
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Navbar, Nav, NavDropdown } from 'react-bootstrap'

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [defaultRole, setDefaultRole] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const { t, i18n } = useTranslation()

  useEffect(() => {
    const token = localStorage.getItem('token')
    const role = localStorage.getItem('role')
    setIsLoggedIn(!!token)
    setDefaultRole(role)
    setIsLoading(false)

    const langCookie = document.cookie.split('; ').find(row => row.startsWith('lang='))
    const language = langCookie ? langCookie.split('=')[1] : 'en'
    i18n.changeLanguage(language)
  }, [])

  const handleLogout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('role')
    setIsLoggedIn(false)
    setDefaultRole('')
  }

  const padding = {
    padding: 5,
  }

  if (isLoading) {
    return <div>{t('label.loading')}</div>
  }

  return (
    <div>
      <Router>
        <div>
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

          <Routes>
            <Route element={<ProtectedRoute isLoggedIn={isLoggedIn} />}>

              <Route path="/" element={defaultRole == 'retailer' ? <Navigate replace to="/retailer" /> : <Navigate replace to="/grower" />} />

              <Route path="/grower" element={<GrowerLayout />}>
                <Route index element={<GrowerHomePage />} />
                <Route path="flowers" element={<GrowerFlowerPage />} />
                <Route path="sites" element={<GrowerSitesPage />} />
              </Route>

              <Route path="/grower/:siteId" element={<GrowerLayout />}>
                <Route index element={<GrowerHomePage />} />
                <Route path="flowers" element={<GrowerFlowerPage />} />
                <Route path="sites" element={<GrowerSitesPage />} />
              </Route>

              <Route path="/retailer" element={<RetailerLayout />}>
                <Route index element={<RetailerHomePage />} />
                <Route path="flowers" element={<RetailerFlowerPage />} />
              </Route>

              <Route path="/user" element={<UserPage />} />
              {/* <Route path="/site" element={<SitePage />} />
              <Route path="/site/:siteId" element={<SitePage />} /> */}

              {/* Lisää kirjautumista vaativat routet tänne */}
            </Route>

            <Route path="/login" element={<LogInPage
              onLogin={handleLogout}
              setIsLoggedIn={setIsLoggedIn}
              setDefaultRole={setDefaultRole}
            />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/terms" element={<TermsPage />} />
            <Route path="/flowers" element={<HomePage />} />
          </Routes>

        </div>
      </Router>
    </div>
  )
}

export default App
