import './Home.css';

function Home() {
  return (
    <div className="Home">
      <div className="header kirang-haerang-regular">
        <div>
          <h2>
       Welcome to GarageBuddy
       </h2>
      <img src="/garagebuddy.png" className="App-logo" alt="logo" />
      </div>
      </div>
      <div className="login">
        <form>
        <h2>Log in</h2>

          <input type="text" placeholder="Email" />
          <input type="password" placeholder="Password" />
          <button>Login</button>
        </form>
      </div>

    </div>
  );
}

export default Home;
