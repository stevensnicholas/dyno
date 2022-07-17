import React, { useEffect, useState } from 'react';

const Login = () => {
  const GITHUB_CLIENT_ID = 'c2f25de6e7cfc7713944';
  const gitHubRedirectURL = 'http://localhost:8080/api/auth/github';
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