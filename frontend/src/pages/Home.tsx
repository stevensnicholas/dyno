import React from 'react';
import { useNavigate } from 'react-router-dom';

interface Props {
  loggedIn: boolean;
}

const Home = ({ loggedIn }: Props) => {
  const navigation = useNavigate();

  function goToLogin() {
    navigation('/login');
  }

  function goToDashboard() {
    navigation('/dashboard');
  }
  return (
    <div>
      {loggedIn ? (
        <div>
          <h5>
            Logged in successfully. Go to Dashboard to view your test result.
          </h5>
          <button className="btn" onClick={() => goToDashboard()}>
            Dashboard
          </button>
        </div>
      ) : (
        <div>
          <h5>You are not logged in. Please login with your GitHub account.</h5>
          <button className="btn" onClick={() => goToLogin()}>
            Login
          </button>
        </div>
      )}
    </div>
  );
};

export default Home;
