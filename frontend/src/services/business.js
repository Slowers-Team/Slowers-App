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
  const url = "/api/business";
  return axios.get(url, config).then((response) => response.data);
};

export default {
  create,
  get,
};
