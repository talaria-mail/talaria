/**
 * auth returns a router gate that requires user authentication
 * to access restricted routes
 *
 * @param {*} store
 * @param {*} window
 */
export const auth = (store, window) => (to, from, next) => {
  // If in development, don't login redirect. Makes it a pain to iterate.
  if (process.env.NODE_ENV == "development") {
    next();
    return;
  }
  if (to.meta && to.meta.authenticated && !store.getters.getAuthToken) {
    window.location.href = "/login";
  } else {
    next();
  }
};
