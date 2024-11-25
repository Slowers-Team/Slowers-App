// Authenticator provides persistent login information for the app.
// It stores its state in local storage to allow reloading of page.
export const Authenticator = {
  isLoggedIn: false,
  role: '',
  username: '',
  login({ token, role, username} ) {
    localStorage.setItem('token', token)
    localStorage.setItem('role', role)
    localStorage.setItem('username', username)
    Authenticator.isLoggedIn = true
    Authenticator.role = role
    Authenticator.username = username
    return null
  },
  logout() {
    localStorage.removeItem('token')
    localStorage.removeItem('role')
    localStorage.removeItem('username')
    Authenticator.isLoggedIn = false
    Authenticator.role = ''
    Authenticator.username = ''
    return null
  },
  setRole(role){
    Authenticator.role = role
    localStorage.setItem('role', role)
  },
  refresh() {
    Authenticator.role = localStorage.getItem("role")
    if (Authenticator.role) {
      Authenticator.isLoggedIn = true
      Authenticator.username = localStorage.getItem("username")
      return true
    }
    return false
  }
}
