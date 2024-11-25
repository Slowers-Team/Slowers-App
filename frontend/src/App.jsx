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
import { Routes, Route, createBrowserRouter, RouterProvider, redirect } from 'react-router-dom'
import { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import NavigationBar from './components/NavigationBar'
import { Authenticator } from './Authenticator'

const Root = () => {
  const { t, i18n } = useTranslation()

  useEffect(() => {
    setLanguage()
  }, [])

  const setLanguage = () => {
    const langCookie = document.cookie.split('; ').find(row => row.startsWith('lang='))
    const language = langCookie ? langCookie.split('=')[1] : 'en'
    i18n.changeLanguage(language)
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


// Redirect user to a default role, if they are logged in
const roleLoader = () => {
  if (Authenticator.isLoggedIn) {
    if (Authenticator.role === 'grower') {
      return redirect('/grower')
    } else {
      return redirect('/retailer')
    }
  }
  return null
}

// Redirect user to login-screen, if they are not logged in
function protectedLoader() {
  if (!Authenticator.isLoggedIn) {
    return redirect("/login")
  }
  return null
}

// Redirect user to a default role if logged in, else to login
const rootLoader = () => {
  if (Authenticator.isLoggedIn) {
    if (Authenticator.role === 'grower') {
      return redirect('/grower')
    } else {
      return redirect('/retailer')
    }
  } else {
    return redirect("/login")
  }
}

const router = createBrowserRouter([
  { path: "/", 
    element: <Root />, 
    id: "root",
    loader() {
      Authenticator.refresh()
      return { 
        role: Authenticator.role,
        isLoggedIn: Authenticator.isLoggedIn,
        username: Authenticator.username
    }},
    children: [
      { index: true, loader: rootLoader }, // rootLoader always redirects to another place
      { path: "login", loader: roleLoader, element: <LogInPage />,
        action() { return roleLoader() },
      },
      { path: "register", loader: roleLoader, element: <RegisterPage /> },
      { path: "terms", element: <TermsPage /> },
      { path: "logout", action() { Authenticator.logout(); return null}},
      { path: "*", element: <ProtectedRoute />, loader: protectedLoader,
        children: [
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
          { path: "user", element: <UserPage /> },
          { path: "*", loader() { return redirect("/")} }
        ]
      }
    ]
  }
])

export default function App() {
  return <RouterProvider router={router} />
}
