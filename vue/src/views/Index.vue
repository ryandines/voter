<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances path="ryandines.voter.voter" />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		  <sp-type-form path="ryandines.voter.voter" type="poll" :fields="[ ['creator', 1, 'string'] , ['title', 2, 'string'] , ['New option', 3, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
