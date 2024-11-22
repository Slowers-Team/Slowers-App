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
import { Routes, Route, Navigate, createBrowserRouter, RouterProvider, redirect } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import NavigationBar from './components/NavigationBar'
import { Authenticator } from './Authenticator'

const getDefaultRole = () => {
  return localStorage.getItem('role') === 'retailer' ? <Navigate replace to="/retailer" /> : <Navigate replace to="/grower" />
}

const Root = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [isLoading, setIsLoading] = useState(true)
  const { t, i18n } = useTranslation()

  useEffect(() => {
    const token = localStorage.getItem('token')
    setIsLoggedIn(!!token)
    setIsLoading(false)
    setLanguage()
  }, [])

  const setLanguage = () => {
    const langCookie = document.cookie.split('; ').find(row => row.startsWith('lang='))
    const language = langCookie ? langCookie.split('=')[1] : 'en'
    i18n.changeLanguage(language)
  }

  if (isLoading) {
    return <div>{t('label.loading')}</div>
  }

  return (
    <div>
      <NavigationBar />
    </div>
  )
}

const RetailerRoutes = () => (
  <Routes>
    <Route element={<RetailerLayout />} >
      <Route index element={<RetailerHomePage />} />
      <Route path="flowers" element={<RetailerFlowerPage />} />
    </Route>
  </Routes>
)

const roleLoader = () => {
  if (Authenticator.refresh()) {
    if (Authenticator.role === 'grower') {
      return redirect('/grower')
    } else {
      return redirect('/retailer')
    }
  }
  return null
}

function protectedLoader() {
  if (!Authenticator.isLoggedIn) {
    return redirect("/login")
  }
  return null
}

const router = createBrowserRouter([
  { path: "/", 
    element: <Root />, 
    id: "root",
    loader() {
      Authenticator.refresh()
      return { 
      role: Authenticator.role, isLoggedIn: Authenticator.isLoggedIn, username: Authenticator.username
    }},
    children: [
      { index: true, element: <ProtectedRoute /> },
      { path: "grower", element: <GrowerLayout />, children: [
        { index: true, element: <RetailerHomePage />},
        { path: "flowers", element: <GrowerFlowerPage />},
        { path: "sites", element: <GrowerSitesPage />},
        {
          path: ":siteId",
          children:
          [
            { index: true, element: <GrowerHomePage />},
            { path: "flowers", element: <GrowerFlowerPage />},
            { path: "sites", element: <GrowerSitesPage />},
            { path: "images", element: <GrowerImagesPage />}
          ]
        }      
      ] },
      { path: "retailer", element: <RetailerRoutes /> },
      { path: "user", loader: protectedLoader, element: <UserPage /> },
      { path: "login", loader: roleLoader, element: <LogInPage />,
      },
      { path: "register", loader: roleLoader, element: <RegisterPage /> },
      { path: "terms", element: <TermsPage /> }
    ]
  }
])

export default function App() {
  return <RouterProvider router={router} />
}
