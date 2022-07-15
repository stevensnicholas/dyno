import React, { useEffect, useState } from 'react';

interface RepoCardProps{
  name: string,
  language: string
}

const RepoCard = (props: RepoCardProps) => {
  return (
    <div className='card'>
      <span><b>Name: </b>{props.name}</span> <hr></hr>
      {props.language !== null ? <span><b>Language: </b>{props.language}</span> : <span><b>Language: </b>null</span>}
    </div>
  );
};

export default RepoCard;