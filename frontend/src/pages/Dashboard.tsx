import React from 'react';

const Dashboard = () => {
  const username =
    window.location.search[0] === '?'
      ? window.location.search.substring(1).split('=')[1]
      : null;

  return (
    <div>
      <p>Dashboard</p>
      <p>{username}</p>
    </div>
  );
};

export default Dashboard;
