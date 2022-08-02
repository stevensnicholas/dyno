import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

const Dashboard = () => {
  const [loggedIn, setLoggedIn] = useState<boolean>(false);

  useEffect(() => {
    const Token = window.localStorage.getItem('token');
    if (Token !== '') {
      setLoggedIn(true);
    }
  }, [setLoggedIn]);

  return (
    <>
      {loggedIn ? (
        <div>
          <p>Dashboard</p>
          <p>Successfully logged in. Your test results will be displayed</p>
        </div>
      ) : (
        <div>
          <Link to="/login">Please login with your GitHub account first</Link>
        </div>
      )}
    </>
  );
};

export default Dashboard;
