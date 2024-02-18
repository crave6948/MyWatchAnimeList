<template>
  <div id="stickyDiv">
    <img @click="open = !open" id="icon" src="../assets/plusicon.svg" alt="plus">
  </div>
  <div v-if="open">
    <div id="buttonContainer">
      <button @click="clicked">Add Anime</button>
    </div>
    <div id="panel">
      <input type="text" v-model="animeName" placeholder="Anime Name" />
      <input type="number" v-model="animeWatch" placeholder="Watched Episodes" />
      <input type="text" v-model="animeImage" placeholder="Anime Image" />
    </div>
  </div>
</template>

<script>
export default {
  props: ['addAnime'],
  data() {
    return {
      animeName: '',
      animeWatch: 0,
      animeImage: '',
      open: false
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
#icon {
  width: 50px;
  height: 50px;
  border-radius: 12px;
  position: relative;
}

#stickyDiv {
  position: fixed;
  bottom: 5%;
  right: 5%;
}

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
}</style>
