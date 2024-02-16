<template>
  <div class="box">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'Box'
  // Your component logic here
}
</script>

<style scoped>
.box {
  border: 1px solid #ffffff;
  border-top-right-radius: 25px;
  border-bottom-left-radius: 25px;
  background-color: #010101a2;
  padding: 10px;
  margin: 10px;
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-shadow: 0 12px 12px 0 rgba(0, 0, 0, 0.2);
}
</style>
