import "./Home.css";
import { Navigate } from "react-router-dom";
import Authorized from "../auth/Authorized";
import MyCars from "../components/MyCars";

function Home() {
  return (
    <div>
      <Authorized anonymous={true}>
        <Navigate to="/login" replace />
      </Authorized>
      <Authorized>
        <h1 className="home">SUPPPP</h1>
        <MyCars />
      </Authorized>
    </div>
  );
}

export default Home;
