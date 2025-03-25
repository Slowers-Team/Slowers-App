// Authenticator provides persistent login information for the app.
// It stores its state in local storage to allow reloading of page.
export const Authenticator = {
  isLoggedIn: false,
  role: "",
  username: "",
  email: "",
  login({ token, role, username, email }) {
    localStorage.setItem("token", token);
    localStorage.setItem("role", role);
    localStorage.setItem("username", username);
    localStorage.setItem("email", email);
    Authenticator.isLoggedIn = true;
    Authenticator.role = role;
    Authenticator.username = username;
    Authenticator.email = email;
    return null;
  },
  logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("role");
    localStorage.removeItem("username");
    localStorage.removeItem("email");
    Authenticator.isLoggedIn = false;
    Authenticator.role = "";
    Authenticator.username = "";
    Authenticator.email = "";
    return null;
  },
  setRole(role) {
    Authenticator.role = role;
    localStorage.setItem("role", role);
  },
  refresh() {
    Authenticator.role = localStorage.getItem("role");
    if (Authenticator.role) {
      Authenticator.isLoggedIn = true;
      Authenticator.username = localStorage.getItem("username");
      return true;
    }
    return false;
  },
};
