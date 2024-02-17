<template>
  <div id="buttonContainer">
    <button @click="clicked">Add Anime</button>
  </div>
  <div id="panel">
    <input type="text" v-model="animeName" placeholder="Anime Name" />
    <input type="number" v-model="animeWatch" placeholder="Watched Episodes" />
    <input type="text" v-model="animeImage" placeholder="Anime Image" />
  </div>
</template>

<script>
export default {
  props: ['addAnime'],
  data() {
    return {
      animeName: '',
      animeWatch: 0,
      animeImage: ''
    }
  },
  methods: {
    clicked() {
      this.addAnime(this.animeName, this.animeWatch, this.animeImage)
    }
  }
}
</script>

<style>
#buttonContainer {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}
#panel {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 20px;
}
</style>
