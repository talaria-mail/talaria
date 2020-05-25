import { UserServiceClient } from "@/proto/users_grpc_web_pb";
import {
  FetchUserRequest,
} from "@/proto/users_pb";

const BACKEND_URL = process.env.API_URL || "http://localhost:8081";
const client = new UserServiceClient(BACKEND_URL, null, null);

const state = {
  current: {},
  all: []
};

const getters = {
  getCurrentUser: state => state.current
};

const actions = {
  fetchCurrentUser({ commit, rootGetters }) {
    let username = rootGetters.getUsername
    console.log(username)
    let req = new FetchUserRequest()
    req.setUsername(username)
    let token = rootGetters.getAuthToken
    client.fetch(req, { authorization: token }, (err, resp) => {
      console.log(err)
      console.log(resp)
      let u = resp.getUser().toObject()
      commit("SET_CURRENT_USER", u)
    })
  }
};

const mutations = {
  SET_CURRENT_USER: (state, user) => (state.current = user),
};

export default {
  state,
  getters,
  actions,
  mutations
};
