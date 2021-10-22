import axios, { AxiosRequestConfig } from "axios";
import { LoginInputs } from "../layouts/login";
import { AuthToken } from "./auth_token";
import { catchAxiosError } from "./error";

export const postLogin = async (inputs) => {
  const data = new URLSearchParams(inputs);
  const res = await post("/user/signin", data).catch(catchAxiosError);
  if (res.error) {
    return res.error;
  }
  if (res.data && res.data.token) {
    alert(`this is my token: (${res.data.token})`);
    return;
  }
  return "Something unexpected happened!";
};

// a base configuration we can extend from
const baseConfig = {
  baseURL: "http://localhost:8080",
};

const post = (url, data) => {
  return axios.post(url, data, baseConfig).catch(catchAxiosError);
};
