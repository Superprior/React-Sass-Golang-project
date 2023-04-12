import React from 'react';
import ReactDOM from 'react-dom/client';
import IndexPage from './pages/IndexPage';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import SamplesPage from './pages/SamplesPage';


const router = createBrowserRouter([
  {
    path: '/',
    element: <IndexPage/>
  },
  {
    path: '/samples',
    element: <SamplesPage />
  }
])

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <RouterProvider router={router} />
);
