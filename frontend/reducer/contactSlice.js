import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  contacts: [],
  loading: false,
};

const contactSlice = createSlice({
  name: "contact",
  initialState: initialState,
  reducers: {
    setContacts(state, action) {
      state.contacts = action.payload;
    },
    setLoading(state, action) {
      state.loading = action.payload;
    },
    addContact(state, action) {
      state.contacts.push(action.payload);
    },
    deleteContact(state, action) {
      state.contacts = state.contacts.filter(contact => contact._id !== action.payload);
    },
    updateContact(state, action) {
      const updatedContactIndex = state.contacts.findIndex(contact => contact._id === action.payload._id);
      if (updatedContactIndex !== -1) {
        state.contacts[updatedContactIndex] = action.payload;
      }
    },
  },
});

export const { setContacts, setLoading, addContact, deleteContact, updateContact } = contactSlice.actions;

export default contactSlice.reducer;
