import axios from "axios";
import tokenService from "./token";
import { parseCSV } from "../utils"


const create = (newBusiness, userEmail) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  };
  const url = '/api/business'
  const request = axios.post(
    url,
    { ...newBusiness, userEmail: userEmail },
    config
  );
  return request.then((response) => response.data);
};

const get = () => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  };
  const url = '/api/business'
  return axios.get(url, config).then((response) => response.data);
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
// businessService.editMember({user_email, business_id, newdesignation})
const editMember = (member) => {
  const config = {
    headers: { Authorization: tokenService.fetchToken(),
    'Content-Type': 'application/json'},
  }
  const url = '/api/membership/edit'
  return axios.post(url, member, config).then(response => response.data)
}

export default {
  create,
  get,
  addMembership,
  deleteMembership,
  getAllMembers,
  editMember,
};