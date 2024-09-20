import userService from '../services/users'
import RegisterForm from '../components/RegisterForm'

const RegisterPage = () => {

    const createNewUser = userObject => {
        if (window.confirm(`Are you sure you want to register a new user with email ${userObject.email}?`)) {
          userService
            .create(userObject)
            .catch(error => {
                console.log(error)
                alert('Registering failed')
            })  
        }
    }

  return (
    <div>
      <RegisterForm createNewUser={createNewUser} />
    </div>
  )
}

export default RegisterPage