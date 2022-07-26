import React, { useEffect, useState } from 'react';
import { Route, Routes } from 'react-router-dom';
import { AppClient } from './client';
import Home from './pages/Home';
import Login from './pages/login/Login';
import Dashboard from './pages/Dashboard';

export interface PageProps {
  client: AppClient;
}

const Navigation = () => {
  const [client, setClient] = useState<AppClient | undefined>();
  const [clientId, setClientId] = useState<string>('');

  useEffect(() => {
    fetch('settings.json')
      .then((res) => res.json())
      .then((settings) => {
        setClient(new AppClient({ BASE: settings.backend }));
        setClientId(settings.client_id);
      });
  }, []);
  return (
    <Routes>
      <Route path="/home" element={<Home />} />
      <Route path="/login" element={<Login clientID={clientId} />} />
      <Route path="/dashboard" element={<Dashboard />} />
    </Routes>
  );
};

export default Navigation;
