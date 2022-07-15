import React, { useEffect, useState } from 'react';
import { Route, Routes, useNavigate, useLocation } from 'react-router-dom';
import RepoCard from './RepoCard';

type RepoProps = {
  username: string,
  name: string,
  language: string
}

const Repo = (props: RepoProps) => {
  const [repoData, setRepoData] = useState([] as any[]);
  const [datafetched, setDataFetched] = useState(false);
  
  useEffect(() => {
    fetch(`https://api.github.com/users/${props.username}/repos`)
        .then(res => res.json())
        .then(data => {
            setRepoData(data)
            setDataFetched(true)
        })
}, [])

  return (
    <div className="center-cards">
      {datafetched === true ? repoData.map((data, index) => <RepoCard key={index} name={data.name} language={data.language} />) : null}
    </div>
  );
};

export default Repo;