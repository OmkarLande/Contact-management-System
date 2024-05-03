import axios from "axios"
import { toast } from "react-hot-toast"
import { setLoading, setSession } from "../reducer/authSlice"


const BASE_URL = process.env.VITE_BASE_URL

const REGISTER_API = BASE_URL + "/register-user"
const LOGIN_API = BASE_URL + "/login-user"
const LOGOUT_API = BASE_URL + "/logout-user"

export const axiosInstance = axios.create({});

export const apiConnector = (method, url, bodyData, headers, params) => {
  return axiosInstance({
    method: `${method}`,
    url: `${url}`,
    data: bodyData ? bodyData : null,
    headers: headers ? headers : null,
    params: params ? params : null,
  });
}

export const registerUser = (username, email, password, navigate) => async (dispatch) => {
  const toastId = toast.loading("Loading...")
  dispatch(setLoading(true))
  try {
    const response = await apiConnector("POST", REGISTER_API, {
      username,
      email,
      password,
    })

    if (!response.data.success) {
      throw new Error(response.data.message)
    }

    dispatch(setLoading(false))
    toast.success(response.data.message)
    navigate("/login-user")
  } catch (error) {
    console.log(error.response.data.message)
    toast.error("Signup Failed")
  }
  dispatch(setLoading(false))
  toast.dismiss(toastId)
}

export const loginUser = (username, password, navigate) => async (dispatch) => {
  const toastId = toast.loading("Loading...")
  dispatch(setLoading(true))
  try {
    const response = await apiConnector("POST", LOGIN_API, {
      username,
      password,
    })

    if (!response.data.success) {
      throw new Error(response.data.message)
    }

    dispatch(setSession(response.data.session))
    dispatch(setLoading(false))
    toast.success(response.data.message)
    navigate("/dashboard")
  } catch (error) {
    console.log(error.response.data.message)
    toast.error("Login Failed")
  }
  dispatch(setLoading(false))
  toast.dismiss(toastId)
}

export const logoutUser = (navigate) => async (dispatch) => {
  const toastId = toast.loading("Loading...")
  dispatch(setLoading(true))
  try {
    const response = await apiConnector("POST", LOGOUT_API)

    if (!response.data.success) {
      throw new Error(response.data.message)
    }

    dispatch(setSession(null))
    dispatch(setLoading(false))
    toast.success(response.data.message)
    navigate("/login-user")
  } catch (error) {
    console.log(error.response.data.message)
    toast.error("Logout Failed")
  }
  dispatch(setLoading(false))
  toast.dismiss(toastId)
}