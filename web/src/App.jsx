import './App.css';

import { StoreProvider } from './store/StoreContext';
import { AuthProvider } from './auth/AuthProvider';

import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Home from './pages/Home';

const router = createBrowserRouter([{
  path: "/",
  element: <Home />,
}]);

function App() {
  return (
    <StoreProvider>
      <AuthProvider>
        <RouterProvider router={router} />
    </AuthProvider>
    </StoreProvider>
  );
}

export default App;
