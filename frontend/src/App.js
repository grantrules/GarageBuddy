import './App.css';

import { StoreProvider } from './store/StoreContext';
import { AuthProvider } from './auth/AuthProvider';

function App() {
  return (
    <StoreProvider>
      <AuthProvider>
    <div className="App">
      <header className="App-header">
       Hi
      </header>
    </div>
    </AuthProvider>
    </StoreProvider>
  );
}

export default App;
