import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Dashboard from "../views/Dashboard.vue"
import SignIn from "../views/SignIn.vue"
import Matches from "../views/Matches.vue"
import Profile from "../views/Profile.vue"
import Discover from "../views/Discover.vue"
import Messages from "../views/Messages.vue"

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue")
  },
  {
    path: "/signin",
    name: "SignIn",
    component: SignIn
  },
  {
    path: "/dashboard/:username",
    name: "Dashboard",
    component: Dashboard,
    props: true
  },
  {
    path: "/matches",
    name: "Matches",
    component: Matches
  },
  {
    path: "/discover",
    name: "Discover",
    component: Discover
  },
  {
    path: "/profile",
    name: "Profile",
    component: Profile
  },
  {
    path: "/messages",
    name: "Messages",
    component: Messages
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
