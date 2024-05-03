import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  signupData: null,
  loading: false,
  session: sessionStorage.getItem("session") ? JSON.parse(sessionStorage.getItem("session")) : null,
};

const authSlice = createSlice({
  name: "auth",
  initialState: initialState,
  reducers: {
    setSignupData(state, action) {
      state.signupData = action.payload;
    },
    setLoading(state, action) {
      state.loading = action.payload;
    },
    setSession(state, action) {
      state.session = action.payload;
      sessionStorage.setItem("session", JSON.stringify(action.payload));
    },
  },
});

export const { setSignupData, setLoading, setSession } = authSlice.actions;

export default authSlice.reducer;
