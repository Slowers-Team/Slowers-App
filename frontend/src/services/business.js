import axios from "axios";
import tokenService from "./token";
import { parseCSV } from "../utils"


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
  const url = "/api/membership"
  return axios.post(url, membership, config).then(response => response.data);
}

const deleteMembership = (membership) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
    data: { email: membership[0] }
  }
  const url = '/api/membership/${businessID}'
  return axios.delete(url, config).then(response => response.data)
}

const getAllMembers = (businessID) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
    responseType: 'json'

  }
  const url = `/api/membership/${businessID}`
  return axios.get(url, config).then(response => {
    console.log("Fetched members:", response.data);
    return parseCSV(response.data);
  }).catch(error => {
    console.error('Error fetching members:', error);
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