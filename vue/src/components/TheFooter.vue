<template>
  <footer>
    <div class="container">
      <router-link class="logo-font" :to="{ name: 'home', params: {} }">
         Distributed Indelible Election Ledger 
      </router-link>
      <span class="attribution">
        The safest way to record elections.
        Engineered with the Cosmos SDK.
        <a rel="noopener noreferrer" target="blank" href="https://docs.cosmos.network/"
          >Read more here.
        </a>
      </span>
    </div>
  </footer>
</template>

<script>
export default {
  name: "RwvFooter"
};
</script>
