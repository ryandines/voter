<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
      <SpH3>
        Polls
      </SpH3>
      <poll-form />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { PollForm, ...sp },
};
</script>
