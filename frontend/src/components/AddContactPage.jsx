import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import toast from 'react-hot-toast';
import { motion } from 'framer-motion';

const AddContactPage = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phoneNumber, setPhoneNumber] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const session = localStorage.getItem('session');
    if (!session) {
      navigate('/');
    } else {
      console.log(session)
    }
  }, [navigate]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const userId = localStorage.getItem('session');
      if (!userId) {
        throw new Error('User id not found');
      }
      const BASE_URL = import.meta.env.VITE_BASE_URL;
      const response = await fetch(BASE_URL + '/add-contact', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, email, phoneNumber, userId }),
      });
      console.log(response)

      if (!response.ok) {
        const data = await response.json();
        throw new Error(data.message || 'Failed to add contact');
      }

      navigate('/dashboard');
      toast.success('Contact added successfully');
    } catch (error) {
      console.error('Add contact error:', error.message);
      toast.error(error.message);
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
        <h2 className="text-3xl font-semibold mb-4 text-center">Add Contact</h2>
        <motion.form
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          onSubmit={handleSubmit}
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
              value={name}
              onChange={(e) => setName(e.target.value)}
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
              value={email}
              onChange={(e) => setEmail(e.target.value)}
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
              value={phoneNumber}
              onChange={(e) => setPhoneNumber(e.target.value)}
              className="w-full px-4 py-2 mt-1 rounded-md border border-gray-300 focus:outline-none focus:border-blue-500"
            />
          </motion.div>
          <motion.button
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
            type="submit"
            className="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:bg-blue-600"
          >
            Add Contact
          </motion.button>
        </motion.form>
      </div>
    </div>
  );
};

export default AddContactPage;
