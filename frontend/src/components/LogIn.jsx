import { useState } from 'react';
import userService from '../services/users';

const LogIn = ({ onLogin, setIsLoggedIn, setDefaultRole }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
       e.preventDefault();

    try {
      const response = await userService.login(email, password)

      const data = await response.json();

      if (response.ok) {
        localStorage.setItem("token", data.token);
        localStorage.setItem("role", data.role);
        setIsLoggedIn(true); 
        setDefaultRole(data.role);
        onLogin();
      } else {
        setError("Invalid email or password");
      }
    } catch (err) {
      setError("An error occurred. Please try again.");
      console.log(err)
    }
  };

  return (
    <div>
      <h2>Log In</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="emailInput">Email:</label>
          <input
            id="emailInput"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="passwordInput">Password:</label>
          <input
            id="passwordInput"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button id="loginButton" type="submit">Log In</button>
      </form>
    </div>
  );
}

export default LogIn;
