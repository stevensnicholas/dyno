import { useEffect, useState } from 'react';
import './App.css';
import { AppClient } from './client';
import Background from './background.jpg';
import { Loading } from './components/Loading/Loading';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export interface PageProps {
  client: AppClient;
}

export function App() {
  const [client, setClient] = useState<AppClient | undefined>();
  const [request, setRequest] = useState<string>('');
  const [response, setResponse] = useState<string>('');

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
        backgroundImage: `url(${Background})`,
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
            <input
              type="text"
              onChange={(e) => {
                e.preventDefault();
                setRequest(e.target.value);
              }}
            />
            <input
              type="submit"
              onClick={() => {
                client.default
                  .cmdEndpointsPostEcho({
                    request,
                  })
                  .then((res) => {
                    setResponse(res.result);
                  })
                  .catch((e) => {
                    toast(e);
                  });
              }}
            />
            <h2>{response}</h2>
          </>
        ) : (
          <Loading />
        )}
      </div>
    </div>
  );
}
