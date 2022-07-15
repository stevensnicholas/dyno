import React, { useEffect, useState } from 'react';

const Login = () => {
  const GITHUB_CLIENT_ID = process.env.REACT_APP_GITHUB_CLIENT_ID;
  const gitHubRedirectURL = 'http://localhost:5000/api/auth/github';
  const path = '/';
  return (
    <div>
      <a
        href={`https://github.com/login/oauth/authorize?client_id=${GITHUB_CLIENT_ID}&redirect_uri=${gitHubRedirectURL}?path=${path}&scope=user:email`}
      >Login with github</a>
    </div>
  );
};

export default Login;
