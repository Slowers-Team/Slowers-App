import axios from 'axios'
const baseUrl = '/api/register'
const logInUrl = '/api/login'


const create = newUser =>  {
    const request = axios.post(baseUrl, newUser)
    return request.then(response => response.data)
  }

  const login = (email, password)  => {
    return fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

  }

  export default {
    create, 
    login,
  }
