import { useEffect, useState } from 'react';
import './App.css';
import { AppClient } from './client';
import { Loading } from './components/Loading/Loading';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Link, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Login from './pages/login/Login';
import Dashboard from './pages/Dashboard';

export interface PageProps {
  client: AppClient;
}

export function App() {
  const [client, setClient] = useState<AppClient | undefined>();
  const [clientId, setClientId] = useState<string>('');
  const [loggedIn, setLoggedIn] = useState<boolean>(false);

  useEffect(() => {
    fetch('settings.json')
      .then((res) => res.json())
      .then((settings) => {
        setClientId(settings.client_id);
        setClient(new AppClient({ BASE: settings.backend }));
      });
  }, []);

  useEffect(() => {
    const code = window.location.search.substring(1).split('=')[2];
    if (client && code !== undefined) {
      client.default.endpointsAuthentication(code).then((res) => {
        window.localStorage.setItem('token', res.token);
        setLoggedIn(true);
      });
    }
  }, [client]);
  return (
    <div
      id="main"
      style={{
        overflowX: 'hidden',
        width: '100%',
        height: '100%',
        backgroundPosition: 'center',
        backgroundRepeat: 'no-repeat',
        backgroundSize: 'cover',
      }}
    >
      <ToastContainer />

      <div className="App">
        <h1>Go Lambda Skeleton</h1>
        {client ? (
          loggedIn ? (
            <>
              <div>
                <nav>
                  <ul>
                    <li>
                      <Link to="/">Home</Link>
                    </li>
                    <li>
                      <Link to="/login">Login</Link>
                    </li>
                    <li>
                      <Link to="/dashboard">Dashboard</Link>
                    </li>
                  </ul>
                </nav>
              </div>

              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login clientID={clientId} />} />
                <Route path="/dashboard" element={<Dashboard />} />
              </Routes>
            </>
          ) : (
            <>
              <div>
                <nav>
                  <ul>
                    <li>
                      <Link to="/">Home</Link>
                    </li>
                    <li>
                      <Link to="/login">Login</Link>
                    </li>
                  </ul>
                </nav>
              </div>

              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login clientID={clientId} />} />
              </Routes>
            </>
          )
        ) : (
          <Loading />
        )}
      </div>
    </div>
  );
}
