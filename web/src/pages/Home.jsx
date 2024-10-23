import './Home.css';
import { Navigate } from 'react-router-dom';
import Authorized from '../auth/Authorized';

function Home() {

  return (<div>    <Authorized anonymous={true}>
      <Navigate to="/login" replace />

  </Authorized>
    <Authorized anonymous={false}>
      Hello
    </Authorized>
  </div>
  )
}

export default Home;
