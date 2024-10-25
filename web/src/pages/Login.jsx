import React from "react";
import PropTypes from "prop-types";

import { AuthContext } from "../auth/AuthProvider";

import "./Login.css";
import Authorized from "../auth/Authorized";
import { Navigate } from "react-router-dom";

function StepOne({ handleChange, toggleLogin, values }) {
  return (
    <div>
      <input
        type="text"
        placeholder="Email"
        onChange={handleChange("email")}
        value={values.email}
        autoFocus
      />
      <div className="loginActions">
        <button>Sign Up</button>{" "}
        <a href="#" onClick={toggleLogin}>
          Back to login
        </a>
      </div>
    </div>
  );
}

function StepTwo({ handleChange, toggleLogin, values }) {
  return (
    <div>
      <h2>Step 2</h2>
      <input
        type="password"
        placeholder="Password"
        name="password"
        value={values.password}
        onChange={handleChange("password")}
        autoFocus
      />
      <input
        type="password"
        placeholder="Confirm Password"
        name="password-confirm"
        values={values["password-confirm"]}
        onChange={handleChange("password-confirm")}
      />
      <button>Next</button>
    </div>
  );
}

function StepThree({ handleChange, toggleLogin, values }) {
  return (
    <div>
      <h2>What should we call you?</h2>
      <input
        type="text"
        placeholder="Your Name"
        name="firstName"
        value={values.name}
        onChange={handleChange("name")}
        autoFocus
      />
      <button>Next</button>
    </div>
  );
}

const stepPropTypes = {
  handleChange: PropTypes.func.isRequired,
  toggleLogin: PropTypes.func.isRequired,
  values: PropTypes.shape({
    email: PropTypes.string.isRequired,
    password: PropTypes.string.isRequired,
    "password-confirm": PropTypes.string.isRequired,
    name: PropTypes.string.isRequired,
  }).isRequired,
};

StepOne.propTypes = stepPropTypes;
StepTwo.propTypes = stepPropTypes;
StepThree.propTypes = stepPropTypes;

function Steps({ step, handleChange, toggleLogin, values }) {
  switch (step) {
    case 1:
      return (
        <StepOne
          handleChange={handleChange}
          toggleLogin={toggleLogin}
          values={values}
        />
      );
    case 2:
      return (
        <StepTwo
          handleChange={handleChange}
          toggleLogin={toggleLogin}
          values={values}
        />
      );
    case 3:
      return (
        <StepThree
          handleChange={handleChange}
          toggleLogin={toggleLogin}
          values={values}
        />
      );
    default:
      return null;
  }
}

Steps.propTypes = {
  step: PropTypes.number.isRequired,
  handleChange: PropTypes.func.isRequired,
  toggleLogin: PropTypes.func.isRequired,
  values: PropTypes.object.isRequired,
};

function RegisterForm({ toggleLogin }) {
  const [step, setStep] = React.useState(1);
  const [values, setValues] = React.useState({
    email: "",
    password: "",
    "password-confirm": "",
    name: "",
  });

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value });
  };

  const register = () => {
    return fetch("/api/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(values),
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (step === 3) {
      register();
    } else {
      setStep(step + 1);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <Steps
        step={step}
        handleChange={handleChange}
        toggleLogin={toggleLogin}
        values={values}
      />
    </form>
  );
}

RegisterForm.propTypes = {
  toggleLogin: PropTypes.func.isRequired,
};

function LoginForm({ toggleLogin }) {
  const { login, loginFailed } = React.useContext(AuthContext);

  const [values, setValues] = React.useState({ email: "", password: "" });

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    login(values);
  };

  return (
    <form onSubmit={handleSubmit}>
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
      <div> {loginFailed && `Failed`}</div>
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
    <>
      <Authorized anonymous={false}>
        <Navigate to="/" replace />
      </Authorized>
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
    </>
  );
}

export default Login;
