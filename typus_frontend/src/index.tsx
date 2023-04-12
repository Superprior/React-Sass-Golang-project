import React from 'react';
import ReactDOM from 'react-dom/client';
import IndexPage from './pages/IndexPage';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import SamplesListPage from './pages/SamplesListPage';
import SamplePage from './pages/SamplePage';


const router = createBrowserRouter([
  {
    path: '/',
    element: <IndexPage/>
  },
  {
    path: '/samples',
    element: <SamplesListPage />
  },
  {
    path: '/samples/:id',
    element: <SamplePage />
  }
])

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <RouterProvider router={router} />
);
