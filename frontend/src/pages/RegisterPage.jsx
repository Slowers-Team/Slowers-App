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
    <section className='vh-100' style = {{ backgroundColor: '#eee' }}>
      <div className='container py-5 h-100'>
        <div className='row d-flex justify-content-center align-items-center h-100'>
          <div className='col-12 col-md-8 col-lg-6 col-xl-5'>
            <div className="card" style={{ borderRadius: '1rem' }}>
              <div className='card-body p-5'>
                <h2 className='mb-5 text-center'>Slowers registration</h2>
                <RegisterForm createNewUser={createNewUser} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}

export default RegisterPage