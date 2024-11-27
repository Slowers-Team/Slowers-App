import RegisterPage from './pages/RegisterPage'
import TermsPage from './pages/TermsPage'
import LogInPage from './pages/LogInPage'
import UserPage from './pages/UserPage'
import RetailerHomePage from './pages/RetailerHomePage'
import RetailerFlowerPage from './pages/RetailerFlowerPage'
import RetailerLayout from './layouts/RetailerLayout'
import { createBrowserRouter, RouterProvider, redirect } from 'react-router-dom'
import { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import NavigationBar from './components/NavigationBar'
import { Authenticator } from './Authenticator'
import growerRoutes from './routes/grower'

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
      Authenticator.refresh() // try to fetch login info from local storage 
      return { 
        role: Authenticator.role,
        isLoggedIn: Authenticator.isLoggedIn,
        username: Authenticator.username
    }},
    children: [
      { index: true,   loader: rootLoader }, // rootLoader always redirects to another place
      { path: "login", loader: roleLoader, element: <LogInPage />,
        action() { return redirect("/") }, // PUT /login -> redirect to homepage
      },
      { path: "register", loader: roleLoader, element: <RegisterPage /> },
      { path: "terms", element: <TermsPage /> },
      { path: "logout",
        action() { return Authenticator.logout() } // PUT /logout -> Authenticator.logout()
      },
      { path: "*", loader: protectedLoader, children: 
        [
          growerRoutes,
          { path: "retailer",  element: <RetailerLayout />, children: 
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
