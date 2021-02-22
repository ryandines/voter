<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances path="ryandines.voter.voter" />
      <sp-token-send /> 
      <!-- this line is used by starport scaffolding # 4 -->
      <!-- <SpH3>
      Polls
      </SpH3> -->
      <!-- <poll-form /> -->
      <new-form path="ryandines.voter.voter" type="poll" :fields="[ ['creator', 1, 'string'] , ['title', 2, 'string'] , ['New option', 3, 'string'] ]" />    
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
import PollForm from "../components/PollForm";
import NewForm from '../components/NewForm.vue';

export default {
  components: {NewForm,  ...sp, PollForm }
};
</script>
