import './App.css'
import RegisterPage from './pages/RegisterPage'
import TermsPage from './pages/TermsPage'
import HomePage from './pages/HomePage'
import LogInPage from './pages/LogInPage'
import SitePage from './pages/SitePage'
import GrowerLayout from './layouts/GrowerLayout'
import GrowerFlowerPage from './pages/GrowerFlowerPage'
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom'
import { useState, useEffect } from 'react'

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [isLoading, setIsLoading] = useState(true)

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
    return <div>Loading...</div>
  }

  return (
    <div>
      <Router>
        <div>
          <nav>
            <Link style={padding} to="/">
              Home
            </Link>
            <Link style={padding} to="/grower">
              Grower Page
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
