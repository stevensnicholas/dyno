import React, { useEffect, useState } from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Home from './pages/Home';
import UserPage from './pages/UserPage';
import Login from './pages/Login';

export function Navigation() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/home" element={<Home />} />
      <Route path="/userpage" element={<UserPage />} />
    </Routes>
  );
}