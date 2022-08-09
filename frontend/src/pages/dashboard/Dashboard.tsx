import React from 'react';
import { useNavigate } from 'react-router-dom';

const Dashboard = () => {
  const navigate = useNavigate();
  function getResult(id: number) {
    const path = `/testresult/${id}`;
    navigate(path, { state: { id: id } });
  }
  return (
    <div className="container">
      <h5>Test History</h5>
      <table>
        <thead>
          <tr>
            <th>Test ID</th>
            <th>OpenAPI file name</th>
            <th>Test Date</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>
              <button
                className="waves-effect waves-teal btn-flat"
                onClick={() => getResult(2)}
              >
                1
              </button>
            </td>
            <td>
              <button className="waves-effect waves-teal btn-flat">
                file 1
              </button>
            </td>
            <td>01/08/2022</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default Dashboard;
