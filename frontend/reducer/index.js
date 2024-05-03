import {combineReducers } from "@reduxjs/toolkit"

import authReducer from "./authSlice"
import contactReducer from "./contactSlice"

const rootReducer = combineReducers({
    auth: authReducer,
    contact: contactReducer
})

export default rootReducer
