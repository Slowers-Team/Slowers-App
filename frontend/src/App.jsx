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
import { createBrowserRouter, RouterProvider, redirect } from 'react-router-dom'
import { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import NavigationBar from './components/NavigationBar'
import { Authenticator } from './Authenticator'
import userService from './services/users'

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
const rootRedirect = () => {
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

const rootLoader = () => {
  Authenticator.refresh() // try to fetch login info from local storage 
  return { 
    role: Authenticator.role,
    isLoggedIn: Authenticator.isLoggedIn,
    username: Authenticator.username
  } 
}

const loginAction = async ({ request }) => {
  const errors = {}
  try {
    const formData = await request.formData()
    const email = formData.get("email")
    const password = formData.get("password")

    const response = await userService.login(email, password)

    if (response.ok) {
      Authenticator.login({ ... await response.json() })
      return redirect("/")
    } else {
      errors.invalidLogin = true
    }
  } catch (err) {
    console.error(err)
    errors.error = err
  }
  return errors
}

const registerAction = async ({ request }) => {
  const loginInfo = Object.fromEntries(await request.formData())

  Authenticator.login(loginInfo)

  return redirect("/")
}

const router = createBrowserRouter([
  { 
    path: "/", 
    element: <Root />, 
    id: "root",
    loader: rootLoader,
    children: [
      { 
        index: true,
        loader: rootRedirect // rootRedirect always redirects to another place
      }, 
      { path: "login", 
        loader: roleLoader, 
        element: <LogInPage />,
        action: loginAction,
      },
      { 
        path: "register", 
        loader: roleLoader, 
        action: registerAction,
        element: <RegisterPage /> 
      },
      { path: "terms", element: <TermsPage /> },
      { path: "logout",
        action() { return Authenticator.logout() } // POST /logout -> Authenticator.logout()
      },
      { path: "*", 
        loader: protectedLoader, 
        children: 
        [
          { 
            path: "grower",
            element: <GrowerLayout />,
            children: 
            [
              { index: true,     element: <GrowerHomePage />},
              { path: "flowers", element: <GrowerFlowerPage />},
              { path: "sites",   element: <GrowerSitesPage />},
              { path: ":siteId", children: 
                [
                  { index: true,     element: <GrowerHomePage />},
                  { path: "flowers", element: <GrowerFlowerPage />},
                  { path: "sites",   element: <GrowerSitesPage />},
                  { path: "images",  element: <GrowerImagesPage />}
                ] 
              } 
            ] 
          },
          { 
            path: "retailer",
            element: <RetailerLayout />,
            children: 
            [
              { index: true,     element: <RetailerHomePage />},
              { path: "flowers", element: <RetailerFlowerPage />}
            ] },
          { path: "user", element: <UserPage /> },
          { path: "*", loader() { return redirect("/")} } // redirect undefined paths to home
        ]
      }
    ]
  }
])

export default function App() {
  return <RouterProvider router={router} />
}
