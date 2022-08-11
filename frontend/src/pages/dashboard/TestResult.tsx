import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { ListResult, ResultList } from './Dashboard';
import '../dashboard/testresult.css';

const TestResult = () => {
  const [result, setResult] = useState<ListResult | undefined>();
  const url = window.location.toString();
  const ID = Number(url.split('http://localhost:3000/testresult/')[1]);
  const navigate = useNavigate();

  function goDashboard() {
    navigate('/dashboard');
  }

  useEffect(() => {
    for (const item of ResultList) {
      if (ID === item.id) {
        setResult(item);
        break;
      }
    }
  }, [ID]);
  return (
    <div className="container">
      {result ? (
        <React.Fragment>
          <div id="result-text">
            <h4>Test Result ID: {result.id}</h4>
            <div className="divider"></div>
            <div className="section" id="section">
              <h5>Title</h5>
              <p>{result.title}</p>
            </div>
            <div className="divider"></div>
            <div className="section" id="section">
              <h5>ErroType</h5>
              <p>{result.ErrorType}</p>
            </div>
            <div className="divider"></div>
            <div className="section" id="section">
              <h5>Details</h5>
              <p>{result.details}</p>
            </div>
            <div className="divider"></div>
            <div className="section" id="section">
              <h5>Endpoint & HTTP Method</h5>
              <p>Endpoint: {result.endpoint}</p>
              <p>Method: {result.httpmethod}</p>
            </div>
            <div className="divider"></div>
            <div className="section" id="section">
              <h5>Time Delay & Async Time</h5>
              <p>Time Delay: {result.TimeDelay}</p>
              <p>Async Time: {result.AsyncTime}</p>
            </div>
            <div className="divider"></div>
            <div className="section" id="section">
              <h5>Previous Response</h5>
              <p>{result.PreviousResponse}</p>
            </div>
            <div id="button-dashboard">
              <button className="btn" onClick={() => goDashboard()}>
                Go Back to Dashboard
              </button>
            </div>
          </div>
        </React.Fragment>
      ) : (
        <p>Loading</p>
      )}
    </div>
  );
};

export default TestResult;
