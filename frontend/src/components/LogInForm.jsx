import { useTranslation } from 'react-i18next'
import { Form, useActionData } from 'react-router-dom' 

const LogIn = () => {
  const { t, i18n } = useTranslation()

  const errors = useActionData()
  let error = null

  if (errors?.invalidLogin) {
    error = t("error.invalidlogininfo")
  } else if (errors?.error) {
    error = t("error.erroroccured")
  }

  return (
    <div className='text-left'>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <Form action="/login" method="post">
        <div className='form-group mb-4'>
          <label htmlFor="emailInput">{t("user.data.email")}</label>
          <input
            type="email"
            id="emailInput"
            name="email"
            placeholder={t("user.input.email")}
            className='form-control'
            required
          />
        </div>
        <div className='form-group mb-4'>
          <label htmlFor="passwordInput">{t("user.data.password")}</label>
          <input
            type="password"
            id="passwordInput"
            name="password"
            placeholder={t("user.input.password")}
            className='form-control'
            required
          />
        </div>
        <div>
          <button type="submit" id="loginButton" className='btn btn-primary' >{t("button.login")}</button>
        </div>
      </Form>
    </div>
  )
}

export default LogIn
