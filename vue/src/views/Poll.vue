<template>
  <div>
    <div class="sp-container">
      <!-- <sp-bank-balances path="ryandines.voter.voter" /> -->
      <!-- <sp-token-send />  -->
      <!-- this line is used by starport scaffolding # 4 -->
      <new-form 
        path="ryandines.voter.voter" t
        type="poll"
        />
      <!-- <poll-list />  -->
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
import NewForm from '../components/NewForm.vue';
import PollList from "../components/PollList";

export default {
  components: { NewForm, PollList, ...sp }
};
</script>
