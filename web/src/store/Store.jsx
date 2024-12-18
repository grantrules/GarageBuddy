import React from "react";
import EventType from "./EventType";

class Store {
  constructor() {
    this.change = new EventType();
    this.data = new Map();
  }

  /// Get a key from the store. Doesn't interact with subscriptions.
  get(key) {
    return this.data.get(key);
  }

  /// Update a value in the store. Triggers a change notification.
  set(key, value) {
    this.data.set(key, value);
    this.change.emit(this.data);
    return this;
  }

  /// Example usage inside a React component:
  ///     const foo = store.use(() => store.get('foo'));
  /// The calling component will update if the *return value* of selector has changed.
  use(selector, deps = undefined) {
    // eslint-disable-next-line
    const [, forceUpdate] = React.useReducer((c) => c + 1, 0);
    // eslint-disable-next-line
    const ref = React.useRef(selector(this));
    this.change.use(
      () => {
        const next = selector(this);
        if (next !== ref.current) {
          ref.current = next;
          forceUpdate();
        }
      },
      deps || [selector],
    );

    return selector(this);
  }
}

export default Store;
