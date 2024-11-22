import { useState } from 'react'
import userService from '../services/users'
import { useTranslation } from 'react-i18next'

const LogIn = ({ onLogin }) => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const { t, i18n } = useTranslation()

  const handleSubmit = async (e) => {
       e.preventDefault()

    try {
      const response = await userService.login(email, password)

      const data = await response.json()

      if (response.ok) {
        onLogin(data.token, data.role, data.username)
      } else {
        setError(t("error.invalidlogininfo"))
      }
    } catch (err) {
      setError(t("error.erroroccured"))
      console.log(err)
    }
  };

  return (
    <div className='text-left'>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div className='form-group mb-4'>
          <label htmlFor="emailInput">{t("user.data.email")}</label>
          <input
            type="email"
            id="emailInput"
            value={email}
            placeholder={t("user.input.email")}
            onChange={(e) => setEmail(e.target.value)}
            className='form-control'
            required
          />
        </div>
        <div className='form-group mb-4'>
          <label htmlFor="passwordInput">{t("user.data.password")}</label>
          <input
            type="password"
            id="passwordInput"
            value={password}
            placeholder={t("user.input.password")}
            onChange={(e) => setPassword(e.target.value)}
            className='form-control'
            required
          />
        </div>
        <div>
          <button type="submit" id="loginButton" className='btn btn-primary' >{t("button.login")}</button>
        </div>
      </form>
    </div>
  )
}

export default LogIn
