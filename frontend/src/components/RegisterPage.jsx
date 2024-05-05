import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-hot-toast';
import { motion } from 'framer-motion';

const RegisterPage = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();
  const BASE_URL = import.meta.env.VITE_BASE_URL;

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(BASE_URL + '/register-user', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, email, password }),
      });

      if (!response.ok) {
        const data = await response.json();
        throw new Error(data.message || 'Registration failed');
      }

      navigate('/');
      toast.success('Registration successful. You can now login.', { appearance: 'success' });
    } catch (error) {
      console.error('Registration error:', error.message);
      toast.error(error.message);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="max-w-md w-full p-8 bg-white shadow-md rounded-md">
        <h2 className="text-3xl font-semibold mb-4 text-center">Register</h2>
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
            <label htmlFor="username" className="block text-sm font-medium text-gray-700">
              Username:
            </label>
            <motion.input
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              type="text"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
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
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Password:
            </label>
            <motion.input
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-4 py-2 mt-1 rounded-md border border-gray-300 focus:outline-none focus:border-blue-500"
            />
          </motion.div>
          <motion.button
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
            type="submit"
            className="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:bg-blue-600"
          >
            Register
          </motion.button>
        </motion.form>
        <p className="mt-4 text-center">
          Already have an account?{' '}
          <motion.button
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
            onClick={() => navigate('/')}
            className="text-blue-500 hover:underline focus:outline-none focus:underline"
          >
            Login
          </motion.button>
        </p>
      </div>
    </div>
  );
};

export default RegisterPage;
