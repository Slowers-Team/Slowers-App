import axios from 'axios'
const baseUrl = '/api/sites'

const create = newSite =>  {
  const request = axios.post(baseUrl, newSite)
  return request.then(response => response.data)
}

const getSite = id =>  {
  const request = axios.get(`${baseUrl}/${id}`)
  return request.then(response => response.data)
}

const getRootSites = () => {
  const request = axios.get(baseUrl)
  return request.then(response => response.data)
}

const remove = id => {
  return axios.delete(`${baseUrl}/${id}`)
}

export default {
  create
}
