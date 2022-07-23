import React from 'react';
import 'materialize-css/dist/css/materialize.min.css';
import styles from '../errors/Error.module.css';

const Error500 = () => {
  return (
    <div className={styles.browser_window}>
      <div className={styles.top_bar}>
        <div id={styles.close_circle} className={styles.circle}></div>
        <div id={styles.minimize_circle} className={styles.circle}></div>
        <div id={styles.maximize_circle} className={styles.circle}></div>
      </div>
      <div className="row">
        <div id={styles.site_layout_example_top} className="col s12">
          <p className={styles.caption_uppercase}>
            <h5 className={'center'}>Internal server error</h5>
          </p>
        </div>
        <div id={styles.site_layout_example_right} className="col s12 m12 l12">
          <div className="row center">
            <h1 className={styles.text_long_shadow + 'col s12 mt-3'}>500</h1>
            <p className="center white-text">
              Something has gone seriously wrong. Please try later.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Error500;
