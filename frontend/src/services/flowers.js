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

const getFlowersBySite = (id=null) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  };
  const url = id ? `/api/sites/${id}/flowers` : baseUrl;

  return axios.get(url, config)
    .then(response => {
      console.log("Fetched flowers:", response.data); 
      return response.data;
    })
    .catch(error => {
      console.error('Error fetching site flowers:', error);
      throw error;
    });
};

const toggleVisibility = (id) => {
  const config = {
    method: 'post',
    url: `${baseUrl}/${id}/visibility`,
    headers: { Authorization: tokenService.fetchToken() },
  }

  return axios(config)
    .then(response => {
      console.log("Visibility of", id, "set to", response.data)
      return response.data
    })
    .catch(error => {
      console.error("Error setting visibility of flower", id,":",error.response)
      throw error.response.data
    })
  
}

export default {
  getAll,
  create,
  remove,
  getUserFlowers,
  getFlowersBySite,
  toggleVisibility
}
