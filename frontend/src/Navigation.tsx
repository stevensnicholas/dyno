import React from 'react';
import { Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Login from './pages/login/Login';
import Dashboard from './pages/Dashboard';
import Error500 from './pages/errors/500';

const Navigation = () => {
  return (
    <Routes>
      <Route path="/home" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/500" element={<Error500 />} />
    </Routes>
  );
};

export default Navigation;
