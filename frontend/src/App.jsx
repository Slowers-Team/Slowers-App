import './App.css';
import RegisterPage from './pages/RegisterPage';
import HomePage from './pages/HomePage';
import LogInPage from './pages/LogInPage';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useState, useEffect } from 'react';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);  
  console.log(isLoggedIn)

  useEffect(() => {
    const token = localStorage.getItem('token');
    setIsLoggedIn(!!token);  
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsLoggedIn(false);
  };

  const padding = {
    padding: 5,
  };

  return (
    <div>
      <Router>
        <div>
          <nav>
            <Link style={padding} to="/">Home</Link>
            {!isLoggedIn && <Link style={padding} to="/register">Register</Link>}
            {!isLoggedIn && <Link style={padding} to="/login">Login</Link>}
            {isLoggedIn && <Link onClick={handleLogout}>Logout</Link>}
          </nav>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/login" element={<LogInPage setIsLoggedIn={setIsLoggedIn} />} />
            <Route path="/register" element={<RegisterPage />} />
          </Routes>
        </div>
      </Router>
    </div>
  );
};

export default App;
