import React from 'react';
import Photo from "../images/logo/login-logo.png"
import "../css/style.css"

const Login = () => {
  return (
    <html>
      <body className="login-background">
        <div id="login-page">
          <div className="col s12 z-depth-4 card-panel">
            <div className="row">
              <div className="input-field col s12 center">
                <img src={Photo} alt="" className="circle responsive-img valign profile-image-login" />
                <p className="center login-form-text">Sign In</p>
                <h4 className="login-form-text">Use your Github Account</h4>
              </div>
            </div>
            <div className="row">
              <div className="login-box">
                <a href="#" className="btn-large waves-effect waves-teal col s12">
                  <span />
                  <span />
                  <span />
                  <span />Login</a>
              </div>
            </div>
          </div>
        </div>
      </body>
    </html>
  );
};

export default Login;
