import React, { useEffect, useState } from 'react';
import { AppClient } from '../client';
import { Loading } from '../components/Loading/Loading';

interface Props {
  client: AppClient;
}

const Dashboard = (props: Props) => {
  const [loggedIn, setLoggedIn] = useState<boolean>(false);

  useEffect(() => {
    const code = window.location.search.substring(1).split('=')[2];
    const url = `http://localhost:8080/login?code=${code}`;
    props.client.default.endpointsAuthentication(url).then((res) => {
      setLoggedIn(true);
    });
  }, [props]);
  return (
    <>
      {loggedIn ? (
        <div>
          <p>Dashboard</p>
        </div>
      ) : (
        <Loading />
      )}
    </>
  );
};

export default Dashboard;
