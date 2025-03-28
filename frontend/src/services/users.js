import axios from 'axios'
import tokenService from './token'

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
    headers: { Authorization: tokenService.fetchToken() },
  }
  const url = "/api/user"
  return axios.get(url, config).then(response => response.data)
}

const setRole = (role) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
  }
  const url = '/api/user/role'
  return axios.post(url, role, config).then(response => response.data)
}

const setDesignation = (designation) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
  }
  const url = '/api/user/designation'
  return axios.post(url, designation, config).then(response => response.data)
}

export default {
  create, 
  login,
  get,
  setRole,
  setDesignation,
}
