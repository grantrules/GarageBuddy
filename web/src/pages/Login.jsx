import React from "react";
import PropTypes from "prop-types";

import "./Login.css";

function RegisterForm({ toggleLogin }) {
  const [values, setValues] = React.useState({ email: "" });

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value });
  };

  const register = (e) => {
    e.preventDefault();
    fetch("/api/register", { method: "POST", body: JSON.stringify(values) });
  };
  return (
    <form onSubmit={register}>
      <input
        type="text"
        placeholder="Email"
        onChange={handleChange("email")}
        value={values.email}
      />
      <div className="loginActions">
        <button>Sign Up</button>{" "}
        <a href="#" onClick={toggleLogin}>
          Back to login
        </a>
      </div>
    </form>
  );
}

RegisterForm.propTypes = {
  toggleLogin: PropTypes.func.isRequired,
};

function LoginForm({ toggleLogin }) {
  const [values, setValues] = React.useState({ email: "", password: "" });

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value });
  };

  const login = (e) => {
    e.preventDefault();
    fetch("/api/login", { method: "POST", body: JSON.stringify(values) });
  };

  return (
    <form onSubmit={login}>
      <input
        type="text"
        placeholder="Email"
        onChange={handleChange("email")}
        value={values.email}
      />
      <input
        type="password"
        placeholder="Password"
        onChange={handleChange("password")}
        value={values.password}
      />
      <div className="loginActions">
        <button>Login</button>{" "}
        <a href="#" onClick={toggleLogin}>
          New? Sign up!
        </a>
      </div>
    </form>
  );
}

LoginForm.propTypes = {
  toggleLogin: PropTypes.func.isRequired,
};

function Login() {
  const [isLogin, setIsLogin] = React.useState(true);

  const toggleLogin = (e) => {
    e.preventDefault();
    setIsLogin(!isLogin);
  };

  return (
    <div className="LoginPage">
      <div className="header kirang-haerang-regular">
        <div>
          <h2>Welcome to GarageBuddy</h2>
          <img src="/garagebuddy.png" className="App-logo" alt="logo" />
        </div>
      </div>
      <div className="login">
        <div>
          <h2>{isLogin ? "Log in" : "Sign up"}</h2>
          <a
            href="/api/oauth2/login/google"
            target="_blank"
            className="button googleButton"
          >
            Sign in with Google
          </a>
          {isLogin ? (
            <LoginForm toggleLogin={toggleLogin} />
          ) : (
            <RegisterForm toggleLogin={toggleLogin} />
          )}
        </div>
      </div>
    </div>
  );
}

export default Login;
