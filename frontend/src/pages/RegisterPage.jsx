import userService from '../services/users'
import RegisterForm from '../components/RegisterForm'
import { useTranslation } from 'react-i18next'

const RegisterPage = () => {
    const { t, i18n } = useTranslation()

    const createNewUser = userObject => {
          userService
            .create(userObject)
            .catch(error => {
                const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, '')
                alert(t('error.error') + ': ' + (i18n.exists(key) ? t(key) : error.response.data))
            })  
    }

  return (
    <div>
      <RegisterForm createNewUser={createNewUser} />
    </div>
  )
}

export default RegisterPage
