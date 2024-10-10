import axios from 'axios'
const baseUrl = '/api/register'

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

const get = () => {
  const config = {
    headers: { Authorization: fetchToken() },
  }
  const url = "/api/user"
  return axios.get(url, config).then(response => response.data)
}

const fetchToken = () => {
  return localStorage.getItem("token")
}


export default {
  create, 
  login,
  get,
}
