<template>
  <div class="block">
    <div class="preloader"></div>
  </div>
</template>

<script>
export default {
  name: "Preloader"
}
</script>

<style scoped>
.block {
  text-align: center;
  margin-bottom: 35px;
}

.preloader {
  display: inline-block;
  width: 40px;
  height: 40px;
  border: solid 3px;
  border-color: #555555 transparent;
  border-radius: 50%;
  -webkit-animation:
      rota 1s ease-in-out infinite;
  -moz-animation:
      rota 1s ease-in-out infinite;
  -ms-animation:
      rota 1s ease-in-out infinite;
  -o-animation:
      rota 1s ease-in-out infinite;
  animation:
      rota 1s ease-in-out infinite;
}

@-webkit-keyframes rota {
  to { -webkit-transform: rotate(360deg);}
}

@-moz-keyframes rota {
  to { -moz-transform: rotate(360deg); }
}

@-ms-keyframes rota {
  to { -ms-transform: rotate(360deg); }
}

@-o-keyframes rota {
  to { -o-transform: rotate(360deg); }
}

@keyframes rota {
  to { transform: rotate(360deg); }
}
</style>