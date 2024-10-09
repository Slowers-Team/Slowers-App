import './App.css'
import RegisterPage from './pages/RegisterPage'
import HomePage from './pages/HomePage'
import LogInPage from './pages/LogInPage'
import SitePage from './pages/SitePage'
import RetailerHomePage from './pages/RetailerHomePage'
import RetailerFlowerPage from './pages/RetailerFlowerPage'
import RetailerLayout from './layouts/RetailerLayout'
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom'
import { useState, useEffect } from 'react'
import './App.css'
import RegisterPage from './pages/RegisterPage'
import TermsPage from './pages/TermsPage'
import HomePage from './pages/HomePage'
import LogInPage from './pages/LogInPage'
import SitePage from './pages/SitePage'
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom'
import { useState, useEffect } from 'react'

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const token = localStorage.getItem('token')
    setIsLoggedIn(!!token)
  }, [])

  const handleLogout = () => {
    localStorage.removeItem('token')
    setIsLoggedIn(false)
  }

  const padding = {
    padding: 5,
  }

  return (
    <div>
      <Router>
        <div>
          <nav>
            <Link style={padding} to="/">
              Home
            </Link>
            <Link style={padding} to="/retailer">
              Retailer Page
            </Link>
            {!isLoggedIn && (
              <Link style={padding} to="/register">
                Register
              </Link>
            )}
            {!isLoggedIn && (
              <Link style={padding} to="/login">
                Login
              </Link>
            )}
            {isLoggedIn && <Link onClick={handleLogout}>Logout</Link>}
            <Link style={padding} to="/terms">
              Terms
            </Link>
          </nav>
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
          </Routes>
        </div>
      </Router>
    </div>
  )
}

export default App
