import axios from 'axios'
const baseUrl = '/api/flowers'

const getAll = () => {
  const config = {
    headers: { Authorization: localStorage.getItem("token") },
  }
  const request = axios.get(baseUrl, config)
  return request.then(response => response.data)
}

const create = newFlower => {
  const config = {
    headers: { Authorization: localStorage.getItem("token") },
  }
  const request = axios.post(baseUrl, newFlower, config)
  return request.then(response => response.data)
}

const remove = id => {
  const config = {
    headers: { Authorization: localStorage.getItem("token") },
  }
  return axios.delete(`${baseUrl}/${id}`, config)
}

export default {
  getAll, create, remove
}
