import React from 'react';
import Photo from '../login/login-logo.png';
import styles from '../login/login.module.css';
const Login = () => {
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
                  href="http://localhost:3000/login"
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
