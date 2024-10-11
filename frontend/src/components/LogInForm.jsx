import { useState } from 'react'
import userService from '../services/users'

const LogIn = ({ onLogin, setIsLoggedIn }) => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  const handleSubmit = async (e) => {
       e.preventDefault()

    try {
      const response = await userService.login(email, password)

      const data = await response.json()

      if (response.ok) {
        localStorage.setItem("token", data.token)
        setIsLoggedIn(true)
        onLogin()
      } else {
        setError("Invalid email or password")
      }
    } catch (err) {
      setError("An error occurred. Please try again.")
      console.log(err)
    }
  };

  return (
    <div className='text-left'>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div className='form-group mb-4'>
          <label htmlFor="emailInput">Email address</label>
          <input
            type="email"
            id="emailInput"
            value={email}
            placeholder='Enter email'
            onChange={(e) => setEmail(e.target.value)}
            className='form-control'
            required
          />
        </div>
        <div className='form-group mb-4'>
          <label htmlFor="passwordInput">Password</label>
          <input
            type="password"
            id="passwordInput"
            value={password}
            placeholder='Enter password'
            onChange={(e) => setPassword(e.target.value)}
            className='form-control'
            required
          />
        </div>
        <button className='btn btn-primary' id="loginButton" type="submit">Log In</button>
      </form>
    </div>
  )
}

export default LogIn
