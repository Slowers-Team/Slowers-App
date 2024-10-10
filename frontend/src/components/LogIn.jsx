import { useState } from 'react';
import userService from '../services/users';
import { useTranslation } from 'react-i18next';

const LogIn = ({ onLogin, setIsLoggedIn }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const { t, i18n } = useTranslation();

  const handleSubmit = async (e) => {
       e.preventDefault();

    try {
      const response = await userService.login(email, password)

      const data = await response.json();

      if (response.ok) {
        localStorage.setItem("token", data.token);
        setIsLoggedIn(true); 
        onLogin();
      } else {
        setError(t("error.invalidlogininfo"));
      }
    } catch (err) {
      setError(t("error.erroroccured"));
      console.log(err)
    }
  };

  return (
    <div>
      <h2>{t("title.login")}</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="emailInput">{t("user.data.email")}:</label>
          <input
            id="emailInput"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="passwordInput">{t("user.data.password")}:</label>
          <input
            id="passwordInput"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button id="loginButton" type="submit">{t("button.login")}</button>
      </form>
    </div>
  );
}

export default LogIn;
