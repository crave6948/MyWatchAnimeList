<template>
  <AnimeTable/>
  <ActionMenu/>
</template>

<script>
import AnimeTable from './components/AnimeTable.vue';
import ActionMenu from './components/ActionMenu.vue';
export default {
  components: {
    AnimeTable,
    ActionMenu,
  },
};
</script>

<style scoped>
</style>
