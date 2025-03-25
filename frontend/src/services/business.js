import axios from "axios";
import tokenService from "./token";

const baseUrl = "/api/business";

const create = (newBusiness, user_email) => {
  console.log("testi api", newBusiness, user_email);
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  };
  const request = axios.post(baseUrl, { newBusiness: newBusiness, user_email: user_email }, config);
  return request.then((response) => response.data);
};

export default {
  create,
};
