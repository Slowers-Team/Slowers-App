import axios from 'axios'
const baseUrl = '/api/flowers'
import tokenService from './token'

const getAll = () => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  }
  const request = axios.get(baseUrl, config)
  return request.then(response => response.data)
}
const getFlowersForSite = (siteId) => {
  return axios.get(`${API_URL}/site/${siteId}`).then(response => response.data);
};


const getUserFlowers = () => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  }
  const request = axios.get(`${baseUrl}/user`, config)
  return request.then(response => response.data)
}

const create = newFlower => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  }
  console.log(newFlower)
  const request = axios.post(baseUrl, newFlower, config)
  return request.then(response => response.data)
}

const remove = id => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  }
  return axios.delete(`${baseUrl}/${id}`, config)
}

export default {
  getAll,
  create,
  remove,
  getUserFlowers,
  getFlowersForSite
}
