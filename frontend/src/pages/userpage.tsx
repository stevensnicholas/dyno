import React, { FC, ReactElement, useEffect, useState } from 'react';
import { Route, Routes, useNavigate, useLocation } from 'react-router-dom';
import Avatar from '../components/Utils/Avatar';
import Repo from '../components/Utils/Repo';


const UserPage = () => {
  const [userData, setUserData] = useState([] as any[]);
  const [datafetched, setDataFetched] = useState(false)
  const [username, setUsername] = useState('')
  const [loggedIn, setLoggedIn] = useState(false);

  const navigate = useNavigate();

  function fetchUserData(name: React.ElementType<any>) {

    fetch(`https://api.github.com/users/${name}`)
        .then(res => res.json())
        .then(data => {
            setUserData(data)
            console.log(data)
        })

  }

  useEffect(() => {
    fetch('/api')
        .then(res => res.json())
        .then((data) => {
            console.log(data.username)
            if (data.username === null) {
                navigate("/");
            }
            setUsername(data.username)
            setLoggedIn(true)
            fetchUserData(data.username)
        }).then(() => setDataFetched(true));
  }, [])

  
  
  return (
    <div className="center">
      {datafetched === false ? <h1>Loading User Data...</h1> :
        <>
        <p>{userData}</p>
        {/* <Avatar avatar={userData.avatar_url} name={userData.name} />
        <Repo username= {username} /> */}
        </>
      }
    </div>
  );
};

export default UserPage;