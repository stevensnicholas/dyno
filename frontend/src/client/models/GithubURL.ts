export function GitHubURL(from: string) {
  const rootUrl = 'https://github.com/login/oauth/authorize';
  const options = {
    client_id: 'c2f25de6e7cfc7713944' as string,
    redirect_url: 'http://localhost:8080/login' as string,
    scope: 'user:email',
    state: from,
  };

  const qs = new URLSearchParams(options);

  return `${rootUrl}?${qs.toString()}`;
}
