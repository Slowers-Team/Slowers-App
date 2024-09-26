import { useNavigate } from 'react-router-dom';
import LogIn from '../components/LogIn';

const LogInPage = ({ setIsLoggedIn }) => {
  const navigate = useNavigate();

  const handleLogin = () => {
    navigate('/');  
  };

  return (
    <div>
      <LogIn onLogin={handleLogin} setIsLoggedIn={setIsLoggedIn} />
    </div>
  );
};

export default LogInPage;
