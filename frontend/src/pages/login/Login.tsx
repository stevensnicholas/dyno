import React from 'react';
import Photo from '../login/login-logo.png';
import styles from '../login/login.module.css';

interface Props {
  clientID?: string;
}

const Login = ({ clientID }: Props) => {
  const hostname = window.location.origin;
  const redirectURL = `${hostname}/oauth`;
  const scope = 'user:email';
  const githubURL = `https://github.com/login/oauth/authorize?client_id=${clientID}&redirect_uri=${redirectURL}?scope=${scope}`;

  function GithubLogin() {
    window.location.replace(githubURL);
  }

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
                <button
                  onClick={() => GithubLogin()}
                  className="btn-large waves-effect waves-teal col s12 black"
                >
                  Login
                </button>
              </div>
            </div>
          </div>
        </div>
      </body>
    </html>
  );
};

export default Login;
