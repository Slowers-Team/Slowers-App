import userService from '../services/users'
import RegisterForm from '../components/RegisterForm'
import { useTranslation } from 'react-i18next'

const RegisterPage = () => {
    const { t, i18n } = useTranslation()

    const createNewUser = userObject => {
          userService
            .create(userObject)
            .catch(error => {
                alert(t('error.error') + ': ' + error.response.data)
            })  
    }

  return (
    <div>
      <RegisterForm createNewUser={createNewUser} />
    </div>
  )
}

export default RegisterPage
