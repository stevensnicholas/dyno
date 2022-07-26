import React from 'react';
import { useLocation } from 'react-router-dom';
import Photo from '../login/login-logo.png';
import styles from '../login/login.module.css';
<<<<<<< HEAD
<<<<<<< HEAD

interface Props {
  clientID?: string;
}

const Login = ({ clientID }: Props) => {
  const redirectURL = 'http://localhost:8080/login';
  const path = '/';
  const scope = 'user:email';
  const githubURL =
    'https://github.com/login/oauth/authorize?' +
    clientID +
    '&redirect_uri=' +
    redirectURL +
    '?path=' +
    path +
    '&scope=' +
    scope;
=======
import { GitHubURL } from '../../client/models/GithubURL';
=======
>>>>>>> github url change

type LocationState = {
  from: {
    path: string;
  };
};

function GitHubURL(from: string) {
  const rootUrl = 'https://github.com/login/oauth/authorize';
  const options = {
    client_id: '' as string,
    redirect_url: 'http://localhost:8080/login' as string,
    scope: 'user:email',
    state: from,
  };

  const qs = new URLSearchParams(options);

  return `${rootUrl}?${qs.toString()}`;
}

const Login = () => {
  const location = useLocation();
  const from = (location.state as LocationState)?.from?.path || '/';
>>>>>>> github link added
  return (
    <html>
      <body className={styles.login_background}>
        <div id="login-page" className="row">
          <div className="col s12 z-depth-4 card-panel">
            <div className="row">
              <div className="input-field col s12 center">
                <img
                  src={Photo}
                  alt=""
                  className={
                    'circle responsive-img valign' + styles.profile_image_login
                  }
                />
                <p className={'center' + styles.login_form_text}>Sign In</p>
                <h4 className={styles.login_form_text}>
                  Use your Github Account
                </h4>
              </div>
            </div>
            <div className="row">
              <div className={styles.login_box}>
                <a
<<<<<<< HEAD
                  href={githubURL}
=======
                  href={GitHubURL(from)}
>>>>>>> github link added
                  className="btn-large waves-effect waves-teal col s12 black"
                >
                  Login
                </a>
              </div>
            </div>
          </div>
        </div>
      </body>
    </html>
  );
};

export default Login;