import { useNavigate } from 'react-router-dom';
import LogIn from '../components/LogIn';

const LogInPage = ({ setIsLoggedIn, setDefaultRole }) => {
  const navigate = useNavigate();

  const handleLogin = () => {
    navigate('/');  
  };

  return (
    <div>
      <LogIn onLogin={handleLogin} setIsLoggedIn={setIsLoggedIn} setDefaultRole={setDefaultRole}/>
    </div>
  );
};

export default LogInPage;
