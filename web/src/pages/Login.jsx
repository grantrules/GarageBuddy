import React from 'react';

import './Login.css';

function Login() {

    const [values, setValues] = React.useState({ email: '', password: '' });

    const handleChange = name => (event) => { setValues({ ...values, [name]: event.target.value }); };

    /*const submitForm = (event) => {
        event.preventDefault();
        console.log(values);
    }*/


  const [isLogin, setIsLogin] = React.useState(true);

  const toggleLogin = (e) => {
    e.preventDefault();
    setIsLogin(!isLogin);
  }
    
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
          <a href="/api/oauth2/login/google" target="_blank" className="button googleButton">Sign in with Google</a>
          {isLogin ?
          <div>
          <input type="text" onChange={handleChange('email')} placeholder="Email" />
          <input type="password" onChange={handleChange('password')} placeholder="Password" />
          <div className="loginActions">
            <button>Login</button> <a href="#" onClick={toggleLogin}>New? Sign up!</a>
          </div>
          </div>
          :

          <div>
          <input type="text" placeholder="Email" />
          <div className="loginActions">
            <button>Sign Up</button> <a href="#" onClick={toggleLogin}>Back to login</a>
          </div>
          </div>
}
        </form>
      </div>

    </div>
  );
}

export default Login;
