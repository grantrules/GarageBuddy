import React, { useState } from 'react';
import PropTypes from 'prop-types';

const AuthContext = React.createContext({
  activeUser: null,
  login: () => {},
  invalidate: () => {},
  loginFailed: false,
});

const LOGGED_IN = { activeUser: true, loginFailed: false };
const LOGGED_OUT = { activeUser: false, loginFailed: false };
const LOGIN_FAILED = { activeUser: false, loginFailed: true };

function AuthProvider({ activeSession, children }) {
  const [{ activeUser, loginFailed }, setUserState] = useState({
    activeUser: activeSession,
    loginFailed: false,
  });

  const loginQuery = async () => {}
  const logoutQuery = async () => {}

  const invalidate = async () => logoutQuery({})
    .then(() => setUserState(LOGGED_OUT));

  const login = async (details) => {
    const { data, error } = await loginQuery(details);
    if (error || data.login === false) {
      setUserState(LOGIN_FAILED);
    } else {
      setUserState(LOGGED_IN);
    }
  };

  const context = {
    activeUser,
    login,
    invalidate,
    loginFailed,
  };
  return (<AuthContext.Provider value={context}>{children}</AuthContext.Provider>);
}

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
  activeSession: PropTypes.bool.isRequired,
};

export { AuthContext, AuthProvider };
