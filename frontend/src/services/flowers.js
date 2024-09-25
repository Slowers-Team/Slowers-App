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
  const request = axios.post(baseUrl, newFlower)
  return request.then(response => response.data)
}

const remove = id => {
  return axios.delete(`${baseUrl}/${id}`)
}

export default {
  getAll, create, remove
}
