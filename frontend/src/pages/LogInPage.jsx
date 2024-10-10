import { useNavigate } from 'react-router-dom';
import LogIn from '../components/LogIn';

const LogInPage = ({ setIsLoggedIn }) => {
  const navigate = useNavigate();

  const handleLogin = () => {
    navigate('/');  
  };

  return (
    <section className='vh-100' style = {{ backgroundColor: '#eee' }}>
      <div className='container py-5 h-100'>
        <div className='row d-flex justify-content-center align-items-center h-100'>
          <div className='col-12 col-md-8 col-lg-6 col-xl-5'>
            <div className="card shadow-2-strong" style={{ borderRadius: '1rem' }} data-mdb-ripple-color="light">
              <div className='card-body p-5 text-center'>
                <LogIn onLogin={handleLogin} setIsLoggedIn={setIsLoggedIn} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default LogInPage;
