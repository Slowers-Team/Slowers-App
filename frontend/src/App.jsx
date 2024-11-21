import ProtectedRoute from './components/ProtectedRoute'
import RegisterPage from './pages/RegisterPage'
import TermsPage from './pages/TermsPage'
import LogInPage from './pages/LogInPage'
import UserPage from './pages/UserPage'
import RetailerHomePage from './pages/RetailerHomePage'
import RetailerFlowerPage from './pages/RetailerFlowerPage'
import RetailerLayout from './layouts/RetailerLayout'
import GrowerLayout from './layouts/GrowerLayout'
import GrowerHomePage from './pages/GrowerHomePage'
import GrowerFlowerPage from './pages/GrowerFlowerPage'
import GrowerSitesPage from './pages/GrowerSitesPage'
import GrowerImagesPage from './pages/GrowerImagesPage'
import { BrowserRouter as Router, Routes, Route, Navigate, createBrowserRouter, RouterProvider } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import NavigationBar from './components/NavigationBar'


const Root = () => {
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
    setLanguage()
  }, [])

  const setLanguage = () => {
    const langCookie = document.cookie.split('; ').find(row => row.startsWith('lang='))
    const language = langCookie ? langCookie.split('=')[1] : 'en'
    i18n.changeLanguage(language)
  }

  const getDefaultRole = () => {
    return defaultRole === 'retailer' ? <Navigate replace to="/retailer" /> : <Navigate replace to="/grower" />
  }

  const handleLogin = (token, role) => {
    localStorage.setItem("token", token)
    localStorage.setItem("role", role)
    setIsLoggedIn(true)
    setDefaultRole(role)
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('role')
    setIsLoggedIn(false)
    setDefaultRole('')
  }

  if (isLoading) {
    return <div>{t('label.loading')}</div>
  }

  return (
    <div>
      <NavigationBar isLoggedIn={isLoggedIn} handleLogout={handleLogout}/>
      <Routes>

        <Route element={<ProtectedRoute isLoggedIn={isLoggedIn} />}>
          <Route path="/" element={getDefaultRole()} />
          <Route path="/grower/*" element={<GrowerRoutes />} />
          <Route path="/retailer/*" element={<RetailerRoutes />} />
          <Route path="/user" element={<UserPage setDefaultRole={setDefaultRole}/>} />
        </Route>

        <Route path="/login" element={isLoggedIn ? getDefaultRole() : (
          <LogInPage
          onLogin={handleLogin}/>
        )} />
        
        <Route path="/register" element={isLoggedIn ? getDefaultRole() : <RegisterPage handleLogin={handleLogin} />} />
        <Route path="/terms" element={<TermsPage />} />

      </Routes>
    </div>
  )
}

const GrowerRoutes = () => (
  <Routes>
    <Route element={<GrowerLayout />}>
      <Route path="/">
        <Route index element={<GrowerHomePage />} />
        <Route path="flowers" element={<GrowerFlowerPage />} />
        <Route path="sites" element={<GrowerSitesPage />} />
      </Route>

      <Route path="/:siteId">
        <Route index element={<GrowerHomePage />} />
        <Route path="flowers" element={<GrowerFlowerPage />} />
        <Route path="sites" element={<GrowerSitesPage />} />
        <Route path="images" element={<GrowerImagesPage />} />
      </Route>
    </Route>
  </Routes>
)

const RetailerRoutes = () => (
  <Routes>
    <Route element={<RetailerLayout />} >
      <Route index element={<RetailerHomePage />} />
      <Route path="flowers" element={<RetailerFlowerPage />} />
    </Route>
  </Routes>
)

const router = createBrowserRouter([
  { path: "*", element: <Root />}
])

export default function App() {
  return <RouterProvider router={router} />
}
