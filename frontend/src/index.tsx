import React from 'react';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { App } from './App';
import { Navigation } from './navigation';
import reportWebVitals from './reportWebVitals';
import { createRoot } from 'react-dom/client';
import ReactDOM from 'react-dom/client';
import {BrowserRouter} from 'react-router-dom'

/* eslint-disable */

// const container = document.getElementById('root');
// const root = createRoot(container!); 
// root.render(
//   // <React.StrictMode>
//   //   <Navigation />
//   // </React.StrictMode>
//   <React.StrictMode>
//     <BrowserRouter>
//       <Navigation />
//     </BrowserRouter>
//   </React.StrictMode>
// );

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
     <BrowserRouter>
       <Navigation />
     </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
