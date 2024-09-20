import userService from '../services/users'
import RegisterForm from '../components/RegisterForm'

const RegisterPage = () => {

    const createNewUser = userObject => {
          userService
            .create(userObject)
            .catch(error => {
                alert('Error: ' + error.response.data)
            })  
    }

  return (
    <div>
      <RegisterForm createNewUser={createNewUser} />
    </div>
  )
}

export default RegisterPage