import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import toast from 'react-hot-toast';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';

const EditContactPage = () => {
  const { contactid, userid } = useParams();
  const [contact, setContact] = useState({
    name: '',
    phoneNumber: '',
    email: ''
  });
  const BASE_URL = import.meta.env.VITE_BASE_URL;
  const navigate = useNavigate();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setContact({
      ...contact,
      [name]: value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${BASE_URL}/edit-contact/${contactid}/${userid}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(contact)
      });
      if (!response.ok) {
        throw new Error('Failed to edit contact');
      }
      toast.success('Contact updated successfully');
      navigate('/dashboard');
    } catch (error) {
      console.error('Edit contact error:', error.message);
      toast.error('Failed to edit contact');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <motion.button
        whileHover={{ scale: 1.05 }}
        whileTap={{ scale: 0.95 }}
        onClick={() => navigate('/dashboard')}
        className="fixed top-4 left-4 z-10 bg-gray-200 hover:bg-gray-300 py-2 px-4 rounded-md focus:outline-none focus:bg-gray-300"
      >
        Back
      </motion.button>
      <div className="max-w-md w-full p-8 bg-white shadow-md rounded-md">
        <h2 className="text-3xl font-semibold mb-4 text-center">Edit Contact</h2>
        <motion.form
          onSubmit={handleSubmit}
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="space-y-4 flex flex-col itemcenter justify-center"
        >
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, delay: 0.2 }}
            className="mb-4"
          >
            <label htmlFor="name" className="block text-sm font-medium text-gray-700">
              Name:
            </label>
            <motion.input
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              type="text"
              id="name"
              name="name"
              value={contact.name}
              onChange={handleInputChange}
              className="w-full px-4 py-2 mt-1 rounded-md border border-gray-300 focus:outline-none focus:border-blue-500"
            />
          </motion.div>
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, delay: 0.4 }}
            className="mb-4"
          >
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
              Email:
            </label>
            <motion.input
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              type="email"
              id="email"
              name="email"
              value={contact.email}
              onChange={handleInputChange}
              className="w-full px-4 py-2 mt-1 rounded-md border border-gray-300 focus:outline-none focus:border-blue-500"
            />
          </motion.div>
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, delay: 0.6 }}
            className="mb-4"
          >
            <label htmlFor="phoneNumber" className="block text-sm font-medium text-gray-700">
              Phone Number:
            </label>
            <motion.input
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              type="text"
              id="phoneNumber"
              name="phoneNumber"
              value={contact.phoneNumber}
              onChange={handleInputChange}
              className="w-full px-4 py-2 mt-1 rounded-md border border-gray-300 focus:outline-none focus:border-blue-500"
            />
          </motion.div>
          <motion.button
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
            type="submit"
            className="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:bg-blue-600"
          >
            Update Contact
          </motion.button>
        </motion.form>
      </div>
    </div>
  );
};

export default EditContactPage;
