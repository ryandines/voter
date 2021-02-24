<template>
  <nav class="navbar navbar-light">
    <div class="container">
      <router-link class="navbar-brand" :to="{ name: 'home' }">
        Tokens
      </router-link>

              <!-- v-if="!isAuthenticated"  -->
      <ul 
        class="nav navbar-nav pull-xs-right ml-5">
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            exact
            :to="{ name: 'poll' }"
          >
          <i class="ion-compose"></i>
            Manage Polls
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            exact
            :to="{ name: 'vote' }"
          >
            <i class="ion-compose"></i>Vote
          </router-link>
        </li>
      </ul>
      <div class="signin">
       <sp-sign-in />
      </div>
      <!-- <ul v-else class="nav navbar-nav pull-xs-right">
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            exact
            :to="{ name: 'home' }"
          >
            Home
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            :to="{ name: 'article-edit' }"
          >
            <i class="ion-compose"></i>&nbsp;New Article
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            exact
            :to="{ name: 'settings' }"
          >
            <i class="ion-gear-a"></i>&nbsp;Settings
          </router-link>
        </li>
        <li class="nav-item" v-if="currentUser.username">
          <router-link
            class="nav-link"
            active-class="active"
            exact
            :to="{
              name: 'profile',
              params: { username: currentUser.username }
            }"
          >
            {{ currentUser.username }}
          </router-link>
        </li>
      </ul> -->
    </div>
  </nav>
</template>
<style scoped>
.signin{
  display: flex;
  flex-direction: column;
  /* align-items: center; */
}
.ml-5{
  margin-left: 1.5rem;
}
</style>
<script>
import { mapGetters } from "vuex";
import * as sp from "@tendermint/vue";

export default {
  name: "RwvHeader",
  computed: {
    ...mapGetters(["currentUser", "isAuthenticated"])
  },
  components: {
    ...sp
  }
};
</script>
