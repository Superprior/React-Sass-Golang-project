import React from 'react';
import ReactDOM from 'react-dom/client';
import IndexPage from './pages/IndexPage';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import SamplesListPage from './pages/SamplesListPage';
import SamplePage from './pages/SamplePage';
import AdminPage from './pages/AdminPage';
import RegisterPage from './pages/RegisterPage';


const router = createBrowserRouter([
  {
    path: '/',
    element: <IndexPage/>
  },
  {
    path: '/register',
    element: <RegisterPage />
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
