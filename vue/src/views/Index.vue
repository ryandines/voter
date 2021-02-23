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
      <new-form 
        path="ryandines.voter.voter" t
        type="poll"
        />    
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
import NewForm from '../components/NewForm.vue';

export default {
  components: { NewForm,  ...sp }
};
</script>
