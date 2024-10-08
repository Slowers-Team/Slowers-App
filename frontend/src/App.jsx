import "./App.css"
import RegisterPage from "./pages/RegisterPage"
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
import { Navbar, Nav } from "react-bootstrap"

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
          <Navbar collapseOnSelect expand="lg" bg="light" variant="light">
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">
              <Nav className="mr.auto">
                <Nav.Link href="#" as="span">
                  <Link style={padding} to="/">Home</Link>
                </Nav.Link>
                <Nav.Link href="#" as="span">
                  {!isLoggedIn && (
                    <Link style={padding} to="/register">Register</Link>
                  )}
                </Nav.Link>
                <Nav.Link href="#" as="span">
                  {!isLoggedIn && (
                    <Link style={padding} to="/login">Login</Link>
                  )}
                </Nav.Link>
                <Nav.Link href="#" as="span">
                  {isLoggedIn && <Link onClick={handleLogout}>Logout</Link>}
                </Nav.Link>
              </Nav>
            </Navbar.Collapse>
          </Navbar>
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
