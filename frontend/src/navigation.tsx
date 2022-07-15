import React, { useEffect, useState } from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Home from './pages/home';
import GitHubLogin from './pages/githublogin';

export function Navigation() {
  return (
    <Routes>
      <Route path="/home" element={<Home />} />
      <Route path="/login" element={<GitHubLogin />} />
    </Routes>
  );
}
