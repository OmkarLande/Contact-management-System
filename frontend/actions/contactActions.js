import axios from "axios"
import { toast } from "react-hot-toast"
import { setLoading } from "../reducer/authSlice"
import { setContacts, setLoading, addContact, deleteContact, updateContact } from "../reducer/contactSlice"

const BASE_URL = process.env.VITE_BASE_URL

const ADD_CONTACT_API = BASE_URL + "/add-contact"
const GET_CONTACTS_API = BASE_URL + "/get-contacts"
const UPDATE_CONTACT_API = BASE_URL + "/update-contact/{contact_id}"
const DELETE_CONTACT_API = BASE_URL + "/delete-contact/{contact_id}"

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

export const addContact = (name, email, phone) => async (dispatch) => {
  const toastId = toast.loading("Loading...")
  dispatch(setLoading(true))
  try {
    const response = await apiConnector("POST", ADD_CONTACT_API, {
      name,
      email,
      phone,
    }, { Authorization: `Bearer ${sessionStorage.getItem("session")}`
})

    if (!response.data.success) {
      throw new Error(response.data.message)
    }

    dispatch(setLoading(false))
    dispatch(addContact(response.data.contact))
    toast.success(response.data.message)
  } catch (error) {
    console.log(error.response.data.message)
    toast.error("Add Contact Failed")
  }
  dispatch(setLoading(false))
  toast.dismiss(toastId)
}

export const getContacts = () => async (dispatch) => {
  const toastId = toast.loading("Loading...")
  dispatch(setLoading(true))
  try {
    const response = await apiConnector("GET", GET_CONTACTS_API, null, { Authorization: `Bearer ${sessionStorage.getItem("session")}` })

    if (!response.data.success) {
      throw new Error(response.data.message)
    }

    dispatch(setLoading(false))
    dispatch(setContacts(response.data.contacts))
    toast.success(response.data.message)
  } catch (error) {
    console.log(error.response.data.message)
    toast.error("Get Contacts Failed")
  }
  dispatch(setLoading(false))
  toast.dismiss(toastId)
}

export const updateContact = (contact_id, name, email, phone) => async (dispatch) => {
  const toastId = toast.loading("Loading...")
  dispatch(setLoading(true))
  try {
    const response = await apiConnector("PUT", UPDATE_CONTACT_API.replace("{contact_id}", contact_id), {
      name,
      email,
      phone,
    }, { Authorization: `Bearer ${sessionStorage.getItem("session")}` })

    if (!response.data.success) {
      throw new Error(response.data.message)
    }

    dispatch(setLoading(false))
    dispatch(updateContact(response.data.contact))
    toast.success(response.data.message)
  } catch (error) {
    console.log(error.response.data.message)
    toast.error("Update Contact Failed")
  }
  dispatch(setLoading(false))
  toast.dismiss(toastId)
}

export const deleteContact = (contact_id) => async (dispatch) => {
    const toastId = toast.loading("Loading...")
    dispatch(setLoading(true))
    try {
        const response = await apiConnector("DELETE", DELETE_CONTACT_API.replace("{contact_id}", contact_id), null, { Authorization: `Bearer ${sessionStorage.getItem("session")}` })
    
        if (!response.data.success) {
        throw new Error(response.data.message)
        }
    
        dispatch(setLoading(false))
        dispatch(deleteContact(contact_id))
        toast.success(response.data.message)
    } catch (error) {
        console.log(error.response.data.message)
        toast.error("Delete Contact Failed")
    }
    dispatch(setLoading(false))
    toast.dismiss(toastId)
}