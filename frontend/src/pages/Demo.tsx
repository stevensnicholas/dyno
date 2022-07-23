import { useState } from 'react';
import { toast } from 'react-toastify';
import { AppClient } from '../client';

interface Props {
  client: AppClient;
}

export function Demo(props: Props) {
  const [request, setRequest] = useState<string>('');
  const [response, setResponse] = useState<string>('');

  return (
    <>
      <input
        type="text"
        onChange={(e) => {
          e.preventDefault();
          setRequest(e.target.value);
        }}
      />
      <input
        type="submit"
        onClick={() => {
          props.client.default
            .endpointsPostEcho({
              request,
            })
            .then((res) => {
              setResponse(res.result);
            })
            .catch((e) => {
              toast(e);
            });
        }}
      />
      <h2>{response}</h2>
    </>
  );
}
