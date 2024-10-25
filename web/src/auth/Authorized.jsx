import React from "react";
import PropTypes from "prop-types";
import { AuthContext } from "./AuthProvider";

function Authorized({ anonymous = false, both = false, children }) {
  const { activeUser } = React.useContext(AuthContext);
  console.log(activeUser);
  if (activeUser !== anonymous || both) return <>{children}</>;
  return null;
}

Authorized.propTypes = {
  children: PropTypes.node.isRequired,
  anonymous: PropTypes.bool,
  both: PropTypes.bool,
};
export default Authorized;
