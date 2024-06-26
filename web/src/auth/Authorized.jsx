import React, { useContext } from 'react';
import PropTypes from 'prop-types';
import { AuthContext } from './AuthProvider';

function Authorized({ anonymous, both, children }) {
  const { activeUser } = useContext(AuthContext);
  if (activeUser !== anonymous || both) return (<>{children}</>);
  return null;
}
Authorized.defaultProps = {
  anonymous: false,
  both: false,
};
Authorized.propTypes = {
  children: PropTypes.node.isRequired,
  anonymous: PropTypes.bool,
  both: PropTypes.bool,
};
export default Authorized;
