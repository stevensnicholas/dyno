import { render, screen } from '@testing-library/react';
import { App } from './App';

test('renders weight', () => {
  render(<App />);
  const linkElement = screen.getByText(/Go Lambda Skeleton/i);
  expect(linkElement).toBeInTheDocument();
});
