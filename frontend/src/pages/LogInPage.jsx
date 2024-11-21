import { useNavigate } from 'react-router-dom'
import LogInForm from '../components/LogInForm'
import { useTranslation } from 'react-i18next'

const LogInPage = ({ setIsLoggedIn, setDefaultRole }) => {
  const navigate = useNavigate()
  const { t, i18n } = useTranslation()

  const handleLogin = () => {
    navigate('/')
  }

  return (
    <div style={{ minHeight: 'calc(100vh - 82px)', backgroundColor: '#eee', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <div className='container'>
        <div className='row justify-content-center'>
          <div className='col-12 col-md-8 col-lg-6 col-xl-5'>
            <div className="card" style={{ borderRadius: '1rem' }}>
              <div className='card-body p-5'>
                <h2 className='mb-5 text-center'>{t('title.login')}</h2>
                <LogInForm onLogin={handleLogin} setIsLoggedIn={setIsLoggedIn} setDefaultRole={setDefaultRole}/>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}


export default LogInPage
