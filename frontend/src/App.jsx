import "./App.css"
import RegisterPage from "./pages/RegisterPage"
import TermsPage from "./pages/TermsPage"
import HomePage from "./pages/HomePage"
import LogInPage from "./pages/LogInPage"
import SitePage from "./pages/SitePage"
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  Navigate,
} from "react-router-dom"
import { useState, useEffect } from "react"
import i18n from "./i18n"

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const token = localStorage.getItem("token")
    setIsLoggedIn(!!token)
  }, [])

  const handleLogout = () => {
    localStorage.removeItem("token")
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
          <div style={{position: "absolute", top: "0", right: "0", padding: "8px"}}>
            <a href="#" onClick={() => i18n.changeLanguage('en')} style={{paddingRight: "0.8rem"}}>en</a>
            <a href="#" onClick={() => i18n.changeLanguage('fi')} style={{paddingRight: "0.8rem"}}>fi</a>
          </div>
          <Routes>
            <Route
              path="/"
              element={
                isLoggedIn ? <SitePage /> : <Navigate replace to="/login" />
              }
            />
            <Route
              path="/login"
              element={
                !isLoggedIn ? (
                  <LogInPage
                    onLogin={handleLogout}
                    setIsLoggedIn={setIsLoggedIn}
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
              element={
                isLoggedIn ? <SitePage /> : <Navigate replace to="/login" />
              }
            />
            <Route
              path="/site/:id"
              element={
                isLoggedIn ? <SitePage /> : <Navigate replace to="/login" />
              }
            />
            <Route
              path="/flowers"
              element={
                isLoggedIn ? <HomePage /> : <Navigate replace to="/login" />
              }
            />
          </Routes>
        </div>
      </Router>
    </div>
  )
}

export default App
