import { useEffect, useState } from 'react';
import './App.css';
import { AppClient } from './client';
import { Loading } from './components/Loading/Loading';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom';
import { Demo } from './pages/Demo';

export interface PageProps {
  client: AppClient;
}

export function App() {
  const [client, setClient] = useState<AppClient | undefined>();

  useEffect(() => {
    fetch('settings.json')
      .then((res) => res.json())
      .then((settings) => setClient(new AppClient({ BASE: settings.backend })));
  }, []);

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
          <>
            <div>
              <nav>
                <ul>
                  <li>
                    <Link to="/">Home</Link>
                  </li>
                  <li>
                    <Link to="/about">About</Link>
                  </li>
                  <li>
                    <Link to="/users">Users</Link>
                  </li>
                </ul>
              </nav>
            </div>

            <BrowserRouter>
              <Routes>
                <Route path="/" element={<Demo client={client} />} />
              </Routes>
            </BrowserRouter>
          </>
        ) : (
          <Loading />
        )}
      </div>
    </div>
  );
}
