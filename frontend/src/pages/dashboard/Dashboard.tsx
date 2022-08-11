import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

export interface ListResult {
  id: number;
  title: string;
  endpoint: string;
  httpmethod: string;
  TimeDelay: string;
  AsyncTime: string;
  PreviousResponse: string;
  ErrorType: string;
  details: string;
}

export const ResultList: ListResult[] = [
  {
    id: 385,
    title: 'UseAfterFreeChecker Invalid 20x Response',
    endpoint: '//api/blog/posts/115',
    httpmethod: 'DELETE',
    TimeDelay: '! producer_timing_delay 0',
    AsyncTime: '! max_async_wait_time 20',
    PreviousResponse: 'HTTP/1.1 204 No Content response:',
    ErrorType: 'UseAfterFreeChecker',
    details:
      'Detects that a deleted resource can still being accessed after deletion.',
  },
  {
    id: 383,
    title: 'InvalidDynamicObjectChecker Invalid 20x Response',
    endpoint: '//api/blog/posts/105?injected_query_string=123',
    httpmethod: 'GET',
    TimeDelay: '! producer_timing_delay 0',
    AsyncTime: '! max_async_wait_time 0',
    PreviousResponse:
      'HTTP/1.1 200 OK response:Connection: keep-alive\r\r{"id":105,"body":"my first blog post"}',
    ErrorType: 'InvalidDynamicObjectChecker',
    details:
      'Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.',
  },
  {
    id: 376,
    title: 'PayloadBodyChecker Invalid 500 Response',
    endpoint: '//api/blog/posts',
    httpmethod: 'POST',
    TimeDelay: '! producer_timing_delay 0',
    AsyncTime: '! max_async_wait_time 0',
    PreviousResponse:
      'HTTP/1.1 201 Created response:Connection: keep-alive {"id":107,"body":"my first blog post"}',
    ErrorType: 'PayloadBodyChecker',
    details: 'Detects 500 errors when fuzzing the JSON bodies of requests.',
  },
  {
    id: 353,
    title: 'InvalidDynamicObjectChecker Invalid 20x Response',
    endpoint: '//api/blog/posts/10973?injected_query_string=123',
    httpmethod: 'GET',
    TimeDelay: '! producer_timing_delay 0',
    AsyncTime: '! max_async_wait_time 0',
    PreviousResponse:
      'HTTP/1.1 200 OK response:Connection: keep-alive\r\r{"id":10973,"body":"my first blog post"}',
    ErrorType: 'InvalidDynamicObjectChecker',
    details:
      'Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.',
  },
  {
    id: 344,
    title: 'InvalidDynamicObjectChecker Invalid 20x Response',
    endpoint: '//api/blog/posts/10976?injected_query_string=123 HTTP/1.1',
    httpmethod: 'PUT',
    TimeDelay: '! producer_timing_delay 0',
    AsyncTime: '! max_async_wait_time 0',
    PreviousResponse: 'HTTP/1.1 204 No Content response:',
    ErrorType: 'InvalidDynamicObjectChecker',
    details:
      'Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.',
  },
  {
    id: 339,
    title: 'PayloadBodyChecker Invalid 500 Response',
    endpoint: '//api/blog/posts',
    httpmethod: 'POST',
    TimeDelay: '! producer_timing_delay 0',
    AsyncTime: '! max_async_wait_time 0',
    PreviousResponse:
      'HTTP/1.1 201 Created response:Connection: keep-alive\r\r{"id":10975,"body":"my first blog post"}',
    ErrorType: 'PayloadBodyChecker',
    details: 'Detects 500 errors when fuzzing the JSON bodies of requests.',
  },
  {
    id: 335,
    title: 'UseAfterFreeChecker Invalid 20x Response',
    endpoint: '//api/blog/posts/10983',
    httpmethod: 'DELETE',
    TimeDelay: '! producer_timing_delay 0',
    AsyncTime: '! max_async_wait_time 20',
    PreviousResponse: 'HTTP/1.1 204 No Content response:',
    ErrorType: 'UseAfterFreeChecker',
    details:
      'Detects that a deleted resource can still being accessed after deletion.',
  },
];

const Dashboard = () => {
  const navigate = useNavigate();
  const [mostError, setMostError] = useState<string>('');
  const [occurrence, setOccurrence] = useState<number>(0);

  function getResult(id: number) {
    const path = `/testresult/${id}`;
    navigate(path, { state: { id: id } });
  }

  useEffect(() => {
    let useAfter = 0;
    let payload = 0;
    let invalidObject = 0;
    for (const item of ResultList) {
      if (item.ErrorType === 'UseAfterFreeChecker') {
        useAfter++;
      }
      if (item.ErrorType === 'PayloadBodyChecker') {
        payload++;
      }
      if (item.ErrorType === 'InvalidDynamicObjectChecker') {
        invalidObject++;
      }
    }
    if (useAfter === Math.max(useAfter, payload, invalidObject)) {
      setMostError('UseAfterFreeChecker');
      setOccurrence(Math.max(useAfter, payload, invalidObject));
    }
    if (payload === Math.max(useAfter, payload, invalidObject)) {
      setMostError('PayloadBodyChecker');
      setOccurrence(Math.max(useAfter, payload, invalidObject));
    }
    if (invalidObject === Math.max(useAfter, payload, invalidObject)) {
      setMostError('InvalidDynamicObjectChecker');
      setOccurrence(Math.max(useAfter, payload, invalidObject));
    }
  }, []);

  return (
    <div className="container">
      <h5>Test History</h5>
      <table>
        <thead>
          <tr>
            <th>Test ID</th>
            <th>Endpoint</th>
            <th>HTTP Method</th>
          </tr>
        </thead>
        <tbody>
          {ResultList.map(({ id, endpoint, httpmethod }: ListResult) => {
            return (
              <tr key={id}>
                <td>
                  <button className="btn-flat" onClick={() => getResult(id)}>
                    {id}
                  </button>
                </td>
                <td>
                  <button className="btn-flat" onClick={() => getResult(id)}>
                    {endpoint}
                  </button>
                </td>
                <td>
                  <button className="btn-flat" onClick={() => getResult(id)}>
                    {httpmethod}
                  </button>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
      <h5>Most Frequent Error Type</h5>
      <table>
        <thead>
          <tr>
            <th>Error Type</th>
            <th>Occurrence</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{mostError}</td>
            <td>{occurrence}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default Dashboard;
