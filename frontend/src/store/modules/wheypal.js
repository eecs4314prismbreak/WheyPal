import axios from "axios";

const state = {
  userRecommendations: null,
  userPreferences: null, // expect an object with various preferences (ie location, activity, etc.) to be displayed on profile page
  userProfileImage: null,
  userName: null,
  userEmail: null,
  userPassword: null,
  userId: null,
  userToken: null,
  userExpiry: null
};

const getters = {};

const actions = {
  async createUser({ commit }, body) {
    console.log("Creating user");
    const url = "http://localhost:8081/user";
    const response = await axios.post(url, body);
    const payload = response.data;
    payload["body"] = body;
    commit("CREATE_USER", payload);
  },
  async getRecommendations({ commit }, body) {
    console.log("Getting recommendations");
    const url = "http://localhost:8081/user";
    const response = await axios.get(url, {
      headers: {
        Authorization: `Bearer ${body.userToken}`
      }
    });
    const payload = response.data.filter(el => el.email !== body.userEmail);
    // console.table(payload)
    // console.log("%c%s", "background: dodgerblue; padding: 5px; border-radius: 5x", "stylized log!")
    commit("GET_RECOMMENDATIONS", payload);
  },
  async loginUser({ commit }, body) {
    console.log("Login user");
    const url = "http://localhost:8081/login";
    const response = await axios.post(url, body);
    const payload = response.data;
    payload["body"] = body;
    commit("LOGIN_USER", payload);
  },
  async logOffUser({ commit }) {
    console.log("Sign out user");
    commit("LOGOFF_USER");
  }
};

const mutations = {
  CREATE_USER: (state, payload) => {
    state.userId = payload.id;
    state.userEmail = payload.email;
    state.userToken = payload.token;
    state.userExpiry = payload.expiry;
    state.userName = payload.body.name;
    state.userEmail = payload.body.email;
    state.userPassword = payload.body.password;
  },
  GET_RECOMMENDATIONS: (state, payload) => {
    state.userRecommendations = payload;
  },
  LOGIN_USER: (state, payload) => {
    state.userToken = payload.token;
    state.userExpiry = payload.expiry;
    state.userId = payload.id;
    state.userEmail = payload.body.email;
    state.userPassword = payload.body.password;
  },
  LOGOFF_USER: state => {
    (state.userRecommendations = null),
      (state.userPreferences = null),
      (state.userProfileImage = null),
      (state.userName = null),
      (state.userEmail = null),
      (state.userPassword = null),
      (state.userId = null),
      (state.userToken = null),
      (state.userExpiry = null);
  }
};

export default {
  state,
  getters,
  actions,
  mutations
};
