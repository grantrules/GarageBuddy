import React from "react";
/// A single-type event emitter (.on(handler) instead of .on(eventName, handler))
export default class EventType {
  constructor() {
    this.listeners = new Set();
  }

  emit(data) {
    this.listeners.forEach((listener) => {
      listener(data);
    });
  }

  on(handler) {
    const listener = (data) => handler(data);

    this.listeners.add(listener);

    return () => this.listeners.delete(listener);
  }

  use(handler, deps = undefined) {
    // React.useEffect(() => this.on(handler), deps || [handler]);
    // eslint-disable-next-line
    React.useEffect(() => this.on(handler), [...deps, handler]);
  }
}
