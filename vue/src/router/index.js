import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import Poll from "../views/Poll.vue";
import Vote from "../views/Vote.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/poll",
    name: "poll",
    component: Poll,
  },
  {
    path: "/vote",
    name: "vote",
    component: Vote,
  },
  {
    path: "/",
    name: "home",
    component: Index,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
