<template>
  <div id="app">
    <RwvHeader />
    <router-view />
    <RwvFooter />
  </div>
</template>

<style>
.sp-container {
  margin: 0 auto;
  max-width: 800px;
  padding: 1rem;
}
</style>

<script>
import RwvHeader from "@/components/TheHeader";
import RwvFooter from "@/components/TheFooter";

export default {
  name: "App",
    components: {
      RwvHeader,
      RwvFooter
  },
  created() {
    this.$store.dispatch("cosmos/init");
    // this.$store.dispatch("cosmos/entityFetch", {type: "poll", module: "voter" });
  },
};
</script>
