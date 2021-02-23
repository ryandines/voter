import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import Poll from "../views/Poll.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "login",
    component: Index,
  },
  {
    path: "/register",
    name: "register",
    component: Poll,
  },
  {
    path: "/home",
    name: "home",
    component: Index,
  },
  {
    path: "/",
    component: Index,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
