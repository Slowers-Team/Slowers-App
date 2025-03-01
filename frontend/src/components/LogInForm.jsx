import { useState } from "react";
import userService from "../services/users";
import { useTranslation } from "react-i18next";

const LogIn = ({ onLogin }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setErrorMessage] = useState("");
  const { t, i18n } = useTranslation();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await userService.login(email, password);

      const data = await response.json();

      if (response.ok) {
        onLogin(data);
      } 
    } catch (err) {
      setErrorMessage(t("error.invalidlogininfo"));
    }
  };

  return (
    <div className="text-left">
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div className="input-group mb-3">
          <span className="input-group-text">
            <i className="bi bi-envelope-fill"></i>
          </span>
          <input
            type="email"
            id="emailInput"
            value={email}
            placeholder={t("user.input.email")}
            onChange={(e) => setEmail(e.target.value)}
            className="form-control"
            required
          />
        </div>
        <div className="input-group mb-3">
          <span className="input-group-text">
            <i className="bi bi-lock-fill"></i>
          </span>
          <input
            type="password"
            id="passwordInput"
            value={password}
            placeholder={t("user.input.password")}
            onChange={(e) => setPassword(e.target.value)}
            className="form-control"
            required
          />
        </div>
        <div>
          <button type="submit" id="loginButton" className="custom-button">
            {t("button.login")}
          </button>
        </div>
      </form>
    </div>
  );
};

export default LogIn;
