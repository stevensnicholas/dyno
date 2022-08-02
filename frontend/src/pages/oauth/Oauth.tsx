import React, { useEffect, useState } from 'react';
import { AppClient } from '../../client';
import { Loading } from '../../components/Loading/Loading';
import { Link } from 'react-router-dom';

interface Props {
  client: AppClient;
}

const OAuth = (props: Props) => {
  const [loggedIn, setLoggedIn] = useState<boolean>(false);

  useEffect(() => {
    const code = window.location.search.substring(1).split('=')[2];
    props.client.default.endpointsAuthentication(code).then((res) => {
      window.localStorage.setItem('token', res.token);
      setLoggedIn(true);
    });
  }, [props]);

  return (
    <>
      {loggedIn ? (
        <div>
          <Link to="/dashboard">Redirect to Dashboard</Link>
        </div>
      ) : (
        <Loading />
      )}
    </>
  );
};

export default OAuth;
