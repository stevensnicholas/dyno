import React, { useEffect, useState } from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Home from './pages/home';
import UserPage from './pages/userpage';
import Login from './pages/login';

export function Navigation() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/home" element={<Home />} />
      <Route path="/userpage" element={<UserPage />} />
    </Routes>
  );
}
