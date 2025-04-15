// Authenticator provides persistent login information for the app.
// It stores its state in local storage to allow reloading of page.
export const Authenticator = {
  isLoggedIn: false,
  role: "",
  designation: "",
  username: "",
  email: "",
  businessType: "",
  login({ token, role, username, email, designation, businessType }) {
    localStorage.setItem("token", token);
    localStorage.setItem("role", role);
    localStorage.setItem("designation", designation);
    localStorage.setItem("businessType", businessType);
    localStorage.setItem("username", username);
    localStorage.setItem("email", email);
    Authenticator.isLoggedIn = true;
    Authenticator.role = role;
    Authenticator.designation = designation;
    Authenticator.businessType = businessType
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
    localStorage.removeItem("businessType");
    Authenticator.isLoggedIn = false;
    Authenticator.role = "";
    Authenticator.designation = "";
    Authenticator.username = "";
    Authenticator.email = "";
    Authenticator.businessType = "";
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
  setBusinessType(businessType) {
    Authenticator.businessType = businessType;
    localStorage.setItem("businessType", businessType);
  },
  refresh() {
    Authenticator.role = localStorage.getItem("role");
    Authenticator.designation = localStorage.getItem("designation");
    Authenticator.businessType = localStorage.getItem("businessType");
    if (Authenticator.role || Authenticator.designation || Authenticator.businessType ) {
      Authenticator.isLoggedIn = true;
      Authenticator.username = localStorage.getItem("username");
      return true;
    }
    return false;
  },
};
