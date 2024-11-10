import { getStorage } from "@/utils/storages";
import axios from "axios";

const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_HOST,
  timeout: 30000,
  headers: {
    "Content-Type": "application/json",
  },
});

axiosInstance.interceptors.request.use(
  (config) => {
    const userId = getStorage("user_id", window.localStorage);

    if (userId) {
      config.headers["X-User-ID"] = userId;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => Promise.reject(error)
);

export default axiosInstance;
