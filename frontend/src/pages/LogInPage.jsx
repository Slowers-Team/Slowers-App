import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import LogIn from '../components/LogIn'

const LogInPage = () => {
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleLogin = async (email, password) => {
    try {
      const response = await fetch('/api/loginReq', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem('token', data.token);
        navigate('/');  // Oletus että kirjautuneille on oma näkymä
      } else {
        setError('Invalid email or password');
      }
    } catch (err) {
      setError('An error occurred during login');
    }
  };

  return (
    <div>
      <h1>Login</h1>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <LogIn onLogin={handleLogin} />
    </div>
  );
};

export default LogInPage;
