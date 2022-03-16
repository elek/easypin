<script setup>

</script>

<template>
  <div>

    <main role="main" class="container" id="box">
      Upload your file with: <pre>curl -X POST -F file=@/path/file.extension
      https://easypin.storj-ipfs.com/api/v0/add</pre>

      And go to /pin/... after the file is uploaded
    </main>
  </div>
</template>

<style scoped>
</style>
