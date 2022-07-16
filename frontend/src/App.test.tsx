import { render, screen } from '@testing-library/react';
import { App } from './App';
import GitHubLogin from './pages/githublogin'; '../src/pages/githublogin';

test('renders weight', () => {
  render(<GitHubLogin />);
  const linkElement = screen.getByText(/Go Lambda Skeleton/i);
  expect(linkElement).toBeInTheDocument();
});
