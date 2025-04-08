import axios from "axios";
import tokenService from "./token";

const baseUrl = "/api/business";

const create = (newBusiness, userEmail) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  };
  const request = axios.post(
    baseUrl,
    { ...newBusiness, userEmail: userEmail },
    config
  );
  return request.then((response) => response.data);
};

const get = () => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  };
  return axios.get(baseUrl, config).then((response) => response.data);
};

// kun kutsutaan tätä:
// businessService.addMembership({user_email, business_id, designation, business_name})
const addMembership = (membership) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
  }
  const url = '/api/membership'
  return axios.post(url, membership, config).then(response => response.data)
}

const deleteMembership = (membership) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
    data: membership
  }
  const url = '/api/membership'
  return axios.delete(url, config).then(response => response.data)
}

const getAllMembers = (businessID) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
  }
  const url = 'api/membership'
  return axios.get(url, businessID, config).then(response => {
    response.data
  })
  .catch(error => {
    console.error('Error fetching business members:', error);
    throw error;
  });
}

export default {
  create,
  get,
  addMembership,
  deleteMembership,
  getAllMembers,
};