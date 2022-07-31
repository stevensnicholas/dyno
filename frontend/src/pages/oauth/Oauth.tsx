import React, { Component } from 'react';
import axios from 'axios';

export default class Oauth extends Component {
  componentDidMount(): void {
    this.oauthToken();
  }
  // eslint-disable-next-line
  render() {
    return (
      <div>
        <h3>Loading</h3>
      </div>
    );
  }
  // eslint-disable-next-line
  oauthToken = async () => {
    const code = window.location.search.substring(1).split('=')[2];
    const accessTokenRsp = await axios({
      method: 'GET',
      url: `http://localhost:8080/login?code=${code}`,
      headers: { accept: 'application/json' },
    });

    const accessToken = accessTokenRsp.data.token;

    const userInfo = await axios({
      method: 'GET',
      url: 'https://api.github.com/user',
      headers: {
        accept: 'application/json',
        Authorization: `token ${accessToken}`,
      },
    });

    window.location.replace(
      `http://localhost:3000/dashboard?name=${userInfo.data.login}`
    );
  };
}
