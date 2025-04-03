// Authenticator provides persistent login information for the app.
// It stores its state in local storage to allow reloading of page.
export const Authenticator = {
  isLoggedIn: false,
  role: "",
  designation: "",
  username: "",
  email: "",
  login({ token, role, username, email, designation }) {
    localStorage.setItem("token", token);
    localStorage.setItem("role", role);
    localStorage.setItem("designation", designation);
    console.log("meneekö", designation)
    localStorage.setItem("username", username);
    localStorage.setItem("email", email);
    Authenticator.isLoggedIn = true;
    Authenticator.role = role;
    Authenticator.designation = designation;
    Authenticator.username = username;
    Authenticator.email = email;
    return null;
  },
  logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("role");
    localStorage.removeItem("designation");
    localStorage.removeItem("username");
    localStorage.removeItem("email");
    Authenticator.isLoggedIn = false;
    Authenticator.role = "";
    Authenticator.designation = "";
    Authenticator.username = "";
    Authenticator.email = "";
    return null;
  },
  setRole(role) {
    Authenticator.role = role;
    localStorage.setItem("role", role);
  },
  setDesignation(designation) {
    Authenticator.designation = designation;
    localStorage.setItem("designation", designation);
  },
  refresh() {
    Authenticator.role = localStorage.getItem("role");
    Authenticator.designation = localStorage.getItem("designation");
    if (Authenticator.role || Authenticator.designation) {
      Authenticator.isLoggedIn = true;
      Authenticator.username = localStorage.getItem("username");
      return true;
    }
    return false;
  },
};
