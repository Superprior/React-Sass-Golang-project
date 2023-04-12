import React from 'react';
import ReactDOM from 'react-dom/client';
import IndexPage from './pages/IndexPage';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import SamplesListPage from './pages/SamplesListPage';
import SamplePage from './pages/SamplePage';
import AdminPage from './pages/AdminPage';


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
  },
  {
    path: '/admin',
    element: <AdminPage />
  }
])

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <RouterProvider router={router} />
);
