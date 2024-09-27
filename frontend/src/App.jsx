import './App.css';
import RegisterPage from './pages/RegisterPage';
import HomePage from './pages/HomePage';
import LogInPage from './pages/LogInPage';
import SitePage from './pages/SitePage';
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom';
import { useState, useEffect } from 'react';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);  

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
            <Route path="/" element={isLoggedIn ? <HomePage /> : <Navigate replace to="/login" />} />
            <Route path="/login" element={!isLoggedIn ? <LogInPage onLogin={handleLogout} setIsLoggedIn={setIsLoggedIn} /> : <Navigate replace to="/" />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/site" element={<SitePage />} />
          </Routes>
        </div>
      </Router>
    </div>
  );
};

export default App;
