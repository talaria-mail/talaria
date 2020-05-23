/**
 * auth returns a router gate that requires user authentication
 * to access restricted routes
 *
 * @param {*} store
 * @param {*} window
 */
export const auth = (store, window) => (to, from, next) => {
    if (to.meta && to.meta.authenticated && !store.getters.getAuthToken) {
      window.location.href = '/login'
    } else {
      next() 
    }
}
