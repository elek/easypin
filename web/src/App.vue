<script setup></script>

<template>
  <div>

    <main role="main" class="container" id="box">
      <Deposit></Deposit>
    </main>
  </div>

</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}


#box {
  max-width: 800px;
}
</style>
<script>
import Deposit from "./components/Deposit.vue";
export default {
  components: {Deposit}
}
</script>