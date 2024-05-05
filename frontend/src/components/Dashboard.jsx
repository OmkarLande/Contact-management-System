import React, { useState, useEffect } from 'react';
import toast from 'react-hot-toast';
import { useNavigate } from 'react-router-dom'
import { motion } from 'framer-motion';

const Dashboard = () => {
  const [contacts, setContacts] = useState([]);
  const navigate = useNavigate()
  const userId = localStorage.getItem('session');
  const BASE_URL = import.meta.env.VITE_BASE_URL;

  useEffect(() => {
    const session = localStorage.getItem('session');
    if (!session) {
      navigate('/');
    } else {
    }
    const fetchContacts = async () => {
      try {

        if (!userId) {
          throw new Error('User id not found');
        }
        const response = await fetch(BASE_URL + `/view-contacts/${userId}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
        });

        if (!response.ok) {
          throw new Error('Failed to fetch contacts');
        }

        const data = await response.json();
        setContacts(data || []);
        toast.success('Contacts fetched successfully');
      } catch (error) {
        console.error('Fetch contacts error:', error.message);
        toast.error('Failed to fetch contacts');
      }
    };

    fetchContacts();
  }, [navigate]);

  const handleDeleteContact = async (contactId) => {
    try {
      const response = await fetch(BASE_URL + `/delete-contact/${contactId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error('Failed to delete contact');
      }

      setContacts((prevContacts) => prevContacts.filter((contact) => contact._id !== contactId));
      toast.success('Contact deleted successfully');
    } catch (error) {
      console.error('Delete contact error:', error.message);
      toast.error('Failed to delete contact');
    }
  };

  return (
    <div className="max-w-4xl mx-auto px-4 py-8">
      <h2 className="text-3xl font-bold mb-4 text-center">Contact Management System</h2>
      <div className="flex justify-between items-center mb-6">
        <div>
          <button
            onClick={() => {
              localStorage.removeItem('session');
              navigate('/');
            }}
            className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none"
          >
            Logout
          </button>
        </div>
        <div>
          <button
            onClick={() => navigate('/add-contact')}
            className="bg-green-500 text-white px-2 py-2 rounded-md hover:bg-green-600 focus:outline-none text-center"
          >
            Add Contact
          </button>
        </div>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {contacts.length > 0 ? (
          contacts.map((contact) => (
            <motion.div
              key={contact._id}
              className="bg-white shadow-md rounded-md p-4 cursor-pointer hover:shadow-lg"
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              >
                <img src={contact.profile_image} alt={contact.name} className="w-24 h-24 rounded-full mx-auto" />
              <div className="mb-2">
                <strong>Name:</strong> {contact.name}
              </div>
              <div className="mb-2">
                <strong>Email:</strong> {contact.email}
              </div>
              <div className="mb-2">
                <strong>Phone Number:</strong> {contact.phoneNumber}
              </div>
              
              <div className='flex justify-around items-center'>
                <button 
                  className='bg-gray-500 text-white px-2 py-2 rounded-md hover:bg-gray-600 focus:outline-none text-center' 
                  onClick={() => navigate(`/edit-contact/${contact._id}/${userId}`)}
                >
                  Edit
                </button>
                <button 
                  className='bg-red-500 text-white px-2 py-2 rounded-md hover:bg-red-600 focus:outline-none text-center'
                  onClick={() => handleDeleteContact(contact._id)}
                >
                    Delete
                </button>
              </div>
            </motion.div>
          ))
        ) : (
          <p className="text-xl text-gray-500">No contacts found</p>
        )}
      </div>
    </div>
  );
};

export default Dashboard;
