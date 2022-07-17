import React from 'react';

type AvatarProps = {
  name: string,
  avatar: string
}

const Avatar = (props: AvatarProps) => {
  return (
    <div>
      <>
        <h2>{props.name}</h2>
        <img className="avatar" src={props.avatar} alt="user image" />
        <a className="logout"
          href={`http://localhost:5000/logout`}
        >Logout</a>
      </>
    </div>
  );
};

export default Avatar;