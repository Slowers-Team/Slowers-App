import './App.css'
import RegisterPage from './pages/RegisterPage'
import TermsPage from './pages/TermsPage'
import HomePage from './pages/HomePage'
import LogInPage from './pages/LogInPage'
import SitePage from './pages/SitePage'
import UserPage from './pages/UserPage'
import RetailerHomePage from './pages/RetailerHomePage'
import RetailerFlowerPage from './pages/RetailerFlowerPage'
import RetailerLayout from './layouts/RetailerLayout'
import GrowerLayout from './layouts/GrowerLayout'
import GrowerFlowerPage from './pages/GrowerFlowerPage'
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom'
import { useState, useEffect } from 'react'
import i18n from "./i18n"
import { useTranslation } from 'react-i18next'
import { Navbar, Nav, NavDropdown } from "react-bootstrap"

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);  
  const [defaultRole, setDefaultRole] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const { t, i18n } = useTranslation()

  useEffect(() => {
    const token = localStorage.getItem('token')
    const role = localStorage.getItem('role')
    setIsLoggedIn(!!token)
    setDefaultRole(role)
    setIsLoading(false)

    const langCookie = document.cookie.split("; ").find(row => row.startsWith("lang="))
    const language = langCookie ? langCookie.split("=")[1] : "en"
    i18n.changeLanguage(language)
  }, [])

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('role');
    setIsLoggedIn(false);
    setDefaultRole('')
  };

  const padding = {
    padding: 5,
  };

  if (isLoading) {
    return <div>{t("label.loading")}</div>
  }

  const changeLanguage = lang => {
    document.cookie = `lang=${lang}; expires=${new Date(Date.now().valueOf() + 2592000000).toUTCString()}; path=/`;
    i18n.changeLanguage(lang);
  };

  return (
    <div>
      <Router>
        <div>
          <Navbar collapseOnSelect expand="lg" bg="light">
            <Navbar.Brand>
              <h1 className='mx-3 text-center'>Slowers</h1>
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">
              <Nav className="me-auto justify-content-start">
                <Nav.Link as={Link} to="/">
                  {t("menu.home")}
                </Nav.Link>
                {!isLoggedIn && (
                  <Nav.Link as={Link} to="/login">
                    {t("menu.login")}
                  </Nav.Link>
                )}
                {!isLoggedIn && (
                  <Nav.Link as={Link} to="/register">
                    {t("menu.register")}
                  </Nav.Link>
                )}
                {isLoggedIn && (
                <NavDropdown title={t("menu.role")} id="collasible-nav-dropdown">
                  <Nav.Link as={Link} to="/retailer">
                    {t("menu.retailer")}
                  </Nav.Link>
                  <Nav.Link as={Link} to="/grower">
                    {t("menu.grower")}
                  </Nav.Link>
                </NavDropdown>
                )}
                {isLoggedIn && (
                  <Nav.Link as={Link} onClick={handleLogout}>
                    {t("menu.logout")}
                  </Nav.Link>
                )}
              </Nav>
              <Nav className="ms-auto">
                {isLoggedIn && (
                  <Nav.Link as={Link} to="/user">
                    {t("menu.profile")}
                  </Nav.Link>
                )}
                <Nav.Link as={Link} to="/terms">
                  {t("menu.terms")}
                </Nav.Link>
                <NavDropdown title={t("menu.language")} id="collasible-nav-dropdown">
                  <Nav.Link href="#" onClick={() => changeLanguage('en')}>
                    en
                  </Nav.Link>
                  <Nav.Link href="#" onClick={() => changeLanguage('fi')}>
                    fi
                  </Nav.Link>
                </NavDropdown>
              </Nav>
            </Navbar.Collapse>
          </Navbar>
          <Routes>
            <Route
              path="/"
              element={
                isLoggedIn 
                  ? defaultRole == 'retailer' 
                    ? <Navigate replace to='/retailer'/> 
                    : <Navigate replace to='/grower'/>
                  : <Navigate replace to="/login" />
              }
            />
            <Route
              path="/login"
              element={
                !isLoggedIn ? (
                  <LogInPage
                    onLogin={handleLogout}
                    setIsLoggedIn={setIsLoggedIn}
                    setDefaultRole={setDefaultRole}
                  />
                ) : (
                  <Navigate replace to="/" />
                )
              }
            />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/terms" element={<TermsPage />} />
            <Route
              path="/site"
              element={isLoggedIn ? <SitePage /> : <Navigate replace to="/login" />}
            />
            <Route
              path="/site/:id"
              element={isLoggedIn ? <SitePage /> : <Navigate replace to="/login" />}
            />
            <Route
              path="/flowers"
              element={isLoggedIn ? <HomePage /> : <Navigate replace to="/login" />}
            />

            <Route
              path="/retailer"
              element={isLoggedIn ? <RetailerLayout /> : <Navigate replace to="/login" />}
            >
              <Route index element={<RetailerHomePage />} />
              <Route path="flowers" element={<RetailerFlowerPage />} />
            </Route>
            <Route
              path="/grower"
              element={isLoggedIn ? <GrowerLayout /> : <Navigate replace to="/login" />}
            >
              <Route index element={<SitePage />} />
              <Route path="flowers" element={<GrowerFlowerPage />} />
            </Route>
            <Route
              path="/user"
              element={isLoggedIn ? <UserPage setDefaultRole={setDefaultRole}/> : <Navigate replace to="/login" />}
            />
          </Routes>
        </div>
      </Router>
    </div>
  );
};

export default App;