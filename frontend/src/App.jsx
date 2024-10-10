import './App.css'
import RegisterPage from './pages/RegisterPage'
import TermsPage from './pages/TermsPage'
import HomePage from './pages/HomePage'
import LogInPage from './pages/LogInPage'
import SitePage from './pages/SitePage'
import RetailerHomePage from './pages/RetailerHomePage'
import RetailerFlowerPage from './pages/RetailerFlowerPage'
import RetailerLayout from './layouts/RetailerLayout'
import GrowerLayout from './layouts/GrowerLayout'
import GrowerFlowerPage from './pages/GrowerFlowerPage'
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom'
import { useState, useEffect } from 'react'
import i18n from "./i18n"
import { useTranslation } from 'react-i18next'

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [isLoading, setIsLoading] = useState(true)
  const { t, i18n } = useTranslation()

  useEffect(() => {
    const token = localStorage.getItem('token')
    setIsLoggedIn(!!token)
    setIsLoading(false)
  }, [])

  const handleLogout = () => {
    localStorage.removeItem('token')
    setIsLoggedIn(false)
  }

  const padding = {
    padding: 5,
  }

  if (isLoading) {
    return <div>{t("label.loading")}</div>
  }

  return (
    <div>
      <Router>
        <div>
          <nav>
            <Link style={padding} to="/">
              {t("menu.home")}
            </Link>
            <Link style={padding} to="/retailer">
              {t("menu.retailer")}
            </Link>
            <Link style={padding} to="/grower">
              {t("menu.grower")}
            </Link>
            {!isLoggedIn && (
              <Link style={padding} to="/register">
                {t("menu.register")}
              </Link>
            )}
            {!isLoggedIn && (
              <Link style={padding} to="/login">
                {t("menu.login")}
              </Link>
            )}
            {isLoggedIn && <Link onClick={handleLogout}>{t("menu.logout")}</Link>}
            <Link style={padding} to="/terms">
              {t("menu.terms")}
            </Link>
          </nav>
          <div style={{position: "absolute", top: "0", right: "0", padding: "8px"}}>
            <a href="#" onClick={() => i18n.changeLanguage('en')} style={{paddingRight: "0.8rem"}}>en</a>
            <a href="#" onClick={() => i18n.changeLanguage('fi')} style={{paddingRight: "0.8rem"}}>fi</a>
          </div>
          <Routes>
            <Route
              path="/"
              element={isLoggedIn ? <SitePage /> : <Navigate replace to="/login" />}
            />
            <Route
              path="/login"
              element={
                !isLoggedIn ? (
                  <LogInPage onLogin={handleLogout} setIsLoggedIn={setIsLoggedIn} />
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
          </Routes>
        </div>
      </Router>
    </div>
  )
}

export default App
