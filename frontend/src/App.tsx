import { useEffect, useState } from 'react';
import './App.css';
import { AppClient } from './client';
import { Loading } from './components/Loading/Loading';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Link, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Login from './pages/login/Login';
import Dashboard from './pages/dashboard/Dashboard';
import TestResult from './pages/dashboard/TestResult';

export interface PageProps {
  client: AppClient;
}

export function App() {
  const [client, setClient] = useState<AppClient | undefined>();
  const [clientId, setClientId] = useState<string>('');
  const [loggedIn, setLoggedIn] = useState<boolean>(false);
  const [token, setToken] = useState<string>('');

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
        window.localStorage.setItem('token', res.jwt);
        setToken(res.jwt);
      });
    }
  }, [client]);

  useEffect(() => {
    if (token) {
      setLoggedIn(true);
    }
  }, [token]);

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
                <Route path="/" element={<Home loggedIn={loggedIn} />} />
                <Route path="/login" element={<Login clientID={clientId} />} />
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/testresult/:id" element={<TestResult />} />
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
                <Route path="/" element={<Home loggedIn={loggedIn} />} />
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
