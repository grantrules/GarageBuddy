import React, { useContext } from "react";
import PropTypes from "prop-types";
import Store from "./Store";

const StoreContext = React.createContext({});

function StoreProvider({ children }) {
  const [store] = React.useState(() => {
    const s = new Store();

    //s.set('studentName', '');

    return s;
  });

  return (
    <StoreContext.Provider value={store}>{children}</StoreContext.Provider>
  );
}

StoreProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

function useStore() {
  return useContext(StoreContext);
}

export { StoreContext, StoreProvider, useStore };
