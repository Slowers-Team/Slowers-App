export const Authenticator = {
  isLoggedIn: false,
  role: '',
  login(token, role) {
    console.log("logging in", token, role)
    localStorage.setItem('token', token)
    localStorage.setItem('role', role)
    Authenticator.isLoggedIn = true
    Authenticator.role = role
  },
  logout() {
    localStorage.removeItem('token')
    localStorage.removeItem('role')
    Authenticator.isLoggedIn = false
    Authenticator.role = ''
  },
  setRole(role){
    Authenticator.role = role
    localStorage.setItem('role', role)
  },
  refresh() {
    Authenticator.role = localStorage.getItem("role")
    if (Authenticator.role) {
      Authenticator.isLoggedIn = true
      return true
    }
    return false
  }
}
  
//   const [isLoggedIn, setIsLoggedIn] = useState(false)
//   const [defaultRole, setDefaultRole] = useState('')
//   const [isLoading, setIsLoading] = useState(true)
//   const { t, i18n } = useTranslation()

//   useEffect(() => {
//     const token = localStorage.getItem('token')
//     const role = localStorage.getItem('role')
//     setIsLoggedIn(!!token)
//     setDefaultRole(role)
//     setIsLoading(false)
//     setLanguage()
//   }, [])

//   const setLanguage = () => {
//     const langCookie = document.cookie.split('; ').find(row => row.startsWith('lang='))
//     const language = langCookie ? langCookie.split('=')[1] : 'en'
//     i18n.changeLanguage(language)
//   }

//   const getDefaultRole = () => {
//     return defaultRole === 'retailer' ? <Navigate replace to="/retailer" /> : <Navigate replace to="/grower" />
//   }

//   const handleLogin = (token, role) => {
//       localStorage.setItem("token", token)
//       localStorage.setItem("role", role)
//       setIsLoggedIn(true)
//       setDefaultRole(role)
//   }

//   const handleLogout = () => {
//     localStorage.removeItem('token')
//     localStorage.removeItem('role')
//     setIsLoggedIn(false)
//     setDefaultRole('')
//   }

