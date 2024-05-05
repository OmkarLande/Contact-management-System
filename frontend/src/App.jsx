import React from 'react';
import { Routes, Route } from 'react-router-dom';
import LoginPage from './components/LoginPage';
import RegisterPage from './components/RegisterPage';
import Dashboard from './components/Dashboard';
import AddContactPage from './components/AddContactPage';
import EditContactPage from './components/EditContactPage';

const App = () => {
  return (
    <>
      
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/add-contact" element={<AddContactPage />} />
        <Route path="/edit-contact/:contactid/:userid" element={<EditContactPage />} />
      </Routes>
    </>
  );
};

export default App;
